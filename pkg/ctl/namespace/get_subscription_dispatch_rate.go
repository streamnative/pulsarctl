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

package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func GetSubscriptionDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the default subscription message dispatch rate of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the default subscription message dispatch rate of namespace (namespace-name)",
		Command: "pulsarctl namespaces get-subscription-dispatch-rate (namespace-name)",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"dispatchThrottlingRateInMsg\" : 0,\n" +
			"  \"dispatchThrottlingRateInByte\" : 0,\n" +
			"  \"ratePeriodInSecond\" : 1\n" +
			"}",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-subscription-dispatch-rate",
		"Get the default subscription message dispatch rate of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetSubscriptionDispatchRate(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetSubscriptionDispatchRate(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	rate, err := admin.Namespaces().GetSubscriptionDispatchRate(*ns)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), rate)
	}

	return err
}
