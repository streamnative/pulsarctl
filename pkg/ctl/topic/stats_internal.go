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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func GetInternalStatsCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the internal stats for a non-partitioned topic or a " +
		"partition of a partitioned topic."
	desc.CommandPermission = "This command requires namespace admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic, partitioned topic"

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get internal stats for an existing non-partitioned-topic (topic-name)",
		Command: "pulsarctl topic internal-stats (topic-name)",
	}

	getPartition := cmdutils.Example{
		Desc:    "Get internal stats for a partition of a partitioned topic",
		Command: "pulsarctl topic internal-stats --partition (partition) (topic-name)",
	}
	examples = append(examples, get, getPartition)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: `{
  "entriesAddedCounter": 0,
  "numberOfEntries": 0,
  "totalSize": 0,
  "currentLedgerEntries": 0,
  "currentLedgerSize": 0,
  "lastLedgerCreatedTimestamp": "",
  "lastLedgerCreationFailureTimestamp": "",
  "waitingCursorsCount": 0,
  "pendingAddEntriesCount": 0,
  "lastConfirmedEntry": "",
  "state": "",
  "ledgers": [
    {
      "ledgerId": 0,
      "entries": 0,
      "size": 0,
      "offloaded": false
    }
  ],
  "cursors": {}
}`,
	}
	out = append(out, successOut, ArgError)

	partitionedTopicInternalStatsError := cmdutils.Output{
		Desc: "the specified topic is not exist or the specified topic is a partitioned topic",
		Out:  "[✖]  code: 404 reason: Topic not found",
	}
	out = append(out, partitionedTopicInternalStatsError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"internal-stats",
		"Get the internal stats of the specified topic",
		desc.ToString(),
		desc.ExampleToString(),
		"")

	var partition int

	vc.SetRunFuncWithNameArg(func() error {
		return doGetInternalStats(vc, partition)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Internal Stats", func(set *pflag.FlagSet) {
		set.IntVarP(&partition, "partition", "p", -1,
			"The partitioned topic index value")
	})
}

func doGetInternalStats(vc *cmdutils.VerbCmd, partition int) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if partition >= 0 {
		topic, err = topic.GetPartition(partition)
		if err != nil {
			return err
		}
	}

	admin := cmdutils.NewPulsarClient()
	stats, err := admin.Topics().GetInternalStats(*topic)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), stats)
	}

	return err
}
