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
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	ctlutils "github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

type setInactiveTopicPoliciesArgs struct {
	deleteWhileInactive                     bool
	deleteInactiveTopicsMaxInactiveDuration string
	inactiveTopicDeleteMode                 string
}

func SetInactiveTopicCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Set the inactive topic policies on a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc: desc.CommandUsedFor,
		Command: "pulsarctl topics set-inactive-topic-policies [topic] \n" +
			"\t--enable-delete-while-inactive true \n" +
			"\t--max-inactive-duration 1h \n" +
			"\t--delete-mode delete_when_no_subscriptions",
	}

	examples = append(examples, set)
	desc.CommandExamples = examples

	vc.SetDescription(
		"set-inactive-topic-policies",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString())

	args := setInactiveTopicPoliciesArgs{}

	vc.FlagSetGroup.InFlagSet("Set Inactive Topic", func(flagSet *pflag.FlagSet) {
		flagSet.BoolVarP(&args.deleteWhileInactive,
			"enable-delete-while-inactive",
			"e",
			false,
			"Control whether deletion is enabled while inactive")

		flagSet.StringVarP(&args.deleteInactiveTopicsMaxInactiveDuration,
			"max-inactive-duration",
			"t",
			"",
			"Max duration of topic inactivity in seconds, "+
				"topics that are inactive for longer than this value will be deleted (eg: 1s, 10s, 1m, 5h, 3d)")
		flagSet.StringVarP(&args.inactiveTopicDeleteMode,
			"delete-mode",
			"m",
			"",
			"Mode of delete inactive topic, "+
				"Valid options are: [delete_when_no_subscriptions, delete_when_subscriptions_caught_up]")

		_ = cobra.MarkFlagRequired(flagSet, "delete-mode")
		_ = cobra.MarkFlagRequired(flagSet, "max-inactive-duration")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doSetInactiveTopic(vc, args)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doSetInactiveTopic(vc *cmdutils.VerbCmd, args setInactiveTopicPoliciesArgs) error {
	inactiveTopicDeleteMode, err := utils.ParseInactiveTopicDeleteMode(args.inactiveTopicDeleteMode)
	if err != nil {
		return err
	}

	maxInactiveDuration, err := ctlutils.ParseRelativeTimeInSeconds(args.deleteInactiveTopicsMaxInactiveDuration)
	if err != nil {
		return err
	}

	body := utils.NewInactiveTopicPolicies(
		&inactiveTopicDeleteMode,
		int(maxInactiveDuration.Seconds()),
		args.deleteWhileInactive)

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetInactiveTopicPolicies(*topic, body)
	if err == nil {
		vc.Command.Printf("Set inactive topic policies successfully for [%s]", topic.String())
	}

	return err
}
