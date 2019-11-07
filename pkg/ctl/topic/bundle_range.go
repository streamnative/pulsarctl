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

func GetBundleRangeCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting namespace bundle range of a topic (partition)."
	desc.CommandPermission = "This command requires super-user permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic"

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get namespace bundle range of a topic (topic-name)",
		Command: "pulsarctl topic bundle-range (topic-name)",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The bundle range of the topic (topic-name) is: (bundle-range)",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"bundle-range",
		"Get the namespace bundle range of a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetBundleRange(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doGetBundleRange(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	bundleRange, err := admin.Topics().GetBundleRange(*topic)
	if err == nil {
		vc.Command.Printf("The bundle range of the topic %s is: %s\n", topic.String(), bundleRange)
	}

	return err
}
