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

package namespace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveMessageTTLCmd(t *testing.T) {
	ns := "public/test-remove-message-ttl"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-message-ttl", ns}
	initialOut, execErr, _, _ := TestNamespaceCommands(getMessageTTL, args)
	assert.Nil(t, execErr)

	args = []string{"set-message-ttl", ns, "-t", "123"}
	_, execErr, _, _ = TestNamespaceCommands(setMessageTTL, args)
	assert.Nil(t, execErr)

	args = []string{"remove-message-ttl", ns}
	removeOut, execErr, _, _ := TestNamespaceCommands(removeMessageTTL, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Removed message TTL successfully for [%s]\n", ns), removeOut.String())

	args = []string{"get-message-ttl", ns}
	currentOut, execErr, _, _ := TestNamespaceCommands(getMessageTTL, args)
	assert.Nil(t, execErr)
	assert.Equal(t, initialOut.String(), currentOut.String())
}

func TestRemoveRetentionCmd(t *testing.T) {
	ns := "public/test-remove-retention"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-retention", ns}
	initialOut, execErr, _, _ := TestNamespaceCommands(getRetention, args)
	assert.Nil(t, execErr)

	args = []string{"set-retention", ns, "--time", "10m", "--size", "10M"}
	_, execErr, _, _ = TestNamespaceCommands(setRetention, args)
	assert.Nil(t, execErr)

	args = []string{"remove-retention", ns}
	removeOut, execErr, _, _ := TestNamespaceCommands(removeRetention, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Removed retention successfully for [%s]\n", ns), removeOut.String())

	args = []string{"get-retention", ns}
	currentOut, execErr, _, _ := TestNamespaceCommands(getRetention, args)
	assert.Nil(t, execErr)
	assert.Equal(t, initialOut.String(), currentOut.String())
}

func TestRemovePersistenceCmd(t *testing.T) {
	ns := "public/test-remove-persistence"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-persistence", ns}
	initialOut, execErr, _, _ := TestNamespaceCommands(getPersistence, args)
	assert.Nil(t, execErr)

	args = []string{"set-persistence", ns,
		"--ensemble-size", "2",
		"--write-quorum-size", "2",
		"--ack-quorum-size", "2",
		"--ml-mark-delete-max-rate", "1.5",
	}
	_, execErr, _, _ = TestNamespaceCommands(setPersistence, args)
	assert.Nil(t, execErr)

	args = []string{"remove-persistence", ns}
	removeOut, execErr, _, _ := TestNamespaceCommands(removePersistence, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Removed persistence successfully for [%s]\n", ns), removeOut.String())

	args = []string{"get-persistence", ns}
	currentOut, execErr, _, _ := TestNamespaceCommands(getPersistence, args)
	assert.Nil(t, execErr)
	assert.Equal(t, initialOut.String(), currentOut.String())
}

func TestRemoveMaxConsumersAndProducersCmds(t *testing.T) {
	ns := "public/test-remove-max-consumers-and-producers"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-max-consumers-per-topic", ns}
	initialConsumersPerTopicOut, execErr, _, _ := TestNamespaceCommands(GetMaxConsumersPerTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-max-consumers-per-topic", "--size", "10", ns}
	_, execErr, _, _ = TestNamespaceCommands(SetMaxConsumersPerTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"remove-max-consumers-per-topic", ns}
	removeConsumersPerTopicOut, execErr, _, _ := TestNamespaceCommands(RemoveMaxConsumersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Removed max consumers per topic successfully for [%s]\n", ns),
		removeConsumersPerTopicOut.String())

	args = []string{"get-max-consumers-per-topic", ns}
	currentConsumersPerTopicOut, execErr, _, _ := TestNamespaceCommands(GetMaxConsumersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, initialConsumersPerTopicOut.String(), currentConsumersPerTopicOut.String())

	args = []string{"get-max-consumers-per-subscription", ns}
	initialConsumersPerSubOut, execErr, _, _ := TestNamespaceCommands(GetMaxConsumersPerSubscriptionCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-max-consumers-per-subscription", "--size", "9", ns}
	_, execErr, _, _ = TestNamespaceCommands(SetMaxConsumersPerSubscriptionCmd, args)
	assert.Nil(t, execErr)

	args = []string{"remove-max-consumers-per-subscription", ns}
	removeConsumersPerSubOut, execErr, _, _ := TestNamespaceCommands(RemoveMaxConsumersPerSubscriptionCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Removed max consumers per subscription successfully for [%s]\n", ns),
		removeConsumersPerSubOut.String())

	args = []string{"get-max-consumers-per-subscription", ns}
	currentConsumersPerSubOut, execErr, _, _ := TestNamespaceCommands(GetMaxConsumersPerSubscriptionCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, initialConsumersPerSubOut.String(), currentConsumersPerSubOut.String())

	args = []string{"get-max-producers-per-topic", ns}
	initialProducersPerTopicOut, execErr, _, _ := TestNamespaceCommands(GetMaxProducersPerTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-max-producers-per-topic", "--size", "8", ns}
	_, execErr, _, _ = TestNamespaceCommands(SetMaxProducersPerTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"remove-max-producers-per-topic", ns}
	removeProducersPerTopicOut, execErr, _, _ := TestNamespaceCommands(RemoveMaxProducersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Removed max producers per topic successfully for [%s]\n", ns),
		removeProducersPerTopicOut.String())

	args = []string{"get-max-producers-per-topic", ns}
	currentProducersPerTopicOut, execErr, _, _ := TestNamespaceCommands(GetMaxProducersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, initialProducersPerTopicOut.String(), currentProducersPerTopicOut.String())
}
