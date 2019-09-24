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

package compact

import (
	"time"

	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func CompactStatusCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting status of compaction on a topic."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	compactStatus := Example{
		Desc:    "Get status of compaction of a persistent topic <topic-name>",
		Command: "pulsarctl topic compact-status <topic-name>",
	}
	desc.CommandExamples = append(examples, compactStatus)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Compacting the topic <topic-name> is done successfully",
	}

	notRun := Output{
		Desc: "Compacting the topic <topic-name> is not running",
		Out:  "Compacting the topic <topic-name> is not running",
	}

	running := Output{
		Desc: "Compacting the topic <topic-name> is running",
		Out:  "Compacting the topic <topic-name> is running",
	}

	errorOut := Output{
		Desc: "Compacting the topic <topic-name> is done with error",
		Out: "Compacting the topic <topic-name> is done with error <error-msg>",
	}
	out = append(out, successOut, notRun, running, errorOut, ArgError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"compact-status",
		"Get status of compaction on a topic",
		desc.ToString())

	var wait bool

	vc.SetRunFuncWithNameArg(func() error {
		return doCompactStatus(vc, wait)
	})
}

func doCompactStatus(vc *cmdutils.VerbCmd, wait bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if !topic.IsPersistent(){
		return errors.New("need to provide a persistent topic")
	}

	admin := cmdutils.NewPulsarClient()
	status, err := admin.Topics().CompactStatus(*topic)
	if err != nil {
		return err
	}

	for wait && status.Status == RUNNING {
		time.Sleep(1 * time.Second)
		status, err = admin.Topics().CompactStatus(*topic)
		if err != nil {
			return err
		}
	}

	switch status.Status {
	case NOT_RUN:
		vc.Command.Printf("Compacting the topic %s is not running", topic.String())
	case RUNNING:
		vc.Command.Printf("Compacting the topic %s is running", topic.String())
	case SUCCESS:
		vc.Command.Printf("Compacting the topic %s is done successfully", topic.String())
	case ERROR:
		vc.Command.Printf("Compacting the topic %s is done with error %s", topic.String(), status.LastError)
	}

	return err
}
