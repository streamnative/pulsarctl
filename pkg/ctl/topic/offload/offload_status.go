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

package offload

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func OffloadStatusCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for checking the status of data offloading" +
		" from a persistent topic to long-term storage."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	offloadStatus := Example{
		Desc:    "Check the status of data offloading from a topic <persistent-topic-name> to long-term storage",
		Command: "pulsarctl topic offload-status <persistent-topic-name>",
	}

	waiting := Example{
		Desc:    "Wait for offloading to complete",
		Command: "pulsarctl topic offload-status --wait <persistent-topic-name>",
	}
	desc.CommandExamples = append(examples, offloadStatus, waiting)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Offloading topic <topic-name> data is done successfully",
	}

	notRun := Output{
		Desc: "Offloading topic is not running",
		Out: "Offloading topic <topic-name> data is not running",
	}

	running := Output{
		Desc: "Offloading topic is running",
		Out: "Offloading topic <topic-name> data is running",
	}

	errorOut := Output{
		Desc: "Offloading topic with error",
		Out:  "Offloading topic <topic-name> data is done with error <error-msg>",
	}
	out = append(out, successOut, notRun, running, errorOut, ArgError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"offload-status",
		"Check the status of data offloading",
		desc.ToString())

	var wait bool

	vc.SetRunFuncWithNameArg(func() error {
		return doOffloadStatus(vc, wait)
	})

	vc.FlagSetGroup.InFlagSet("OffloadStatus", func(set *pflag.FlagSet) {
		set.BoolVarP(&wait, "wait", "w", false, "Wait for offloading to complete")
	})
}

func doOffloadStatus(vc *cmdutils.VerbCmd, wait bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
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

	for wait && status.Status == RUNNING {
		time.Sleep(1 * time.Second)
		status, err = admin.Topics().OffloadStatus(*topic)
		if err != nil {
			return err
		}
	}

	switch status.Status {
	case NOT_RUN:
		vc.Command.Printf("Offloading topic %s data is not running", topic.String())
	case RUNNING:
		vc.Command.Printf("Offloading topic %s data is running", topic.String())
	case SUCCESS:
		vc.Command.Printf("Offloading topic %s data is done successfully", topic.String())
	case ERROR:
		vc.Command.Printf("Offloading topic %s data is done with error %s", topic.String(), status.LastError)
	}

	return err
}
