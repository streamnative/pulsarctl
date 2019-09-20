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
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestLookupTopicCmd(t *testing.T) {
	args := []string{"create", "test-lookup-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"lookup", "test-lookup-topic"}
	out, execErr, _, _  := TestTopicCommands(LookupTopicCmd, args)
	assert.Nil(t, execErr)

	var data pulsar.LookupData
	err := json.Unmarshal(out.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}

	brokerUrl, err := regexp.Compile("^pulsar://[a-z-A-Z0-9]*:6650$")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, brokerUrl.MatchString(data.BrokerUrl))

	httpUrl, err :=  regexp.Compile("^http://[a-z-A-Z0-9]*:8080$")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, httpUrl.MatchString(data.HttpUrl))
}

func TestLookupTopicArgError(t *testing.T)  {
	args  := []string{"lookup"}
	_, _, nameErr, _ := TestTopicCommands(LookupTopicCmd,  args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}
