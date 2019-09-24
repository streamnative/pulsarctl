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

func GetStatsCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the stats for an existing topic and its " +
		"connected producers and consumers. (All the rates are computed over a 1 minute window " +
		"and are relative the last completed 1 minute period)"
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get the non-partitioned topic <topic-name> stats",
		Command: "pulsarctl topic stats <topic-name>",
	}

	getPartition := Example{
		Desc:    "Get the partitioned topic <topic-name> stats",
		Command: "pulsarctl topic stats --partitioned-topic <topic-name>",
	}

	getPerPartition := Example{
		Desc:    "Get the partitioned topic <topic-name> stats and per partition stats",
		Command: "pulsarctl topic stats --partitioned-topic --per-partition <topic-name>",
	}

	desc.CommandExamples = append(examples, get, getPartition, getPerPartition)

	var out []Output
	successOut := Output{
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

	partitionOutput := Output{
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

	perPartitionOutput := Output{
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

	topicNotFoundError := Output{
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
		desc.ToString())

	var partition bool
	var perPartition bool
	vc.SetRunFuncWithNameArg(func() error {
		return doGetStats(vc, partition, perPartition)
	})

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

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()

	if partitionedTopic {
		stats, err := admin.Topics().GetPartitionedStats(*topic, perPartition)
		if err == nil {
			cmdutils.PrintJson(vc.Command.OutOrStdout(), stats)
		}
		return err
	} else {
		topicStats, err := admin.Topics().GetStats(*topic)
		if err == nil {
			cmdutils.PrintJson(vc.Command.OutOrStdout(), topicStats)
		}
		return err
	}
}
