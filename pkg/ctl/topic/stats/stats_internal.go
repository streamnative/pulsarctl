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

package stats

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetInternalStatsCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the internal stats for a non-partitioned topic or a " +
		"partition of a partitioned topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get internal stats for an existing non-partitioned-topic (topic-name)",
		Command: "pulsarctl topic internal-stats (topic-name)",
	}

	getPartition := Example{
		Desc:    "Get internal stats for a partition of a partitioned topic",
		Command: "pulsarctl topic internal-stats --partition (partition) (topic-name)",
	}
	desc.CommandExamples = append(examples, get, getPartition)

	var out []Output
	successOut := Output{
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

	partitionedTopicInternalStatsError := Output{
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
	})

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

	topic, err := GetTopicName(vc.NameArg)
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
		cmdutils.PrintJson(vc.Command.OutOrStdout(), stats)
	}

	return err
}
