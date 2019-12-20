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
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper/bkdata"
	"github.com/stretchr/testify/assert"
)

func TestGetCmd(t *testing.T) {
	var result map[int64]bkdata.LedgerMetadata
	args := []string{"get", "0"}
	out, execErr, nameErr, err := testLedgerCommands(getCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, nameErr)
	assert.Nil(t, execErr)

	json.Unmarshal(out.Bytes(), &result)

	assert.Equal(t, 1, len(result))
	meta := result[0]
	assert.Equal(t, false, meta.StoreCtime)
	assert.Equal(t, false, meta.HasPassword)
	assert.Equal(t, 3, meta.MetadataFormatVersion)
	assert.Equal(t, 1, meta.Ensemble)
	assert.Equal(t, 1, meta.WriteQuorum)
	assert.Equal(t, 1, meta.AckQuorum)
	assert.Equal(t, int64(1000), meta.Length)
	assert.Equal(t, int64(9), meta.LastEntryID)
	assert.Equal(t, "CLOSED", meta.State)
	assert.Equal(t, "CRC32C", meta.DigestType)
	assert.Equal(t, 1, len(meta.Ensembles))
	assert.Equal(t, "", string(meta.Password))
}

func TestGetNonExistentLedger(t *testing.T) {
	args := []string{"get", "10"}
	_, execErr, nameErr, err := testLedgerCommands(getCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, nameErr)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 500 reason: Internal Server Error", execErr.Error())
}

func TestGetArgError(t *testing.T) {
	args := []string{"get"}
	_, _, nameErr, _ := testLedgerCommands(getCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the ledger id is not specified or the ledger id is specified more than one",
		nameErr.Error())

	args = []string{"get", "a"}
	_, execErr, _, _ := testLedgerCommands(getCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid ledger id a", execErr.Error())

	args = []string{"get", "--", "-1"}
	_, execErr, _, _ = testLedgerCommands(getCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid ledger id -1", execErr.Error())
}
