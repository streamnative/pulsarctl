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
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func SetPublishRateCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set message publish rate for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc: "Set message publish rate for a topic",
		Command: "pulsarctl topics set-publish-rate topic " +
			"--msg-publish-rate 4 --byte-publish-rate 5 --publish-rate-period 6 --relative-to-publish-rate",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set message publish rate successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, TopicLevelPolicyNotEnabledError)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-publish-rate",
		"Set message publish rate for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-publish-rate",
	)
	publishRateData := &utils.PublishRateData{}
	vc.SetRunFuncWithNameArg(func() error {
		return doSetPublishRate(vc, publishRateData)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("PublishRate", func(set *pflag.FlagSet) {
		set.Int64VarP(
			&publishRateData.PublishThrottlingRateInMsg,
			"msg-publish-rate",
			"",
			-1,
			"message-publish-rate (default -1 will be overwrite if not passed)")
		set.Int64VarP(
			&publishRateData.PublishThrottlingRateInByte,
			"byte-publish-rate",
			"",
			-1,
			"byte-publish-rate (default -1 will be overwrite if not passed)")
	})
	vc.EnableOutputFlagSet()
}

func doSetPublishRate(vc *cmdutils.VerbCmd, publishRateData *utils.PublishRateData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetPublishRate(*topic, *publishRateData)
	if err == nil {
		vc.Command.Printf("Set message publish rate successfully for [%s]\n", topic.String())
	}
	return err
}
