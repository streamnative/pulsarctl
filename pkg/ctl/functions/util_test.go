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

package functions

import (
	"encoding/json"
	"os"
	"path"
	"strings"
	"testing"

	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestMarshalSecretsAndUserConfig(t *testing.T) {
	configFile := path.Join(ResourceDir(), "example-function-config.yaml")
	yamlFile, err := os.ReadFile(configFile)
	assert.Nil(t, err)

	funcConf := &util.FunctionConfig{}
	err = yaml.Unmarshal(yamlFile, funcConf)
	assert.Nil(t, err)

	// should report error when marshal directly
	_, err = json.Marshal(funcConf)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "json: unsupported type: map[interface {}]interface {}"))

	// should not report error when marshal after format
	formatFuncConf(funcConf)
	jsonData, err := json.Marshal(funcConf)
	assert.Nil(t, err)
	expectedResult := `{"cleanupSubscription":false,"retainOrdering":false,"retainKeyOrdering":false,` +
		`"forwardSourceMessageProperty":false,"autoAck":true,"parallelism":1,"output":"test_result","tenant":"public",` +
		`"namespace":"default","name":"example-functions-config","className":"org.apache.pulsar.functions.api.examples.` +
		`ExclamationFunction","inputs":["test_src"],"userConfig":{"arrayMapKey":[{"key":"value2","path":"config2"}],` +
		`"mapKey":{"key":"value","path":"config"},"stringKey":"stringValue"},"secrets":{"arrayMapKey":` +
		`[{"key":"password2","path":"secret2"}],"mapKey":{"key":"password","path":"secret"},"stringKey":"stringSecret"},` +
		`"exposePulsarAdminClientEnabled":false,"skipToLatest":false}`
	assert.Equal(t, expectedResult, string(jsonData))
}
