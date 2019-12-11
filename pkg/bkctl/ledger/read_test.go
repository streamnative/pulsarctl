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

package ledger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCmd(t *testing.T)  {
	args := []string{"read", "1"}
	_, execErr, nameErr, err := testLedgerCommands(readCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, nameErr)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 500 reason: Unknown pulsar error", execErr.Error())
}

func TestReadArgError(t *testing.T) {
	args := []string{"read"}
	_, _, nameErr, _ := testLedgerCommands(readCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the ledger id is not specified or the ledger id is specified more than one",
		nameErr.Error())

	args = []string{"read", "a"}
	_, execErr, _, _ := testLedgerCommands(readCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid ledger id a", execErr.Error())

	args = []string{"read", "--", "-1"}
	_, execErr, _, _ = testLedgerCommands(readCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid ledger id -1", execErr.Error())

	args = []string{"read", "--start", "-2", "1"}
	_, execErr, _, _ = testLedgerCommands(readCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid start ledger id -2", execErr.Error())

	args = []string{"read", "--end", "-2", "1"}
	_, execErr, _, _ = testLedgerCommands(readCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid end ledger id -2", execErr.Error())
}
