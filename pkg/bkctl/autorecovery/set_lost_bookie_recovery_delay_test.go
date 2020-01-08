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

package autorecovery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLostBookieRecoveryDelayArgsErr(t *testing.T) {
	// no args specified
	args := []string{"set-lost-bookie-recovery-delay"}
	_, _, nameErr, err := testAutoRecoveryCommands(setLostBookieRecoveryDelayCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, nameErr)
	assert.Equal(t, "the delay time is not specified or the delay time is specified more than one",
		nameErr.Error())

	// specify more than one args
	args = []string{"set-lost-bookie-recovery-delay", "1", "2"}
	_, _, nameErr, err = testAutoRecoveryCommands(setLostBookieRecoveryDelayCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, nameErr)
	assert.Equal(t, "the delay time is not specified or the delay time is specified more than one",
		nameErr.Error())

	// specify invalid args
	args = []string{"set-lost-bookie-recovery-delay", "a"}
	_, execErr, nameErr, err := testAutoRecoveryCommands(setLostBookieRecoveryDelayCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	if nameErr != nil {
		t.Fatal(nameErr)
	}

	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid delay times a", execErr.Error())

	args = []string{"set-lost-bookie-recovery-delay", "--", "-1"}
	_, execErr, nameErr, err = testAutoRecoveryCommands(setLostBookieRecoveryDelayCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	if nameErr != nil {
		t.Fatal(nameErr)
	}

	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid delay times -1", execErr.Error())
}
