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
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

const (
	RecommendedConfigPathEnvVar = "PULSARCONFIG"
	RecommendedHomeDir          = ".config"
	RecommendedFileName         = "pulsar/config"
)

var (
	RecommendedConfigDir = path.Join(utils.HomeDir(), RecommendedHomeDir)
	RecommendedHomeFile  = path.Join(RecommendedConfigDir, RecommendedFileName)
)

// currentMigrationRules returns a map that holds the history of recommended home directories used in previous versions.
// Any future changes to RecommendedHomeFile and related are expected to add a migration rule here, in order to make
// sure existing config files are migrated to their new locations properly.
func currentMigrationRules() map[string]string {
	oldRecommendedHomeFile := path.Join(os.Getenv("HOME"), "/.pulsar/.pulsarconfig")
	oldRecommendedWindowsHomeFile := path.Join(os.Getenv("HOME"), RecommendedHomeDir, RecommendedFileName)

	migrationRules := map[string]string{}
	migrationRules[RecommendedHomeFile] = oldRecommendedHomeFile
	if runtime.GOOS == "windows" {
		migrationRules[RecommendedHomeFile] = oldRecommendedWindowsHomeFile
	}
	return migrationRules
}

type ClientConfigLoader interface {
	ConfigAccess
	// Load returns the latest config
	Load() (*cmdutils.Config, error)
}

// ClientConfigLoadingRules is an ExplicitPath and string slice of specific locations that are used
// for merging together a Config
// Callers can put the chain together however they want, but we'd recommend:
// EnvVarPathFiles if set (a list of files if set) OR the HomeDirectoryPath
// ExplicitPath is special, because if a user specifically requests a certain file be used and error
// is reported if this file is not present
type ClientConfigLoadingRules struct {
	Precedence []string

	// MigrationRules is a map of destination files to source files.  If a destination file is not present,
	// then the source file is checked.
	// If the source file is present, then it is copied to the destination file BEFORE any further loading happens.
	MigrationRules map[string]string

	// DoNotResolvePaths indicates whether or not to resolve paths with respect to the originating files.
	// This is phrased as a negative so that a default object that doesn't set this will usually get
	// the behavior it wants.
	DoNotResolvePaths bool

	// DefaultClientConfig is an optional field indicating what rules to use to calculate a default configuration.
	// This should match the overrides passed in to ClientConfig loader.
	DefaultClientConfig ClientConfig

	// WarnIfAllMissing indicates whether the configuration files pointed by pulsarCONFIG environment
	// variable are present or not.
	// In case of missing files, it warns the user about the missing files.
	WarnIfAllMissing bool
}

// ClientConfigLoadingRules implements the ClientConfigLoader interface.
var _ ClientConfigLoader = &ClientConfigLoadingRules{}

// NewDefaultClientConfigLoadingRules returns a ClientConfigLoadingRules object with default fields filled in.
// You are not required to use this constructor
func NewDefaultClientConfigLoadingRules() *ClientConfigLoadingRules {
	chain := []string{}
	warnIfAllMissing := false

	envVarFiles := os.Getenv(RecommendedConfigPathEnvVar)
	if len(envVarFiles) != 0 {
		fileList := filepath.SplitList(envVarFiles)
		// prevent the same path load multiple times
		chain = append(chain, deduplicate(fileList)...)
		warnIfAllMissing = true

	} else {
		chain = append(chain, RecommendedHomeFile)
	}

	return &ClientConfigLoadingRules{
		Precedence:       chain,
		MigrationRules:   currentMigrationRules(),
		WarnIfAllMissing: warnIfAllMissing,
	}
}

// Load starts by running the MigrationRules and then
// takes the loading rules and returns a Config object based on following rules.
//
//	if the ExplicitPath, return the unmerged explicit file
//	Otherwise, return a merged config based on the Precedence slice
//
// A missing ExplicitPath file produces an error. Empty filenames or other missing files are ignored.
// Read errors or files with non-deserializable content produce errors.
// The first file to set a particular map key wins and map key's value is never changed.
// BUT, if you set a struct value that is NOT contained inside of map, the value WILL be changed.
// This results in some odd looking logic to merge in one direction, merge in the other, and then merge the two.
// It also means that if two files specify a "red-user", only values from the first file's red-user are used.  Even
// non-conflicting entries from the second file's "red-user" are discarded.
// Relative paths inside of the .pulsarconfig files are resolved against the .pulsarconfig file's parent folder
// and only absolute file paths are returned.
func (rules *ClientConfigLoadingRules) Load() (*cmdutils.Config, error) {
	if err := rules.Migrate(); err != nil {
		return nil, err
	}

	missingList := []string{}
	pulsarConfigFiles := []string{}
	pulsarConfigFiles = append(pulsarConfigFiles, rules.Precedence...)
	pulsarconfigs := []*cmdutils.Config{}
	// read and cache the config files so that we only look at them once
	for _, filename := range pulsarConfigFiles {
		if len(filename) == 0 {
			// no work to do
			continue
		}
		config, err := LoadFromFile(filename)
		if os.IsNotExist(err) {
			// skip missing files
			// Add to the missing list to produce a warning
			missingList = append(missingList, filename)
			continue
		}

		if err != nil {
			fmt.Println(fmt.Errorf("error loading config file \"%s\": %v", filename, err))
			continue
		}

		pulsarconfigs = append(pulsarconfigs, config)
	}

	if rules.WarnIfAllMissing && len(missingList) > 0 && len(pulsarconfigs) == 0 {
		fmt.Printf("Config not found: %s", strings.Join(missingList, ", "))
	}

	// first merge all of our maps
	mapConfig := cmdutils.NewConfig()

	for _, pulsarconfig := range pulsarconfigs {
		mergo.MergeWithOverwrite(mapConfig, pulsarconfig)
	}

	// merge all of the struct values in the reverse order so that priority is given correctly
	// errors are not added to the list the second time
	nonMapConfig := cmdutils.NewConfig()
	for i := len(pulsarconfigs) - 1; i >= 0; i-- {
		pulsarconfig := pulsarconfigs[i]
		mergo.MergeWithOverwrite(nonMapConfig, pulsarconfig)
	}

	// since values are overwritten, but maps values are not, we can merge the non-map config on top of the map config and
	// get the values we expect.
	config := cmdutils.NewConfig()
	mergo.MergeWithOverwrite(config, mapConfig)
	mergo.MergeWithOverwrite(config, nonMapConfig)

	return config, nil
}

