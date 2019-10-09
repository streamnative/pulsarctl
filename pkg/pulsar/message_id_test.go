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

package pulsar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMessageId(t *testing.T) {
	id, err := ParseMessageID("1:1")
	assert.Nil(t, err)
	assert.Equal(t, MessageID{LedgerID: 1, EntryID: 1, PartitionedIndex: -1}, *id)
}

func TestParseMessageIdErrors(t *testing.T) {
	id, err := ParseMessageID("1;1")
	assert.Nil(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid message id string. 1;1", err.Error())

	id, err = ParseMessageID("a:1")
	assert.Nil(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid ledger id string. a:1", err.Error())

	id, err = ParseMessageID("1:a")
	assert.Nil(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid entry id string. 1:a", err.Error())
}
