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

func GetBacklogQuotaCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the backlog quota policy for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	getBacklog := cmdutils.Example{
		Desc:    desc.CommandUsedFor,
		Command: "pulsarctl topics get-backlog-quotas topic",
	}
	examples = append(examples, getBacklog)
	desc.CommandExamples = examples

	vc.SetDescription(
		"get-backlog-quotas",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString(),
		"get-backlog-quotas",
	)

	var applied bool
	vc.FlagSetGroup.InFlagSet("Get Backlog Quota", func(flagSet *pflag.FlagSet) {
		flagSet.BoolVarP(
			&applied,
			"applied",
			"",
			false,
			"Get the applied policy for the topic")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doGetBacklogQuota(vc, applied)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doGetBacklogQuota(vc *cmdutils.VerbCmd, applied bool) error {
	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	backlogQuotasMap, err := admin.Topics().GetBacklogQuotaMap(*topic, applied)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), &backlogQuotasMap)
	}

	return err
}
