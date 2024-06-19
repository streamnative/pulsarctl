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

package context

import (
	"fmt"
	"os"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func TestSetContextCmd(t *testing.T) {
	home := utils.HomeDir()
	path := fmt.Sprintf("%s/.config/pulsar/config", home)
	defer os.Remove(path)

	setArgs := []string{"set", "test-set-context"}
	out, _, err := TestConfigCommands(setContextCmd, setArgs)
	assert.Nil(t, err)
	assert.Equal(t, "Context \"test-set-context\" created.\n", out.String())

	out, _, err = TestConfigCommands(setContextCmd, setArgs)
	assert.Nil(t, err)
	assert.Equal(t, "Context \"test-set-context\" modified.\n", out.String())

	delArgs := []string{"delete", "test-set-context"}
	out, execErr, err := TestConfigCommands(deleteContextCmd, delArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	warnOut := "warning: this removed your active context, " +
		"use \"pulsarctl context use\" to select a different one\n"
	expectedOut := fmt.Sprintf("deleted context test-set-context from %s\n", path)
	assert.Equal(t, warnOut+expectedOut, out.String())
}

func TestOauthConfiguration(t *testing.T) {
	home := utils.HomeDir()
	path := fmt.Sprintf("%s/.config/pulsar/config", home)
	defer os.Remove(path)

	setOauthConfigArgs := []string{"set", "oauth",
		"--issuer-endpoint", "https://test-endpoint",
		"--client-id", "clientid",
		"--audience", "audience",
		"--scope", "profile api://test-endpoint",
	}
	_, execErr, err := TestConfigCommands(setContextCmd, setOauthConfigArgs)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Nil(t, execErr)

	config := cmdutils.LoadFromEnv()
	assert.Equal(t, "https://test-endpoint", config.IssuerEndpoint)
	assert.Equal(t, "clientid", config.ClientID)
	assert.Equal(t, "audience", config.Audience)
	assert.Equal(t, "profile api://test-endpoint", config.Scope)
}

func TestParseOauthConfiguration(t *testing.T) {
	home := utils.HomeDir()
	path := fmt.Sprintf("%s/.config/pulsar/config", home)
	defer os.Remove(path)

	setOauthConfigArgs := []string{"set", "oauth",
		"--auth-params", "{\"audience\":\"audience\",\"issuerUrl\":\"https://test-endpoint\",\"privateKey\":\"/tmp/auth.json\",\"scope\":\"profile api://test-endpoint\",\"clientId\":\"clientid\"}",
	}
	_, execErr, err := TestConfigCommands(setContextCmd, setOauthConfigArgs)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Nil(t, execErr)

	config := cmdutils.LoadFromEnv()
	assert.Equal(t, "https://test-endpoint", config.IssuerEndpoint)
	assert.Equal(t, "clientid", config.ClientID)
	assert.Equal(t, "audience", config.Audience)
	assert.Equal(t, "profile api://test-endpoint", config.Scope)
	assert.Equal(t, "/tmp/auth.json", config.KeyFile)
}

func TestParseWrongFormatOauthConfiguration(t *testing.T) {
	home := utils.HomeDir()
	path := fmt.Sprintf("%s/.config/pulsar/config", home)
	defer os.Remove(path)

	setOauthConfigArgs := []string{"set", "oauth",
		"--auth-params", "wrong_format",
	}
	_, execErr, err := TestConfigCommands(setContextCmd, setOauthConfigArgs)
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Nil(t, execErr)

	config := cmdutils.LoadFromEnv()
	assert.Equal(t, "", config.IssuerEndpoint)
	assert.Equal(t, "", config.ClientID)
	assert.Equal(t, "", config.Audience)
	assert.Equal(t, "", config.Scope)
	assert.Equal(t, "", config.KeyFile)
}
