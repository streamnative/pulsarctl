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

package plugin

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/plugin"
	"github.com/stretchr/testify/assert"
)

func TestListCommand(t *testing.T) {
	plugin.ValidPluginFilenamePrefixes = []string{"plugin-test"}

	args := []string{"list", "-o", "json"}
	out, execErr, err := testPluginCommands(listCmd, args)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, "[]", out.String())

	// create a temp plugin file pulsarctl-foo to test plugin list command
	dir, err := ioutil.TempDir("", "plugins")
	defer os.RemoveAll(dir)
	if err != nil {
		t.Fatal(err)
	}

	pluginFile := filepath.Join(dir, "plugin-test-foo")
	if err := ioutil.WriteFile(pluginFile, []byte{}, 0777); err != nil {
		t.Fatal(err)
	}

	os.Setenv("PATH", dir)

	// test list the pulsarctl-foo plugin

	args = []string{"list", "-o", "json"}
	out, execErr, err = testPluginCommands(listCmd, args)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, "[\n  \"foo\"\n]", out.String())
}
