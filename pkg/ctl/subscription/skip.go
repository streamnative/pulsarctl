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
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func SkipCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for skipping messages for a subscription."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []pulsar.Example
	skip := pulsar.Example{
		Desc:    "Skip (n) messages for the subscription (subscription-name) of the topic (topic-name)",
		Command: "pulsarctl subscription skip --count (n) (topic-name) (subscription-name)",
	}

	skipAll := pulsar.Example{
		Desc:    "Skip all messages for the subscription (subscription-name) under the topic (topic-name) (clear-backlog)",
		Command: "pulsarctl subscription skip --all (topic-name) (subscription-name)",
	}
	examples = append(examples, skip, skipAll)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "The subscription (subscription-name) skips (n) messages of the topic <topic-name> successfully",
	}
	out = append(out, successOut, ArgsError, TopicNotFoundError, SubNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"skip",
		"Skip messages for the subscription of a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"skip")

	var count int64
	var all bool

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doSkip(vc, count, all)
	}, CheckSubscriptionNameTwoArgs)

	vc.FlagSetGroup.InFlagSet("Skip Messages", func(set *pflag.FlagSet) {
		set.Int64VarP(&count, "count", "n", -1,
			"number of messages to skip")
		set.BoolVarP(&all, "all", "a", false, "skip all messages")
	})
}

func doSkip(vc *cmdutils.VerbCmd, count int64, all bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	if !all && count < 0 {
		return errors.New("the skip message number is not specified")
	}

	topic, err := pulsar.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	sName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	if all {
		err = admin.Subscriptions().ClearBacklog(*topic, sName)
	} else {
		err = admin.Subscriptions().SkipMessages(*topic, sName, count)
	}

	if err == nil {
		vc.Command.Printf("The subscription %s skips %d messages of the topic %s successfully\n",
			sName, count, topic.String())
	}

	return err
}
