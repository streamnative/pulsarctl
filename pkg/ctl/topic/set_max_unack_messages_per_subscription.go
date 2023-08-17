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
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func SetMaxUnackMessagesPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set max unacked messages per subscription for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Set max unacked messages per subscription for a topic",
		Command: "pulsarctl topics set-max-unacked-messages-per-subscription topic -m 10",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set max unacked messages per subscription successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-max-unacked-messages-per-subscription",
		"Set max unacked messages per subscription for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-max-unacked-messages-per-subscription",
	)
	var maxUnackedNum int
	vc.SetRunFuncWithNameArg(func() error {
		return doSetMaxUnackMessagesPerSubscription(vc, maxUnackedNum)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("MaxUnackedMessagesPerSubscription", func(set *pflag.FlagSet) {
		set.IntVarP(
			&maxUnackedNum,
			"maxNum",
			"m",
			0,
			"Max unacked messages per subscription for a topic")
	})
	vc.EnableOutputFlagSet()
}

func doSetMaxUnackMessagesPerSubscription(vc *cmdutils.VerbCmd, maxUnackedNum int) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetMaxUnackMessagesPerSubscription(*topic, maxUnackedNum)
	if err == nil {
		vc.Command.Printf("Set max unacked messages per subscription successfully for [%s]\n", topic.String())
	}
	return err
}
