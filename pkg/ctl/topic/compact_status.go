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
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func StatusCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting compaction status of a topic " +
		"or a partition of a partitioned topic."
	desc.CommandPermission = "This command requires tenant admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic"

	var examples []pulsar.Example
	compactStatus := pulsar.Example{
		Desc:    "Get compaction status of a persistent topic (topic-name)",
		Command: "pulsarctl topic compact-status (topic-name)",
	}

	compactPartitionStatus := pulsar.Example{
		Desc:    "Get compaction status of a partition of partitioned topic",
		Command: "pulsarctl topic compact-status --partition (partition) (topic-name)",
	}
	examples = append(examples, compactStatus, compactPartitionStatus)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Compacting the topic (topic-name) is done successfully",
	}

	notRun := pulsar.Output{
		Desc: "Compacting the topic (topic-name) is not running",
		Out:  "Compacting the topic (topic-name) is not running",
	}

	running := pulsar.Output{
		Desc: "Compacting the topic (topic-name) is running",
		Out:  "Compacting the topic (topic-name) is running",
	}

	errorOut := pulsar.Output{
		Desc: "Compacting the topic (topic-name) is done with error",
		Out:  "Compacting the topic (topic-name) is done with error <error-msg>",
	}
	out = append(out, successOut, notRun, running, errorOut, ArgError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"compact-status",
		"Get status of compaction on a topic",
		desc.ToString(),
		desc.ExampleToString())

	var wait bool
	var partition int

	vc.SetRunFuncWithNameArg(func() error {
		return doCompactStatus(vc, wait, partition)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Compact Status", func(set *pflag.FlagSet) {
		set.BoolVarP(&wait, "wait", "w", false, "Wait for compacting to complete")
		set.IntVarP(&partition, "partition", "p", -1,
			"The partitioned topic index value")
	})
}

func doCompactStatus(vc *cmdutils.VerbCmd, wait bool, partition int) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if !topic.IsPersistent() {
		return errors.New("need to provide a persistent topic")
	}

	if partition >= 0 {
		topic, err = topic.GetPartition(partition)
		if err != nil {
			return err
		}
	}

	admin := cmdutils.NewPulsarClient()
	status, err := admin.Topics().CompactStatus(*topic)
	if err != nil {
		return err
	}

	for wait && status.Status == pulsar.RUNNING {
		time.Sleep(1 * time.Second)
		status, err = admin.Topics().CompactStatus(*topic)
		if err != nil {
			return err
		}
	}

	switch status.Status {
	case pulsar.NOTRUN:
		vc.Command.Printf("Compacting the topic %s is not running\n", topic.String())
	case pulsar.RUNNING:
		vc.Command.Printf("Compacting the topic %s is running\n", topic.String())
	case pulsar.SUCCESS:
		vc.Command.Printf("Compacting the topic %s is done successfully\n", topic.String())
	case pulsar.ERROR:
		vc.Command.Printf("Compacting the topic %s is done with error %s\n", topic.String(), status.LastError)
	}

	return err
}
