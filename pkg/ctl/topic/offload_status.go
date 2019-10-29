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

func OffloadStatusCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for checking the status of offloading data" +
		" from a persistent topic to a long-term storage."
	desc.CommandPermission = "This command requires tenant admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic"

	var examples []pulsar.Example
	offloadStatus := pulsar.Example{
		Desc:    "Check the status of offloading data from a topic (persistent-topic-name) to a long-term storage",
		Command: "pulsarctl topic offload-status (persistent-topic-name)",
	}

	waiting := pulsar.Example{
		Desc:    "Wait for offloading to complete",
		Command: "pulsarctl topic offload-status --wait (persistent-topic-name)",
	}
	examples = append(examples, offloadStatus, waiting)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Offloading topic (topic-name) data is done successfully",
	}

	notRun := pulsar.Output{
		Desc: "Offloading topic is not running",
		Out:  "Offloading topic (topic-name) data is not running",
	}

	running := pulsar.Output{
		Desc: "Offloading topic is running",
		Out:  "Offloading topic (topic-name) data is running",
	}

	errorOut := pulsar.Output{
		Desc: "Offloading topic with error",
		Out:  "Offloading topic (topic-name) data is done with error (error-msg)",
	}
	out = append(out, successOut, notRun, running, errorOut, ArgError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"offload-status",
		"Check the status of data offloading",
		desc.ToString(),
		desc.ExampleToString())

	var wait bool

	vc.SetRunFuncWithNameArg(func() error {
		return doOffloadStatus(vc, wait)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("OffloadStatus", func(set *pflag.FlagSet) {
		set.BoolVarP(&wait, "wait", "w", false, "Wait for offloading to complete")
	})
}

func doOffloadStatus(vc *cmdutils.VerbCmd, wait bool) error {
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

	admin := cmdutils.NewPulsarClient()
	status, err := admin.Topics().OffloadStatus(*topic)
	if err != nil {
		return err
	}

	for wait && status.Status == pulsar.RUNNING {
		time.Sleep(1 * time.Second)
		status, err = admin.Topics().OffloadStatus(*topic)
		if err != nil {
			return err
		}
	}

	switch status.Status {
	case pulsar.NOTRUN:
		vc.Command.Printf("Offloading topic %s is not running\n", topic.String())
	case pulsar.RUNNING:
		vc.Command.Printf("Offloading topic %s is running\n", topic.String())
	case pulsar.SUCCESS:
		vc.Command.Printf("Offloading topic %s is done successfully\n", topic.String())
	case pulsar.ERROR:
		vc.Command.Printf("Offloading topic %s is done with error %s\n", topic.String(), status.LastError)
	}

	return err
}
