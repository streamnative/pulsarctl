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

package lookup

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func TopicCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for looking up the owner broker of a topic."
	desc.CommandPermission = "This command does not require permissions. "

	var examples []pulsar.Example
	lookup := pulsar.Example{
		Desc:    "Lookup the owner broker of the topic (topic-name)",
		Command: "pulsarctl topic lookup (topic-name)",
	}
	examples = append(examples, lookup)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "",
		Out: "{\n" +
			"  \"brokerUlr\": \"\",\n" +
			"  \"brokerUrlTls\": \"\",\n" +
			"  \"httpUrl\": \"\",\n" +
			"  \"httpUrlTls\": \"\",\n" +
			"}",
	}
	out = append(out, successOut, e.ArgError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"lookup",
		"Look up a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doLookupTopic(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doLookupTopic(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	lookup, err := admin.Topics().Lookup(*topic)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), lookup)
	}
	return err
}
