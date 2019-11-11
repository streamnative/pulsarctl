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

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func CompactCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for compacting a persistent topic or a partition of a partitioned topic."
	desc.CommandPermission = "This command is requires tenant admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic"

	var examples []cmdutils.Example
	compact := cmdutils.Example{
		Desc:    "Compact a persistent topic (topic-name)",
		Command: "pulsarctl topic compact (topic-name)",
	}

	compactPartition := cmdutils.Example{
		Desc:    "Compact a partition of a partitioned topic",
		Command: "pulsarctl topic compact --partition (index) (topic-name)",
	}
	examples = append(examples, compact, compactPartition)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Sending compact topic (topic-name) request successfully",
	}
	out = append(out, successOut, ArgError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"compact",
		"Compact a topic",
		desc.ToString(),
		desc.ExampleToString())

	var partition int

	vc.SetRunFuncWithNameArg(func() error {
		return doCompact(vc, partition)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Compact", func(set *pflag.FlagSet) {
		set.IntVarP(&partition, "partition", "p", -1,
			"The partitioned topic index value")
	})
}

func doCompact(vc *cmdutils.VerbCmd, partition int) error {
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

	if !topic.IsPersistent() {
		return errors.New("need to provide a persistent topic")
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().Compact(*topic)
	if err == nil {
		vc.Command.Printf("Successfully triggered compacting topic %s\n", topic.String())
	}

	return err
}
