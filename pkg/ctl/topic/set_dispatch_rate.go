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
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func SetDispatchRateCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set message dispatch rate for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc: "Set message dispatch rate for a topic",
		Command: "pulsarctl topics set-dispatch-rate topic " +
			"--msg-dispatch-rate 4 --byte-dispatch-rate 5 --dispatch-rate-period 6 --relative-to-publish-rate",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set message dispatch rate successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-dispatch-rate",
		"Set message dispatch rate for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-dispatch-rate",
	)
	dispatchRateData := &utils.DispatchRateData{}
	vc.SetRunFuncWithNameArg(func() error {
		return doSetDispatchRate(vc, dispatchRateData)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("DispatchRate", func(set *pflag.FlagSet) {
		set.Int64VarP(
			&dispatchRateData.DispatchThrottlingRateInMsg,
			"msg-dispatch-rate",
			"",
			-1,
			"message-dispatch-rate (default -1 will be overwrite if not passed)")
		set.Int64VarP(
			&dispatchRateData.DispatchThrottlingRateInByte,
			"byte-dispatch-rate",
			"",
			-1,
			"byte-dispatch-rate (default -1 will be overwrite if not passed)")
		set.Int64VarP(
			&dispatchRateData.RatePeriodInSecond,
			"dispatch-rate-period",
			"",
			1,
			"dispatch-rate-period in second type (default 1 second will be overwrite if not passed)")
		set.BoolVarP(
			&dispatchRateData.RelativeToPublishRate,
			"relative-to-publish-rate",
			"",
			false,
			"dispatch rate relative to publish-rate (if publish-relative flag is enabled "+
				"then broker will apply throttling value to (publish-rate + dispatch rate))")
	})
	vc.EnableOutputFlagSet()
}

func doSetDispatchRate(vc *cmdutils.VerbCmd, dispatchRateData *utils.DispatchRateData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetDispatchRate(*topic, *dispatchRateData)
	if err == nil {
		vc.Command.Printf("Set message dispatch rate successfully for [%s]\n", topic.String())
	}
	return err
}
