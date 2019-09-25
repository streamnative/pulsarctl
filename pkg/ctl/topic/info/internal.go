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

package info

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

var normalOutput =
`{
  "version": 1,
  "creationDate": "",
  "modificationData": "",
  "ledgers": [
    {
      "ledgerId": 0,
      "entries": 0,
      "size": 0,
      "timestamp": 0
    }
  ],
  "terminatedPosition": {
    "ledgerId": 0,
    "entryId": 0
  },
  "cursors": {
    "hello": {
      "version": 0,
      "creationDate": "",
      "modificationDate": "",
      "cursorsLedgerId": 0,
      "markDelete": {
        "ledgerId": 0,
        "entryId": 0
      },
      "individualDeletedMessages": null,
      "Properties": null
    }
  }
}
`

func GetInternalInfoCmd(vc *cmdutils.VerbCmd)  {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the internal info of a topic " +
		"which has messages or subscriptions."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	get := Example{
		Desc: "Get the internal info of a topic <topic-name>",
		Command: "pulsarctl topic internal-info <topic-name>",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: normalOutput,
	}

	failOut := Output{
		Desc: "the specified topic does not exist",
		Out: "[✖]  code: 500 reason: Unknown pulsar error",
	}
	out = append(out, successOut, ArgError, failOut)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out
	
	vc.SetDescription("internal-info",
		"Get the topic internal info",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetInternalInfo(vc)
	})
}

func doGetInternalInfo(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	info, err := admin.Topics().GetInternalInfo(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), info)
	}

	return err
}
