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

func TestModifyConfigWithDefaultLoader(t *testing.T) {
	_, exists := os.LookupEnv("CI")
	if !exists {
		t.Skip("skip this test, it will modify the your local file")
		return
	}

	// cleanup
	defer func() {
		_ = os.RemoveAll(filepath.Dir(RecommendedHomeFile))
	}()

	data, err := ioutil.ReadFile(filepath.Join("testdata", "config"))
	assert.NoError(t, err)

	// write configuration file to $HOME/.config/pulsar/config.
	_ = os.Mkdir(filepath.Dir(RecommendedHomeFile), 0755)
	err = ioutil.WriteFile(RecommendedHomeFile, data, 0644)
	assert.NoError(t, err)

	loader := NewDefaultClientConfigLoadingRules()
	testModifyConfig(loader, t)
}

func TestModifyConfigWithCustomLoader(t *testing.T) {
	// write config file to temp file.
	f, err := ioutil.TempFile("", "pulsarctl")
	assert.NoError(t, err)
	defer os.RemoveAll(f.Name())

	data, err := ioutil.ReadFile(filepath.Join("testdata", "config"))
	assert.NoError(t, err)
	err = ioutil.WriteFile(f.Name(), data, 0644)
	assert.NoError(t, err)

	loader := NewClientConfigLoadingRules(f.Name(), nil)

	testModifyConfig(loader, t)
}

func testModifyConfig(loader ClientConfigLoader, t *testing.T) {
	config, err := loader.Load()
	assert.NoError(t, err)

	expectConfigFromTestdataConfigFile(config, t)

	newConfig := NewConfig()
	err = ModifyConfig(loader, *newConfig)
	assert.NoError(t, err)

	config, err = loader.Load()
	assert.NoError(t, err)
	assert.NotNil(t, config)
}
