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
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

var normalOutput = `{
  "version": 1,
  "creationDate": "",
  "modificationData": "",
  "ledgers": [
    {
      "ledgerId": 0,
      "entries": 0,
      "size": 0,
      "offloaded": false
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

func GetInternalInfoCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the internal info of a topic " +
		"which has messages or subscriptions."
	desc.CommandPermission = "This command requires tenant admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic"

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the internal info of a topic <topic-name>",
		Command: "pulsarctl topic internal-info <topic-name>",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  normalOutput,
	}

	failOut := cmdutils.Output{
		Desc: "the specified topic does not exist",
		Out:  "[✖]  code: 500 reason: Unknown error",
	}
	out = append(out, successOut, ArgError, failOut)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription("internal-info",
		"Get the topic internal info",
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetInternalInfo(vc)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.EnableOutputFlagSet()
}

func doGetInternalInfo(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	info, err := admin.Topics().GetInternalInfo(*topic)
	if err == nil {
		oc := cmdutils.NewOutputContent().WithObject(info)
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
	}

	return err
}
