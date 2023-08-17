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

func RemoveBacklogQuotaCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Remove a backlog quota policy from a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	removeBacklog := cmdutils.Example{
		Desc:    "Remove a backlog quota policy from a topic",
		Command: "pulsarctl topics remove-backlog-quota topic -t <destination_storage|message_age>",
	}
	examples = append(examples, removeBacklog)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Remove backlog quota successfully for [topic]",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"remove-backlog-quota",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString(),
		"remove-backlog-quota",
	)

	var backlogQuotaType string

	vc.FlagSetGroup.InFlagSet("Remove backlog quota", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&backlogQuotaType,
			"type",
			"t",
			string(utils.DestinationStorage),
			"Backlog quota type to remove",
		)
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveBacklogQuota(vc, utils.BacklogQuotaType(backlogQuotaType))
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doRemoveBacklogQuota(vc *cmdutils.VerbCmd, backlogQuotaType utils.BacklogQuotaType) error {
	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().RemoveBacklogQuota(*topic, backlogQuotaType)
	if err == nil {
		vc.Command.Printf("Remove backlog quota successfully for [%s]\n", topic)
	}

	return err
}
