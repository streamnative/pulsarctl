// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package internal

import (
	"errors"
	"os"
	"path"
	"reflect"
	"sort"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

type ConfigAccess interface {
	// GetLoadingPrecedence returns the slice of files that should be used for loading and inspecting the config
	GetLoadingPrecedence() []string
	// GetStartingConfig returns the config that subcommands should being operating against.
	// It may or may not be merged depending on loading rules
	GetStartingConfig() (*cmdutils.Config, error)
	// GetDefaultFilename returns the name of the file you should write into (create if necessary),
	// if you're trying to create a new stanza as opposed to updating an existing one.
	GetDefaultFilename() string
}

type PathOptions struct {
	// GlobalFile is the full path to the file to load as the global (final) option
	GlobalFile string

	// GlobalFileSubpath is an optional value used for displaying help
	GlobalFileSubpath string

	LoadingRules *ClientConfigLoadingRules
}

type ClientConfig interface {
	// RawConfig returns the merged result of all overrides
	RawConfig() (cmdutils.Config, error)
	// ConfigAccess returns the rules for loading/persisting the config.
	ConfigAccess() ConfigAccess
}

func (o *PathOptions) GetLoadingPrecedence() []string {
	return []string{o.GlobalFile}
}

func (o *PathOptions) GetStartingConfig() (*cmdutils.Config, error) {
	// don't mutate the original
	loadingRules := *o.LoadingRules
	loadingRules.Precedence = o.GetLoadingPrecedence()

	clientConfig := NewNonInteractiveDeferredLoadingClientConfig(&loadingRules, new(cmdutils.ConfigOverrides))
	rawConfig, err := clientConfig.RawConfig()
	if os.IsNotExist(err) {
		return cmdutils.NewConfig(), nil
	}
	if err != nil {
		return nil, err
	}

	return &rawConfig, nil
}

func (o *PathOptions) GetDefaultFilename() string {
	return o.GlobalFile
}

func NewDefaultPathOptions() *PathOptions {
	ret := &PathOptions{
		GlobalFile: RecommendedHomeFile,

		GlobalFileSubpath: path.Join(RecommendedHomeDir, RecommendedFileName),

		LoadingRules: NewDefaultClientConfigLoadingRules(),
	}
	ret.LoadingRules.DoNotResolvePaths = true

	return ret
}

// ModifyConfig takes a Config object and write filed of Config struct to file
func ModifyConfig(configAccess ConfigAccess, newConfig cmdutils.Config, relativizePaths bool) error {
	possibleSources := configAccess.GetLoadingPrecedence()
	// sort the possible pulsar config files so we always "lock" in the same order
	// to avoid deadlock (note: this can fail w/ symlinks, but... come on).
	sort.Strings(possibleSources)
	for _, filename := range possibleSources {
		if err := lockFile(filename); err != nil {
			return err
		}
		defer unlockFile(filename)
	}

	startingConfig, err := configAccess.GetStartingConfig()
	if err != nil {
		return err
	}

	// We need to find all differences, locate their original files, read a partial config to modify
	// only that stanza and write out the file.
	// Special case the test for current context and preferences since those always write to the default file.
	if reflect.DeepEqual(*startingConfig, newConfig) {
		// nothing to do
		return nil
	}

	if startingConfig.CurrentContext != newConfig.CurrentContext {
		if err := writeCurrentContext(configAccess, newConfig.CurrentContext); err != nil {
			return err
		}
	}

	// seenConfigs stores a map of config source filenames to computed config objects
	seenConfigs := map[string]*cmdutils.Config{}

	for key, context := range newConfig.Contexts {
		startingContext, exists := startingConfig.Contexts[key]
		if !reflect.DeepEqual(context, startingContext) || !exists {
			destinationFile := configAccess.GetDefaultFilename()
			// we only obtain a fresh config object from its source file
			// if we have not seen it already - this prevents us from
			// reading and writing to the same number of files repeatedly
			// when multiple / all contexts share the same destination file.
			configToWrite, seen := seenConfigs[destinationFile]
			if !seen {
				var err error
				configToWrite, err = getConfigFromFile(destinationFile)
				if err != nil {
					return err
				}
				seenConfigs[destinationFile] = configToWrite
			}

			configToWrite.Contexts[key] = context
		}
	}

	// actually persist config object changes
	for destinationFile, configToWrite := range seenConfigs {
		if err := WriteToFile(*configToWrite, destinationFile); err != nil {
			return err
		}
	}

	for key, authInfo := range newConfig.AuthInfos {
		startingAuthInfo, exists := startingConfig.AuthInfos[key]
		if !reflect.DeepEqual(authInfo, startingAuthInfo) || !exists {
			destinationFile := authInfo.LocationOfOrigin
			if len(destinationFile) == 0 {
				destinationFile = configAccess.GetDefaultFilename()
			}

			configToWrite, err := getConfigFromFile(destinationFile)
			if err != nil {
				return err
			}
			t := *authInfo
			configToWrite.AuthInfos[key] = &t
			configToWrite.AuthInfos[key].LocationOfOrigin = destinationFile
			//if relativizePaths {
			//	if err := RelativizeAuthInfoLocalPaths(configToWrite.AuthInfos[key]); err != nil {
			//		return err
			//	}
			//}

			if err := WriteToFile(*configToWrite, destinationFile); err != nil {
				return err
			}
		}
	}

	for key := range startingConfig.Contexts {
		if _, exists := newConfig.Contexts[key]; !exists {
			destinationFile := configAccess.GetDefaultFilename()

			configToWrite, err := getConfigFromFile(destinationFile)
			if err != nil {
				return err
			}
			delete(configToWrite.Contexts, key)

			if err := WriteToFile(*configToWrite, destinationFile); err != nil {
				return err
			}
		}
	}

	for key, authInfo := range startingConfig.AuthInfos {
		if _, exists := newConfig.AuthInfos[key]; !exists {
			destinationFile := authInfo.LocationOfOrigin
			if len(destinationFile) == 0 {
				destinationFile = configAccess.GetDefaultFilename()
			}

			configToWrite, err := getConfigFromFile(destinationFile)
			if err != nil {
				return err
			}
			delete(configToWrite.AuthInfos, key)

			if err := WriteToFile(*configToWrite, destinationFile); err != nil {
				return err
			}
		}
	}

	return nil
}

// writeCurrentContext takes three possible paths.
// If newCurrentContext is the same as the startingConfig's current context, then we exit.
// If newCurrentContext has a value, then that value is written into the default destination file.
// If newCurrentContext is empty, then we find the config file that is setting the CurrentContext and
// clear the value from that file
func writeCurrentContext(configAccess ConfigAccess, newCurrentContext string) error {
	if startingConfig, err := configAccess.GetStartingConfig(); err != nil {
		return err
	} else if startingConfig.CurrentContext == newCurrentContext {
		return nil
	}

	if len(newCurrentContext) > 0 {
		destinationFile := configAccess.GetDefaultFilename()
		config, err := getConfigFromFile(destinationFile)
		if err != nil {
			return err
		}
		config.CurrentContext = newCurrentContext

		if err := WriteToFile(*config, destinationFile); err != nil {
			return err
		}

		return nil
	}

	// we're supposed to be clearing the current context.  We need to find the first spot in the chain that
	// is setting it and clear it
	for _, file := range configAccess.GetLoadingPrecedence() {
		if _, err := os.Stat(file); err == nil {
			currConfig, err := getConfigFromFile(file)
			if err != nil {
				return err
			}

			if len(currConfig.CurrentContext) > 0 {
				currConfig.CurrentContext = newCurrentContext
				if err := WriteToFile(*currConfig, file); err != nil {
					return err
				}

				return nil
			}
		}
	}

	return errors.New("no config found to write context")
}

// getConfigFromFile tries to read a pulsarconfig file and if it can't, returns an error.
// One exception, missing files result in empty configs, not an error.
func getConfigFromFile(filename string) (*cmdutils.Config, error) {
	config, err := LoadFromFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if config == nil {
		config = cmdutils.NewConfig()
	}
	return config, nil
}
