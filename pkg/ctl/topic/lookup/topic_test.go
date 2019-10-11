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

package lookup

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestLookupTopicCmd(t *testing.T) {
	args := []string{"create", "test-lookup-topic", "0"}
	_, execErr, _, _ := test.TestTopicCommands(crud.CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"lookup", "test-lookup-topic"}
	out, execErr, _, _ := test.TestTopicCommands(TopicCmd, args)
	assert.Nil(t, execErr)

	var data pulsar.LookupData
	err := json.Unmarshal(out.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}

	brokerURL := regexp.MustCompile("^pulsar://[a-z-A-Z0-9]*:6650$")
	assert.True(t, brokerURL.MatchString(data.BrokerURL))

	httpURL := regexp.MustCompile("^http://[a-z-A-Z0-9]*:8080$")
	assert.True(t, httpURL.MatchString(data.HTTPURL))
}

func TestLookupTopicArgError(t *testing.T) {
	args := []string{"lookup"}
	_, _, nameErr, _ := test.TestTopicCommands(TopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}
