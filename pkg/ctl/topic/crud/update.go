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

package crud

import (
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/args"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
)

func UpdateTopicCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for updating the partition number of an exist topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example
	updateTopic := pulsar.Example{
		Desc:    "",
		Command: "pulsarctl topics update (topic-name) (partition-num)",
	}
	examples = append(examples, updateTopic)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Update topic (topic-name) with (partition-num) partitions successfully",
	}

	topicNotExist := pulsar.Output{
		Desc: "the topic is not exist",
		Out:  "[✖]  code: 409 reason: Topic is not partitioned topic",
	}
	out = append(out, successOut, e.ArgsError, e.InvalidPartitionsNumberError, topicNotExist)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"update",
		"Update partitioned topic partitions",
		desc.ToString(),
		desc.ExampleToString(),
		"up")

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doUpdateTopic(vc)
	}, args.CheckTopicNameTwoArgs)
}

func doUpdateTopic(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	partitions, err := strconv.Atoi(vc.NameArgs[1])
	if err != nil || partitions <= 0 {
		return errors.Errorf("invalid partition number '%s'", vc.NameArgs[1])
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().Update(*topic, partitions)
	if err == nil {
		vc.Command.Printf("Update topic %s with %d partitions successfully\n", topic.String(), partitions)
	}

	return err
}
