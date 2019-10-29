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
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestPersistence(t *testing.T) {
	args := []string{"create", "public/test-persistent-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-persistent-namespace successfully\n")

	setArgs := []string{"set-persistence", "public/test-persistent-namespace",
		"--ensemble-size", "2",
		"--write-quorum-size", "2",
		"--ack-quorum-size", "2",
		"--ml-mark-delete-max-rate", "2.0",
	}
	setOut, execErr, _, _ := TestNamespaceCommands(setPersistence, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set the persistence policies successfully for [public/test-persistent-namespace]\n")

	var persistence pulsar.PersistencePolicies

	getArgs := []string{"get-persistence", "public/test-persistent-namespace"}
	getOut, execErr, _, _ := TestNamespaceCommands(getPersistence, getArgs)
	assert.Nil(t, execErr)
	err = json.Unmarshal(getOut.Bytes(), &persistence)
	assert.Nil(t, err)
	assert.Equal(t, persistence.BookkeeperEnsemble, 2)
	assert.Equal(t, persistence.BookkeeperWriteQuorum, 2)
	assert.Equal(t, persistence.BookkeeperAckQuorum, 2)
	assert.Equal(t, persistence.ManagedLedgerMaxMarkDeleteRate, float64(2))
}

func TestFailurePersistence(t *testing.T) {
	setArgs := []string{"set-persistence", "public/test-persistent-namespace",
		"--ensemble-size", "2",
		"--write-quorum-size", "5",
		"--ack-quorum-size", "2",
		"--ml-mark-delete-max-rate", "2.0",
	}
	_, execErr, _, _ := TestNamespaceCommands(setPersistence, setArgs)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "code: 412 reason: Bookkeeper Ensemble (2) >= WriteQuorum (5) >= AckQuoru (2)")

	setArgs = []string{"set-persistence", "public/test-persistent-namespace",
		"--ensemble-size", "2",
		"--write-quorum-size", "2",
		"--ack-quorum-size", "3",
		"--ml-mark-delete-max-rate", "2.0",
	}
	_, execErr, _, _ = TestNamespaceCommands(setPersistence, setArgs)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "code: 412 reason: Bookkeeper Ensemble (2) >= WriteQuorum (2) >= AckQuoru (3)")
}
