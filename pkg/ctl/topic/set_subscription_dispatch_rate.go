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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func SetSubscriptionDispatchRateCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set subscription message dispatch rate for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc: "Set subscription message dispatch rate for a topic",
		Command: "pulsarctl topics set-subscription-dispatch-rate topic " +
			"--msg-dispatch-rate 4 --byte-dispatch-rate 5 --dispatch-rate-period 6 --relative-to-publish-rate",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set subscription message dispatch rate successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, TopicLevelPolicyNotEnabledError)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-subscription-dispatch-rate",
		"Set subscription message dispatch rate for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-subscription-dispatch-rate",
	)
	dispatchRateData := &utils.DispatchRateData{}
	vc.SetRunFuncWithNameArg(func() error {
		return doSetSubscriptionDispatchRate(vc, dispatchRateData)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("SubscriptionDispatchRate", func(set *pflag.FlagSet) {
		set.Int64VarP(
			&dispatchRateData.DispatchThrottlingRateInMsg,
			"msg-dispatch-rate",
			"",
			-1,
			"message-dispatch-rate (defaults to -1 and overwrites the existing value when omitted)")
		set.Int64VarP(
			&dispatchRateData.DispatchThrottlingRateInByte,
			"byte-dispatch-rate",
			"",
			-1,
			"byte-dispatch-rate (defaults to -1 and overwrites the existing value when omitted)")
		set.Int64VarP(
			&dispatchRateData.RatePeriodInSecond,
			"dispatch-rate-period",
			"",
			1,
			"dispatch-rate-period in second type (defaults to 1 second and overwrites the existing value when omitted)")
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

func doSetSubscriptionDispatchRate(vc *cmdutils.VerbCmd, dispatchRateData *utils.DispatchRateData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	client, err := cmdutils.NewPulsarRESTClientWithAPIVersion(config.V2)
	if err != nil {
		return err
	}

	endpoint := cmdutils.BuildAdminEndpoint(config.V2, "/persistent", topic.GetRestPath(), "subscriptionDispatchRate")
	err = client.Post(endpoint, dispatchRateData)
	if err == nil {
		vc.Command.Printf("Set subscription message dispatch rate successfully for [%s]\n", topic.String())
	}
	return err
}
