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
	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

type backlogQuota struct {
	LimitSize string
	LimitTime int64
	Policy    string
	Type      string
}

func SetBacklogQuotaCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set a backlog quota policy for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	setBacklog := cmdutils.Example{
		Desc: desc.CommandUsedFor,
		Command: "pulsarctl topics set-backlog-quota topic \n" +
			"\t--limit-size 16G \n" +
			"\t--limit-time -1 \n" +
			"\t--policy <producer_request_hold|producer_exception|consumer_backlog_eviction> \n" +
			"\t--type <destination_storage|message_age>",
	}
	examples = append(examples, setBacklog)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set backlog quota successfully for [topic]",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-backlog-quota",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString(),
		"set-backlog-quota",
	)

	backlogQuota := backlogQuota{}

	vc.FlagSetGroup.InFlagSet("Set Backlog Quota", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&backlogQuota.LimitSize,
			"limit-size",
			"",
			"",
			"Size limit (eg: 10M, 16G)")

		flagSet.Int64VarP(
			&backlogQuota.LimitTime,
			"limit-time",
			"",
			-1,
			"Time limit in seconds")

		flagSet.StringVarP(
			&backlogQuota.Policy,
			"policy",
			"p",
			"",
			"Retention policy to enforce when the limit is reached.\n"+
				"Valid options are: [producer_request_hold, producer_exception, consumer_backlog_eviction]")

		flagSet.StringVarP(&backlogQuota.Type,
			"type",
			"t",
			string(util.DestinationStorage),
			"Backlog quota type to set.\n"+
				"Valid options are: [destination_storage, message_age]")

		_ = cobra.MarkFlagRequired(flagSet, "policy")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doSetBacklogQuota(vc, backlogQuota)
	}, "the topic name is not specified or the topic name is specified more than one")

}

func doSetBacklogQuota(vc *cmdutils.VerbCmd, data backlogQuota) error {
	topic, err := util.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()

	sizeLimit, err := utils.ValidateSizeString(data.LimitSize)
	if err != nil {
		return err
	}

	policy, err := util.ParseRetentionPolicy(data.Policy)
	if err != nil {
		return err
	}

	backlogQuotaType, err := util.ParseBacklogQuotaType(data.Type)
	if err != nil {
		return err
	}

	err = admin.Topics().SetBacklogQuota(*topic, util.NewBacklogQuota(sizeLimit, data.LimitTime, policy), backlogQuotaType)
	if err == nil {
		vc.Command.Printf("Set backlog quota successfully for [%s]\n", topic)
	}

	return err
}
