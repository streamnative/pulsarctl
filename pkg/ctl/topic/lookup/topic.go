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
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func LookupTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for looking up the owner broker of a topic."
	desc.CommandPermission = "This command does not require permissions. "

	var examples []Example
	lookup := Example{
		Desc:    "Lookup the owner broker of the topic (topic-name)",
		Command: "pulsarctl topic lookup (topic-name)",
	}
	desc.CommandExamples = append(examples, lookup)

	var out []Output
	successOut := Output{
		Desc: "",
		Out: "{\n" +
			"  \"brokerUlr\": \"\",\n" +
			"  \"brokerUrlTls\": \"\",\n" +
			"  \"httpUrl\": \"\",\n" +
			"  \"httpUrlTls\": \"\",\n" +
			"}",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out  = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"lookup",
		"Look up a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doLookupTopic(vc)
	})
}

func doLookupTopic(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	lookup, err := admin.Topics().Lookup(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), lookup)
	}
	return err
}
