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

func GetLastMessageIDCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the last message id of a topic (partition)."
	desc.CommandPermission = "This command requires tenant admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic"

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the last message id of a topic (persistent-topic-name)",
		Command: "pulsarctl topic last-message-id (persistent-topic-name)",
	}

	getPartitionedTopic := cmdutils.Example{
		Desc:    "Get the last message id of a partition of a partitioned topic (topic-name)",
		Command: "pulsarctl topic last-message-id --partition (partition) (topic-name)",
	}
	examples = append(examples, get, getPartitionedTopic)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"LedgerID\": 0,\n" +
			"  \"EntryID\": 0,\n" +
			"  \"PartitionedIndex\": 0" +
			"\n}",
	}
	out = append(out, successOut, ArgError)

	topicNotFoundError := cmdutils.Output{
		Desc: "the topic (persistent-topic-name) does not exist in the cluster",
		Out:  "[✖]  code: 404 reason: Topic not found",
	}
	out = append(out, topicNotFoundError)

	notAllowedError := cmdutils.Output{
		Desc: "the topic (persistent-topic-name) does not a persistent topic",
		Out:  "[✖]  code: 405 reason: GetLastMessageId on a non-persistent topic is not allowed",
	}
	out = append(out, notAllowedError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	var partition int

	vc.SetDescription(
		"last-message-id",
		"Get the last message id of a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"lmi")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetLastMessageID(vc, partition)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("LastMessageId", func(set *pflag.FlagSet) {
		set.IntVarP(&partition, "partition", "p", -1,
			"The partitioned topic index value")
	})
}

func doGetLastMessageID(vc *cmdutils.VerbCmd, partition int) error {
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
	messageID, err := admin.Topics().GetLastMessageID(*topic)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), messageID)
	}

	return err
}
