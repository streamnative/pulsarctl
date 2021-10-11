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
	"fmt"
	"reflect"

	"github.com/pkg/errors"

	"github.com/gofrs/flock"
)

type ConfigAccess interface {
	// GetLoadingPrecedence returns the slice of files that should be used for loading and inspecting the config
	GetLoadingPrecedence() []string
	// GetStartingConfig returns the config that subcommands should being operating against.
	// It may or may not be merged depending on loading rules
	GetStartingConfig() (*Config, error)
	// GetCurrentConfigFilename returns the name of the file you should write into (create if necessary),
	// if you're trying to create a new stanza as opposed to updating an existing one.
	GetCurrentConfigFilename() string
}

// ModifyConfig takes a Config object and write filed of Config struct to file
func ModifyConfig(configAccess ConfigAccess, newConfig Config) error {
	configPath := configAccess.GetCurrentConfigFilename()
	if configPath == "" {
		return errors.New("failed to get current configuration file path")
	}

	fileLock := flock.New(configPath)
	locked, err := fileLock.TryLock()
	if err != nil {
		return err
	}
	if !locked {
		return fmt.Errorf("cannot modify the %s configuration file that has been locked by another process",
			configPath)
	}
	defer fileLock.Unlock()

	startingConfig, err := configAccess.GetStartingConfig()
	if err != nil {
		return err
	}

	if reflect.DeepEqual(*startingConfig, newConfig) {
		// nothing to do
		return nil
	}

	return WriteToFile(newConfig, configPath)
}
