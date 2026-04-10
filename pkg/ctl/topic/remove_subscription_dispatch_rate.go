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

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func RemoveSubscriptionDispatchRateCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Remove subscription message dispatch rate for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Remove subscription message dispatch rate for a topic",
		Command: "pulsarctl topics remove-subscription-dispatch-rate topic",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Remove subscription message dispatch rate successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, TopicLevelPolicyNotEnabledError)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"remove-subscription-dispatch-rate",
		"Remove subscription message dispatch rate for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"remove-subscription-dispatch-rate",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveSubscriptionDispatchRate(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doRemoveSubscriptionDispatchRate(vc *cmdutils.VerbCmd) error {
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
	err = client.Delete(endpoint)
	if err == nil {
		vc.Command.Printf("Remove subscription message dispatch rate successfully for [%s]\n", topic.String())
	}
	return err
}
