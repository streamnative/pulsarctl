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

func GetStatsCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the stats for an existing topic and its " +
		"connected producers and consumers. (All the rates are computed over a 1 minute window " +
		"and are relative the last completed 1 minute period)"
	desc.CommandPermission = "This command requires namespace admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic, partitioned topic"

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the non-partitioned topic (topic-name) stats",
		Command: "pulsarctl topic stats (topic-name)",
	}

	getPartition := cmdutils.Example{
		Desc:    "Get the partitioned topic (topic-name) stats",
		Command: "pulsarctl topic stats --partitioned-topic (topic-name)",
	}

	getPerPartition := cmdutils.Example{
		Desc:    "Get the partitioned topic (topic-name) stats and per partition stats",
		Command: "pulsarctl topic stats --partitioned-topic --per-partition (topic-name)",
	}
	examples = append(examples, get, getPartition, getPerPartition)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Get the non-partitioned topic stats",
		Out: `{
  "msgRateIn": 0,
  "msgRateOut": 0,
  "msgThroughputIn": 0,
  "msgThroughputOut": 0,
  "averageMsgSize": 0,
  "storageSize": 0,
  "publishers": [],
  "subscriptions": {},
  "replication": {},
  "deduplicationStatus": "Disabled"
}`,
	}

	partitionOutput := cmdutils.Output{
		Desc: "Get the partitioned topic stats",
		Out: `{
  "msgRateIn": 0,
  "msgRateOut": 0,
  "msgThroughputIn": 0,
  "msgThroughputOut": 0,
  "averageMsgSize": 0,
  "storageSize": 0,
  "publishers": [],
  "subscriptions": {},
  "replication": {},
  "deduplicationStatus": "",
  "metadata": {
    "partitions": 1
  },
  "partitions": {}
}`,
	}

	perPartitionOutput := cmdutils.Output{
		Desc: "Get the partitioned topic stats and per partition topic stats",
		Out: `{
  "msgRateIn": 0,
  "msgRateOut": 0,
  "msgThroughputIn": 0,
  "msgThroughputOut": 0,
  "averageMsgSize": 0,
  "storageSize": 0,
  "publishers": [],
  "subscriptions": {},
  "replication": {},
  "deduplicationStatus": "",
  "metadata": {
    "partitions": 1
  },
  "partitions": {
    "<topic-name>": {
      "msgRateIn": 0,
      "msgRateOut": 0,
      "msgThroughputIn": 0,
      "msgThroughputOut": 0,
      "averageMsgSize": 0,
      "storageSize": 0,
      "publishers": [],
      "subscriptions": {},
      "replication": {},
      "deduplicationStatus": ""
    }
  }
}`,
	}
	out = append(out, successOut, partitionOutput, perPartitionOutput, ArgError)

	topicNotFoundError := cmdutils.Output{
		Desc: "the specified topic does not exist " +
			"or the specified topic is a partitioned-topic and you don't specified --partitioned-topic " +
			"or the specified topic is a non-partitioned topic and you specified --partitioned-topic",
		Out: "code: 404 reason: Topic not found",
	}
	out = append(out, topicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"stats",
		"Get the stats of an existing topic",
		desc.ToString(),
		desc.ExampleToString(),
		"stats",
	)

	var partition bool
	var perPartition bool
	vc.SetRunFuncWithNameArg(func() error {
		return doGetStats(vc, partition, perPartition)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Stats", func(set *pflag.FlagSet) {
		set.BoolVarP(&partition, "partitioned-topic", "p", false,
			"Get the partitioned topic stats")
		set.BoolVarP(&perPartition, "per-partition", "", false,
			"Get the per partition topic stats")
	})
}

func doGetStats(vc *cmdutils.VerbCmd, partitionedTopic, perPartition bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()

	if partitionedTopic {
		stats, err := admin.Topics().GetPartitionedStats(*topic, perPartition)
		if err == nil {
			cmdutils.PrintJSON(vc.Command.OutOrStdout(), stats)
		}
		return err
	}

	topicStats, err := admin.Topics().GetStats(*topic)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), topicStats)
	}
	return err
}
