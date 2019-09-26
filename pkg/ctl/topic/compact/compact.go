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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
)

func CompactCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for compacting a persistent topic."
	desc.CommandPermission = "This command is requires tenant admin permissions."

	var examples []pulsar.Example
	compact := pulsar.Example{
		Desc:    "Compact a persistent topic <topic-name>",
		Command: "pulsarctl topic compact <topic-name>",
	}
	desc.CommandExamples = append(examples, compact)

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Sending compact topic <topic-name> request successfully",
	}
	out = append(out, successOut, e.ArgError, e.TopicNotFoundError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"compact",
		"Compact a topic",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doCompact(vc)
	})
}

func doCompact(vc *cmdutils.VerbCmd) error {
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
	err = admin.Topics().Compact(*topic)
	if err == nil {
		vc.Command.Printf("Sending compact topic %s request successfully/n", topic.String())
	}

	return err
}
