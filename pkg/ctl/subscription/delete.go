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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func DeleteCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for deleting a subscription of a topic."
	desc.CommandPermission = "This command requires tenant admin and namespace consume permissions."

	var examples []cmdutils.Example
	deleteSub := cmdutils.Example{
		Desc:    "Delete the subscription (subscription-name) of the topic (topic-name)",
		Command: "pulsarctl subscriptions delete (topic-name) (subscription-name)",
	}
	examples = append(examples, deleteSub)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Delete the subscription %s of the topic %s successfully",
	}
	out = append(out, successOut, ArgsError, SubNotFoundError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete a subscription of a topic",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doDelete(vc)
	}, CheckSubscriptionNameTwoArgs)
}

func doDelete(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	sName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	err = admin.Subscriptions().Delete(*topic, sName)
	if err == nil {
		vc.Command.Printf("Delete the subscription %s of the topic %s successfully\n",
			sName, topic.String())
	}

	return err
}
