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

package topicpolicies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopicPoliciesNameError(t *testing.T) {
	_, _, nameErr, _ := TestTopicPoliciesCommands(t, GetMessageTTLCmd, []string{"get-message-ttl"})
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestTopicPoliciesRequiredFlag(t *testing.T) {
	_, _, _, err := TestTopicPoliciesCommands(t, SetMessageTTLCmd, []string{"set-message-ttl", "persistent://public/default/test"})
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"ttl\" not set", err.Error())
}

func TestTopicPoliciesEnableDisableValidation(t *testing.T) {
	_, execErr, _, _ := TestTopicPoliciesCommands(t, SetDeduplicationCmd, []string{"set-deduplication", "persistent://public/default/test"})
	assert.NotNil(t, execErr)
	assert.Equal(t, "need to specify either --enable or --disable", execErr.Error())
}

func TestTopicPoliciesSchemaCompatibilityRequiredFlag(t *testing.T) {
	_, _, _, err := TestTopicPoliciesCommands(t, SetSchemaCompatibilityStrategyCmd, []string{
		"set-schema-compatibility-strategy",
		"persistent://public/default/test",
	})
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"compatibility\" not set", err.Error())
}

func TestTopicPoliciesReplicationClustersRequiredFlag(t *testing.T) {
	_, _, _, err := TestTopicPoliciesCommands(t, SetReplicationClustersCmd, []string{
		"set-replication-clusters",
		"persistent://public/default/test",
	})
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"clusters\" not set", err.Error())
}

func TestTopicPoliciesReplicationClustersValidation(t *testing.T) {
	_, execErr, _, _ := TestTopicPoliciesCommands(t, SetReplicationClustersCmd, []string{
		"set-replication-clusters",
		"--clusters", "usw,,use",
		"persistent://public/default/test",
	})
	assert.NotNil(t, execErr)
	assert.Equal(t, "cluster names must be non-empty", execErr.Error())
}

func TestTopicPoliciesSetRetentionHelp(t *testing.T) {
	assert.NotPanics(t, func() {
		_, _, _, err := TestTopicPoliciesCommands(t, SetRetentionCmd, []string{"set-retention", "--help"})
		assert.NoError(t, err)
	})
}
