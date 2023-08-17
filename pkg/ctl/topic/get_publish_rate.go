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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetPublishRateCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get message publish rate for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Get message publish rate for a topic",
		Command: "pulsarctl topics get-publish-rate topic",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Get message publish rate successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, TopicLevelPolicyNotEnabledError)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-publish-rate",
		"Get message publish rate for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"get-publish-rate",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetPublishRate(vc)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.EnableOutputFlagSet()
}

func doGetPublishRate(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	publishRateData, err := admin.Topics().GetPublishRate(*topic)
	if err == nil {
		oc := cmdutils.NewOutputContent().WithObject(publishRateData)
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
	}
	return err
}
