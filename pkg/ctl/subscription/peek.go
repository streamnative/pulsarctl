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

package subscription

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func PeekCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for peeking some messages of a subscription."
	desc.CommandPermission = "This command requires tenant admin permissions or namespace consumer permissions."

	var example []cmdutils.Example
	peek := cmdutils.Example{
		Desc:    "Peek some messages of a subscription",
		Command: "pulsarctl subscriptions peek --count (n) (topic-name) (subscription-name)",
	}
	example = append(example, peek)
	desc.CommandExamples = example

	var out []cmdutils.Output
	success := cmdutils.Output{
		Desc: "normal output",
		Out: `Message ID :
ledgerID:entryID:PartitionIndex:BatchIndex
Properties :
Message :`,
	}
	out = append(out, success, ArgsError)
	out = append(out, TopicNameErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"peek",
		"Peek some messages of a subscription",
		desc.ToString(),
		desc.ExampleToString())

	var count int

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doPeek(vc, count)
	}, CheckSubscriptionNameTwoArgs)

	vc.FlagSetGroup.InFlagSet("Peek", func(set *pflag.FlagSet) {
		set.IntVarP(&count, "count", "n", 1, "Number of messages (default 1)")
	})
}

func doPeek(vc *cmdutils.VerbCmd, n int) error {

	topic, err := utils.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	if topic.GetDomain().String() != "persistent" {
		return errors.New("the specified topic name is not a persistent topic")
	}

	sName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	msgs, err := admin.Subscriptions().PeekMessages(*topic, sName, n)
	if err == nil {
		pos := 0
		out := ""

		for _, v := range msgs {
			m := *v
			pos++
			if pos != 1 {
				out += fmt.Sprintln("-------------------------------------------------------------------------")
			}
			out += fmt.Sprintln("Message ID : " + m.GetMessageID().String())
			out += fmt.Sprintln("Properties : ")
			p, _ := json.MarshalIndent(m.GetProperties(), "", "    ")
			out += fmt.Sprintln(string(p))
			out += fmt.Sprintln("Message :")
			out += fmt.Sprintln(hex.Dump(m.GetPayload()))
		}

		vc.Command.Println(out)
	}

	return err
}
