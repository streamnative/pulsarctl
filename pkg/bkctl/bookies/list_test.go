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

package bookies

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testListData = []struct {
	bookieType   string
	showHostName bool
	outString    string ``
}{
	{"ro", false, "{}\n"},
	{"ro", true, "{}\n"},
	{"rw", false, "3181"},
	{"rw", true, "3181"},
}

func TestListCmd(t *testing.T) {
	for _, data := range testListData {
		t.Logf("test case: %+v\n", data)
		args := []string{"list", data.bookieType}
		if data.showHostName {
			args = append(args, "--show-hostname")
		}
		out, execErr, nameErr, err := testBookiesCommands(args)
		assert.Nil(t, err)
		assert.Nil(t, nameErr)
		assert.Nil(t, execErr)
		assert.True(t, strings.Contains(out.String(), data.outString))
	}
}

func TestListArgError(t *testing.T) {
	args := []string{"list"}
	_, _, nameErr, _ := testBookiesCommands(args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the bookie type is not specified or the bookie type is specified more than one",
		nameErr.Error())

	args = []string{"list", "invalid"}
	_, execErr, _, _ := testBookiesCommands(args)
	assert.NotNil(t, execErr)
	assert.Equal(t, fmt.Sprintf("invalid bookie type %s, the bookie type only can "+
		"be specified as 'rw' or 'ro'", "invalid"), execErr.Error())
}
