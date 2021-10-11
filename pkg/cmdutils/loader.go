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

package cmdutils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

const (
	RecommendedConfigPathEnvVar = "PULSARCONFIG"
	RecommendedHomeDir          = ".config"
	RecommendedFileName         = "pulsar/config"
)

var (
	OldRecommendedHomeFile = filepath.Join(utils.HomeDir(), "/.pulsar/.pulsarconfig")
	RecommendedHomeFile    = filepath.Join(utils.HomeDir(), RecommendedHomeDir, RecommendedFileName)
)

type ClientConfigLoader interface {
	ConfigAccess
	// Load returns the latest config
	Load() (*Config, error)
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
	migrationRules map[string]string
}

// ClientConfigLoadingRules implements the ClientConfigLoader interface.
var _ ClientConfigLoader = &ClientConfigLoadingRules{}

// NewDefaultClientConfigLoadingRules returns a ClientConfigLoadingRules object with default fields filled in.
func NewDefaultClientConfigLoadingRules() *ClientConfigLoadingRules {
	return NewClientConfigLoadingRules(RecommendedHomeFile, map[string]string{
		RecommendedHomeFile: OldRecommendedHomeFile,
	})
}

// NewClientConfigLoadingRules returns a ClientConfigLoadingRules object.
func NewClientConfigLoadingRules(configPath string, migrationRules map[string]string) *ClientConfigLoadingRules {
	if configPath == "" {
		configPath = RecommendedHomeFile
	}

	var chain []string
	envVarFiles := os.Getenv(RecommendedConfigPathEnvVar)
	if len(envVarFiles) != 0 {
		logger.Debug("found the configuration path %s from %s environment variable", envVarFiles,
			RecommendedConfigPathEnvVar)
		fileList := filepath.SplitList(envVarFiles)
		// prevent the same path load multiple times
		chain = append(chain, deduplicate(fileList)...)
	} else {
		chain = append(chain, configPath)
	}

	return &ClientConfigLoadingRules{
		Precedence:     chain,
		migrationRules: migrationRules,
	}
}

// Load starts by running the MigrationRules and then
// takes the loading rules and returns a Config object based on following rules.
// Find the first valid path from the Precedence slice and load the path as the configuration.
func (rules *ClientConfigLoadingRules) Load() (*Config, error) {
	if err := rules.Migrate(); err != nil {
		return nil, err
	}

	pulsarConfigFiles := []string{}
	pulsarConfigFiles = append(pulsarConfigFiles, rules.Precedence...)

	for _, filename := range pulsarConfigFiles {
		logger.Debug("stat the config file %s", filename)

		exists, err := fileExists(filename)
		if err != nil {
			logger.Debug(
				errors.Wrapf(err, "failed to stat the config file %s", filename).Error())
			continue
		}

		if !exists {
			logger.Warning("config file %s does not exist")
			continue
		}

		return LoadFromFile(filename)
	}

	return nil, errors.New("no config file loaded")
}

// Migrate uses the MigrationRules map. If a destination file is not present, then the source file is checked.
// If the source file is present, then it is copied to the destination file BEFORE any further loading happens.
func (rules *ClientConfigLoadingRules) Migrate() error {
	// migrates the old configuration file.
	for destination, source := range rules.migrationRules {
		exists, err := fileExists(source)
		if err != nil || !exists {
			continue
		}

		exists, err = fileExists(destination)
		if err != nil || !exists {
			continue
		}

		in, err := os.Open(source)
		if err != nil {
			return err
		}

		out, err := os.Create(destination)
		if err != nil {
			_ = in.Close()
			return err
		}

		_, err = io.Copy(out, in)
		_ = out.Close()
		_ = in.Close()

		if err != nil {
			return err
		}
	}

	configPath := rules.GetCurrentConfigFilename()
	defaultConfig := NewConfig()

	// creates a default configuration file if the configuration file doesn't exist.
	if configPath == "" {
		err := errors.New("cannot migrate the configuration to local file, " +
			"if the " + RecommendedConfigPathEnvVar + "environment variable is set, " +
			"ensure that the path is readable and writeable")

		if len(rules.GetLoadingPrecedence()) == 0 {
			return err
		}

		for _, path := range rules.GetLoadingPrecedence() {
			err := WriteToFile(*defaultConfig, path)
			if err == nil {
				return nil
			}
		}

		return err
	}

	exists, err := fileExists(configPath)
	if err != nil {
		return err
	}
	if !exists {
		err = WriteToFile(*defaultConfig, configPath)
		return err
	}

	return nil
}

// GetLoadingPrecedence implements ConfigAccess
func (rules *ClientConfigLoadingRules) GetLoadingPrecedence() []string {
	return rules.Precedence
}

// GetStartingConfig implements ConfigAccess
func (rules *ClientConfigLoadingRules) GetStartingConfig() (*Config, error) {
	return rules.Load()
}

// GetCurrentConfigFilename implements ConfigAccess
func (rules *ClientConfigLoadingRules) GetCurrentConfigFilename() string {
	for _, filename := range rules.GetLoadingPrecedence() {
		if _, err := os.Stat(filename); err == nil {
			return filename
		}
	}
	return ""
}

// LoadFromFile takes a filename and deserializes the contents into Config object
func LoadFromFile(filename string) (*Config, error) {
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
		config.AuthInfos[key] = obj
	}

	for key, obj := range config.Contexts {
		config.Contexts[key] = obj
	}

	if config.AuthInfos == nil {
		config.AuthInfos = map[string]*AuthInfo{}
	}

	if config.Contexts == nil {
		config.Contexts = map[string]*Context{}
	}

	return config, nil
}

// Load takes a byte slice and deserializes the contents into Config object.
// Encapsulates deserialization without assuming the source is a file.
func Load(data []byte) (config *Config, err error) {
	config = NewConfig()
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
func WriteToFile(config Config, filename string) error {
	content, err := toYaml(config)
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

// toYaml serializes the config to yaml.
// Encapsulates serialization without assuming the destination is a file.
func toYaml(config Config) ([]byte, error) {
	return yaml.Marshal(&config)
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

// fileExists checks whether the given path exists and is a file.
func fileExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil {
		return !f.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
