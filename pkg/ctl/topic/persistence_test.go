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

package topic

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestPersistence(t *testing.T) {
	topicName := "persistent://public/default/test-persistence-topic-10"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	setArgs := []string{"set-persistence", topicName, "-e", "5", "-w", "4", "-a", "3", "-r", "2.2"}
	setOut, execErr, _, _ := TestTopicCommands(SetPersistenceCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set persistence successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	getArgs := []string{"get-persistence", topicName}
	getOut, execErr, _, _ := TestTopicCommands(GetPersistenceCmd, getArgs)
	var persistenceData utils.PersistenceData
	err := json.Unmarshal(getOut.Bytes(), &persistenceData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, persistenceData.BookkeeperEnsemble, int64(5))
	assert.Equal(t, persistenceData.BookkeeperWriteQuorum, int64(4))
	assert.Equal(t, persistenceData.BookkeeperAckQuorum, int64(3))
	assert.Equal(t, persistenceData.ManagedLedgerMaxMarkDeleteRate, float64(2.2))

	setArgs = []string{"remove-persistence", topicName}
	setOut, execErr, _, _ = TestTopicCommands(RemovePersistenceCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Remove persistence successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	getArgs = []string{"get-persistence", topicName}
	getOut, execErr, _, _ = TestTopicCommands(GetPersistenceCmd, getArgs)
	err = json.Unmarshal(getOut.Bytes(), &persistenceData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, persistenceData.BookkeeperEnsemble, int64(0))
	assert.Equal(t, persistenceData.BookkeeperWriteQuorum, int64(0))
	assert.Equal(t, persistenceData.BookkeeperAckQuorum, int64(0))
	assert.Equal(t, persistenceData.ManagedLedgerMaxMarkDeleteRate, float64(0))

	// test value
	setArgs = []string{"set-persistence", topicName, "-e", "1", "-w", "4", "-a", "3", "-r", "2.2"}
	_, execErr, _, _ = TestTopicCommands(SetPersistenceCmd, setArgs)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "code: 400 reason: Bookkeeper Ensemble (1) >= WriteQuorum (4) >= AckQuoru (3)")
}
