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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func GetInactiveTopicCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Get the inactive topic policies on a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    desc.CommandUsedFor,
		Command: "pulsarctl topics get-inactive-topic-policies [topic]",
	})
	desc.CommandExamples = examples

	vc.SetDescription(
		"get-inactive-topic-policies",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString())

	var applied bool
	vc.FlagSetGroup.InFlagSet("Get Inactive Topic", func(flagSet *pflag.FlagSet) {
		flagSet.BoolVarP(&applied, "applied", "", false, "Get the applied policy for the topic")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doGetInactiveTopic(vc, applied)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doGetInactiveTopic(vc *cmdutils.VerbCmd, applied bool) error {
	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	response, err := admin.Topics().GetInactiveTopicPolicies(*topic, applied)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), &response)
	}

	return err
}
