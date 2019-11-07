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
)

func LookUpTopicCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for looking up the owner broker of a topic."
	desc.CommandPermission = "This command does not require permissions. "
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic, partitioned topic"

	var examples []cmdutils.Example
	lookup := cmdutils.Example{
		Desc:    "Lookup the owner broker of the topic (topic-name)",
		Command: "pulsarctl topic lookup (topic-name)",
	}
	examples = append(examples, lookup)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
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
	out = append(out, NamespaceErrors...)
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

	topic, err := utils.GetTopicName(vc.NameArg)
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