// Migrate uses the MigrationRules map.  If a destination file is not present, then the source file is checked.
// If the source file is present, then it is copied to the destination file BEFORE any further loading happens.
func (rules *ClientConfigLoadingRules) Migrate() error {
	if rules.MigrationRules == nil {
		return nil
	}

	for destination, source := range rules.MigrationRules {
		_, err := os.Stat(destination)
		if err == nil || os.IsPermission(err) {
			// if the destination already exists, do nothing and if we can't access the file, skip it
			continue
		} else if !os.IsNotExist(err) {
			// if we had an error other than non-existence, fail
			return err
		}

		if sourceInfo, err := os.Stat(source); err != nil {
			if os.IsNotExist(err) || os.IsPermission(err) {
				// if the source file doesn't exist or we can't access it, there's no work to do.
				continue
			}

			// if we had an error other than non-existence, fail
			return err
		} else if sourceInfo.IsDir() {
			return fmt.Errorf("cannot migrate %v to %v because it is a directory", source, destination)
		}

		in, err := os.Open(source)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.Create(destination)
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err = io.Copy(out, in); err != nil {
			return err
		}
	}

	return nil
}

// GetLoadingPrecedence implements ConfigAccess
func (rules *ClientConfigLoadingRules) GetLoadingPrecedence() []string {
	return rules.Precedence
}

// GetStartingConfig implements ConfigAccess
func (rules *ClientConfigLoadingRules) GetStartingConfig() (*cmdutils.Config, error) {
	clientConfig := NewNonInteractiveDeferredLoadingClientConfig(rules, &cmdutils.ConfigOverrides{})
	rawConfig, err := clientConfig.RawConfig()
	if os.IsNotExist(err) {
		return cmdutils.NewConfig(), nil
	}
	if err != nil {
		return nil, err
	}

	return &rawConfig, nil
}

// GetDefaultFilename implements ConfigAccess
func (rules *ClientConfigLoadingRules) GetDefaultFilename() string {
	// first existing file from precedence.
	for _, filename := range rules.GetLoadingPrecedence() {
		if _, err := os.Stat(filename); err == nil {
			return filename
		}
	}
	// If none exists, use the first from precedence.
	if len(rules.Precedence) > 0 {
		return rules.Precedence[0]
	}
	return ""
}

// LoadFromFile takes a filename and deserializes the contents into Config object
func LoadFromFile(filename string) (*cmdutils.Config, error) {
	pulsarconfigBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config, err := Load(pulsarconfigBytes)
	if err != nil {
		return nil, err
	}

	// set LocationOfOrigin on every BrokerServiceURL, User, and Context
	for key, obj := range config.AuthInfos {
		obj.LocationOfOrigin = filename
		config.AuthInfos[key] = obj
	}

	for key, obj := range config.Contexts {
		config.Contexts[key] = obj
	}

	if config.AuthInfos == nil {
		config.AuthInfos = map[string]*cmdutils.AuthInfo{}
	}

	if config.Contexts == nil {
		config.Contexts = map[string]*cmdutils.Context{}
	}

	return config, nil
}

// Load takes a byte slice and deserializes the contents into Config object.
// Encapsulates deserialization without assuming the source is a file.
func Load(data []byte) (config *cmdutils.Config, err error) {
	config = cmdutils.NewConfig()
	// if there's no data in a file, return the default object instead of failing (DecodeInto reject empty input)
	if len(data) == 0 {
		return config, nil
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// WriteToFile serializes the config to yaml and writes it out to a file.  If not present,
// it creates the file with the mode 0600.  If it is present it stomps the contents
func WriteToFile(config cmdutils.Config, filename string) error {
	content, err := Write(config)
	if err != nil {
		return err
	}
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(filename, content, 0600); err != nil {
		return err
	}
	return nil
}

func lockFile(filename string) error {
	// TODO: find a way to do this with actual file locks. Will
	// probably need separate solution for windows and Linux.

	// Make sure the dir exists before we try to create a lock file.
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	f, err := os.OpenFile(lockName(filename), os.O_CREATE|os.O_EXCL, 0)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

func unlockFile(filename string) error {
	return os.Remove(lockName(filename))
}

func lockName(filename string) string {
	return filename + ".lock"
}

// Write serializes the config to yaml.
// Encapsulates serialization without assuming the destination is a file.
func Write(config cmdutils.Config) ([]byte, error) {
	return yaml.Marshal(&config)
}

func (rules ClientConfigLoadingRules) ResolvePaths() bool {
	return !rules.DoNotResolvePaths
}

// deduplicate removes any duplicated values and returns a new slice, keeping the order unchanged
func deduplicate(s []string) []string {
	encountered := map[string]bool{}
	ret := make([]string, 0)
	for i := range s {
		if encountered[s[i]] {
			continue
		}
		encountered[s[i]] = true
		ret = append(ret, s[i])
	}
	return ret
}
