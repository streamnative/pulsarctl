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

package stop

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func TopicTerminateCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for terminating a non-partitioned topic or a partition of " +
		"a partitioned topic. Upon termination, no more messages are allowed to published to it."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	terminate := pulsar.Example{
		Desc:    "Terminate a non-partitioned topic (topic-name)",
		Command: "pulsarctl topic terminate (topic-name)",
	}

	terminateWithPartition := pulsar.Example{
		Desc:    "Terminate a partition of a partitioned topic",
		Command: "pulsarctl topic terminate --partition (partition) (topic-name)",
	}
	examples = append(examples, terminate, terminateWithPartition)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Topic (topic-name) is successfully terminated at (message-id)",
	}

	partitionError := pulsar.Output{
		Desc: "the specified topic is a partitioned topic",
		Out:  "[✖]  code: 405 reason: Termination of a partitioned topic is not allowed",
	}
	out = append(out, successOut, e.ArgError, e.TopicNotFoundError, partitionError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"terminate",
		"Terminate a non-partitioned topic",
		desc.ToString(),
		desc.ExampleToString())

	var partition int

	vc.SetRunFuncWithNameArg(func() error {
		return doTerminate(vc, partition)
	})

	vc.FlagSetGroup.InFlagSet("Terminate", func(set *pflag.FlagSet) {
		set.IntVarP(&partition, "partition", "p", -1,
			"The partitioned topic index value")
	})
}

func doTerminate(vc *cmdutils.VerbCmd, partition int) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if !topic.IsPersistent() {
		return errors.New("only support terminating a persistent topic")
	}

	if partition >= 0 {
		topic, err = topic.GetPartition(partition)
		if err != nil {
			return err
		}
	}

	admin := cmdutils.NewPulsarClient()
	messageID, err := admin.Topics().Terminate(*topic)
	if err == nil {
		vc.Command.Printf("Topic %s is successfully terminated at %+v", topic.String(), messageID)
	}

	return err
}
