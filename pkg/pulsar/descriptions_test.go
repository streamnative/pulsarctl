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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongDescription_exampleToString(t *testing.T) {
	desc := LongDescription{}
	example := Example{
		Desc:    "command description",
		Command: "command",
	}
	desc.CommandExamples = []Example{example}
	res := desc.ExampleToString()

	expect := "    #command description\n" +
		"    command\n\n"

	assert.Equal(t, expect, res)
}

func TestLongDescription_ToString(t *testing.T) {
	desc := LongDescription{}
	desc.CommandUsedFor = "command used for"
	desc.CommandPermission = "command permission"
	example := Example{}
	example.Desc = "command description"
	example.Command = "command"
	desc.CommandExamples = []Example{example}
	out := Output{
		Desc: "Output",
		Out:  "Out line 1\nOut line 2",
	}
	desc.CommandOutput = []Output{out}

	expect := "USED FOR:\n" +
		"    " + desc.CommandUsedFor + "\n\n" +
		"REQUIRED PERMISSION:\n" +
		"    " + desc.CommandPermission + "\n\n" +
		"OUTPUT:\n" +
		"    " + "#" + out.Desc + "\n" +
		"    " + "Out line 1" + "\n" +
		"    " + "Out line 2" + "\n\n"

	result := desc.ToString()

	assert.Equal(t, expect, result)
}
