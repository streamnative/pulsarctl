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
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientConfigLoaderWithDefaultConfig_Migrate(t *testing.T) {
	_, exists := os.LookupEnv("CI")
	if !exists {
		t.Skip("skip this test, it will modify the your local file")
		return
	}

	// cleanup
	defer func() {
		_ = os.RemoveAll(filepath.Dir(OldRecommendedHomeFile))
		_ = os.RemoveAll(RecommendedHomeFile)
	}()

	data, err := ioutil.ReadFile(filepath.Join("testdata", "config"))
	assert.NoError(t, err)

	// write configuration data to $HOME/.pulsar/.pulsarconfig.
	_ = os.Mkdir(filepath.Dir(OldRecommendedHomeFile), 0755)
	err = ioutil.WriteFile(OldRecommendedHomeFile, data, 0644)
	assert.NoError(t, err)

	loader := NewDefaultClientConfigLoadingRules()

	// migrate configuration from $HOME/.pulsar/.pulsarconfig to $HOME/.config/pulsar/config.
	err = loader.Migrate()
	assert.NoError(t, err)
}

func TestClientConfigLoaderWithDefaultConfig_Load(t *testing.T) {
	_, exists := os.LookupEnv("CI")
	if !exists {
		t.Skip("skip this test, it will modify the your local file")
		return
	}

	// cleanup
	defer func() {
		_ = os.RemoveAll(filepath.Dir(OldRecommendedHomeFile))
		_ = os.RemoveAll(filepath.Dir(RecommendedHomeFile))
	}()

	data, err := ioutil.ReadFile(filepath.Join("testdata", "config"))
	assert.NoError(t, err)

	// write configuration file to $HOME/.config/pulsar/config.
	_ = os.Mkdir(filepath.Dir(RecommendedHomeFile), 0755)
	err = ioutil.WriteFile(RecommendedHomeFile, data, 0644)
	assert.NoError(t, err)

	loader := NewDefaultClientConfigLoadingRules()
	config, err := loader.Load()
	assert.NoError(t, err)

	expectConfigFromTestdataConfigFile(config, t)
}

func TestClientConfigLoaderWithDefaultConfig_Load_When_ConfigFile_Not_Exist(t *testing.T) {
	_, exists := os.LookupEnv("CI")
	if !exists {
		t.Skip("skip this test, it will modify the your local file")
		return
	}

	// cleanup
	defer func() {
		_ = os.RemoveAll(filepath.Dir(RecommendedHomeFile))
	}()

	_ = os.RemoveAll(RecommendedHomeFile)
	loader := NewDefaultClientConfigLoadingRules()
	config, err := loader.Load()
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestClientConfigLoaderWithCustom_Load(t *testing.T) {
	// write config file to temp file.
	f, err := ioutil.TempFile("", "pulsarctl")
	assert.NoError(t, err)
	defer os.RemoveAll(f.Name())

	data, err := ioutil.ReadFile(filepath.Join("testdata", "config"))
	assert.NoError(t, err)
	err = ioutil.WriteFile(f.Name(), data, 0644)
	assert.NoError(t, err)

	loader := NewClientConfigLoadingRules(f.Name(), nil)

	config, err := loader.Load()
	assert.NoError(t, err)

	expectConfigFromTestdataConfigFile(config, t)
}

func expectConfigFromTestdataConfigFile(config *Config, t *testing.T) {
	key := "standalone"
	assert.NotNil(t, config.AuthInfos[key])
	assert.Equal(t, "", config.AuthInfos[key].TLSTrustCertsFilePath)
	assert.Equal(t, false, config.AuthInfos[key].TLSAllowInsecureConnection)
	assert.Equal(t, "", config.AuthInfos[key].Token)
	assert.Equal(t, "", config.AuthInfos[key].TokenFile)
	assert.Equal(t, "https://login.cluster-1.pulsar.demo.local/11969f5b-8b91-4ea8-bd79-f33ede77397a/v2.0",
		config.AuthInfos[key].IssuerEndpoint)
	assert.Equal(t, "d69780eb-0cc7-44fc-b183-08d124d1f712", config.AuthInfos[key].ClientID)
	assert.Equal(t, "api://cluster-1.pulsar.demo.local", config.AuthInfos[key].Audience)
	assert.Equal(t, "api://cluster-1.pulsar.demo.local/.default", config.AuthInfos[key].Scope)
	assert.Equal(t, "", config.AuthInfos[key].KeyFile)

	assert.NotNil(t, config.Contexts[key])
	assert.Equal(t, "http://localhost:8080", config.Contexts[key].BookieServiceURL)
	assert.Equal(t, "http://localhost:8080", config.Contexts[key].BrokerServiceURL)

	assert.Equal(t, key, config.CurrentContext)
}
