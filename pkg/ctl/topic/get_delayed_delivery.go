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

func GetDelayedDeliveryCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get delayed delivery policy for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Get delayed delivery policy for a topic",
		Command: "pulsarctl topics get-delayed-delivery topic",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Get delayed delivery policy successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-delayed-delivery",
		"Get delayed delivery policy for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"get-delayed-delivery",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetDelayedDelivery(vc)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.EnableOutputFlagSet()
}

func doGetDelayedDelivery(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	delayedDeliveryData, err := admin.Topics().GetDelayedDelivery(*topic)
	if err == nil {
		oc := cmdutils.NewOutputContent().WithObject(delayedDeliveryData)
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
	}
	return err
}
