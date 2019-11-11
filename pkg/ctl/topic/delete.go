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

func DeleteTopicCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for deleting an existing topic."
	desc.CommandPermission = "This command requires namespace admin permissions."
	desc.CommandScope = "non-partitioned topic, partitioned topic"

	var examples []cmdutils.Example
	deleteTopic := cmdutils.Example{
		Desc:    "Delete a partitioned topic (topic-name)",
		Command: "pulsarctl topics delete (topic-name)",
	}

	deleteNonPartitionedTopic := cmdutils.Example{
		Desc:    "Delete a non-partitioned topic (topic-name)",
		Command: "pulsarctl topics delete --non-partitioned (topic-name)",
	}

	examples = append(examples, deleteTopic, deleteNonPartitionedTopic)
	desc.CommandExamples = examples
	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Delete topic (topic-name) successfully",
	}

	partitionedTopicNotExistError := cmdutils.Output{
		Desc: "the partitioned topic does not exist",
		Out:  "[✖]  code: 404 reason: Partitioned topic does not exist",
	}

	nonPartitionedTopicNotExistError := cmdutils.Output{
		Desc: "the non-partitioned topic does not exist",
		Out:  "[✖]  code: 404 reason: Topic not found",
	}
	out = append(out, successOut, ArgError,
		partitionedTopicNotExistError, nonPartitionedTopicNotExistError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"d")

	var force bool
	var deleteSchema bool
	var nonPartitioned bool

	vc.FlagSetGroup.InFlagSet("Delete Topic", func(set *pflag.FlagSet) {
		set.BoolVarP(&nonPartitioned, "non-partitioned", "n", false,
			"Delete a non-partitioned topic")
		set.BoolVarP(&force, "force", "f", false,
			"Close all producer/consumer/replicator and delete topic forcefully")
		set.BoolVarP(&deleteSchema, "delete-schema", "d", false,
			"Delete schema while deleting topic")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteTopic(vc, force, deleteSchema, nonPartitioned)
	}, "the topic name is not specified or the topic name is specified more than one")
}

// TODO add delete schema
func doDeleteTopic(vc *cmdutils.VerbCmd, force, deleteSchema, nonPartitioned bool) error {
	_ = deleteSchema
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().Delete(*topic, force, nonPartitioned)
	if err == nil {
		vc.Command.Printf("Delete topic %s successfully\n", topic.String())
	}

	return err
}
