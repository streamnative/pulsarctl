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
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func CreateCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for creating a subscription on a topic."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc:    "Create a subscription (subscription-name) on a topic (topic-name) from latest position",
		Command: "pulsarctl subscriptions create (topic-name) (subscription-name)",
	}

	createWithFlag := cmdutils.Example{
		Desc: "Create a subscription (subscription-name) on a topic (topic-name) from the specified " +
			"position (position)",
		Command: "pulsarctl subscription create --messageId (position) (topic-name) (subscription-name)",
	}
	examples = append(examples, create, createWithFlag)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Create subscription (subscription-name) on topic (topic-name) from (position) successfully",
	}
	out = append(out, successOut, ArgsError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	out = append(out, MessageIDErrors...)
	desc.CommandOutput = out

	var ID string

	vc.SetDescription(
		"create",
		"Create a subscription on a topic from latest or a specific position",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doCreate(vc, ID)
	}, CheckSubscriptionNameTwoArgs)

	vc.FlagSetGroup.InFlagSet("Create Subscription", func(set *pflag.FlagSet) {
		set.StringVarP(&ID, "messageId", "m", "latest",
			"message id where the subscription starts from. It can be either 'latest', "+
				"'earliest' or (ledgerId:entryId)")
	})
}

func doCreate(vc *cmdutils.VerbCmd, id string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	sName := vc.NameArgs[1]

	var messageID utils.MessageID
	switch id {
	case "latest":
		messageID = utils.Latest
	case "earliest":
		messageID = utils.Earliest
	default:
		s := strings.Split(id, ":")
		if len(s) != 2 {
			return errors.Errorf("invalid position value : %s", id)
		}
		i, err := utils.ParseMessageID(id)
		if err != nil {
			return err
		}
		messageID = *i
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Subscriptions().Create(*topic, sName, messageID)
	if err == nil {
		vc.Command.Printf("Create subscription %s on topic %s starting from %s successfully\n",
			sName, topic.String(), id)
	}

	return err
}
