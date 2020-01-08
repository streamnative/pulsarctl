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
	"io"
	"sync"
)

// DeferredLoadingClientConfig is a ClientConfig interface that is backed by a client config loader.
type DeferredLoadingClientConfig struct {
	loader         ClientConfigLoader
	overrides      *ConfigOverrides
	fallbackReader io.Reader

	clientConfig ClientConfig
	loadingLock  sync.Mutex
}

func (config *DeferredLoadingClientConfig) RawConfig() (Config, error) {
	mergedConfig, err := config.createClientConfig()
	if err != nil {
		return Config{}, err
	}

	return mergedConfig.RawConfig()
}

// ConfigAccess implements ClientConfig
func (config *DeferredLoadingClientConfig) ConfigAccess() ConfigAccess {
	return config.loader
}

// NewNonInteractiveDeferredLoadingClientConfig creates a ConfigClientClientConfig using the passed context name
func NewNonInteractiveDeferredLoadingClientConfig(loader ClientConfigLoader, overrides *ConfigOverrides) ClientConfig {
	return &DeferredLoadingClientConfig{loader: loader, overrides: overrides}
}

// DirectClientConfig is a ClientConfig interface that is backed by a clientcmdapi.Config,
// options overrides, and an optional fallbackReader for auth information
type DirectClientConfig struct {
	config         Config
	contextName    string
	overrides      *ConfigOverrides
	fallbackReader io.Reader
	configAccess   ConfigAccess
}

// NewNonInteractiveClientConfig creates a DirectClientConfig using the passed context
// name and does not have a fallback reader for auth information
func NewNonInteractiveClientConfig(config Config, contextName string, overrides *ConfigOverrides,
	configAccess ConfigAccess) ClientConfig {
	return &DirectClientConfig{config, contextName, overrides, nil, configAccess}
}

// NewInteractiveClientConfig creates a DirectClientConfig using the passed context
// name and a reader in case auth information is not provided via files or flags
func NewInteractiveClientConfig(config Config, contextName string, overrides *ConfigOverrides, fallbackReader io.Reader,
	configAccess ConfigAccess) ClientConfig {
	return &DirectClientConfig{config, contextName, overrides, fallbackReader, configAccess}
}

func (config *DirectClientConfig) RawConfig() (Config, error) {
	return config.config, nil
}

// ConfigAccess implements ClientConfig
func (config *DirectClientConfig) ConfigAccess() ConfigAccess {
	return config.configAccess
}

func (config *DeferredLoadingClientConfig) createClientConfig() (ClientConfig, error) {
	if config.clientConfig == nil {
		config.loadingLock.Lock()
		defer config.loadingLock.Unlock()

		if config.clientConfig == nil {
			mergedConfig, err := config.loader.Load()
			if err != nil {
				return nil, err
			}

			var mergedClientConfig ClientConfig
			if config.fallbackReader != nil {
				mergedClientConfig = NewInteractiveClientConfig(*mergedConfig, config.overrides.CurrentContext,
					config.overrides, config.fallbackReader, config.loader)
			} else {
				mergedClientConfig = NewNonInteractiveClientConfig(*mergedConfig, config.overrides.CurrentContext,
					config.overrides, config.loader)
			}

			config.clientConfig = mergedClientConfig
		}
	}

	return config.clientConfig, nil
}
