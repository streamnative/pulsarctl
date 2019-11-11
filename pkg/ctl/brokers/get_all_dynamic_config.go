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

package brokers

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getAllDynamicConfigsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get all overridden dynamic-configuration values"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "Get all overridden dynamic-configuration values",
		Command: "pulsarctl brokers get-all-dynamic-config",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"dispatchThrottlingRatePerTopicInMsg\": \"true\"\n" +
			"}",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-all-dynamic-config",
		"Get all overridden dynamic-configuration values",
		desc.ToString(),
		desc.ExampleToString(),
		"get-all-dynamic-config")

	vc.SetRunFunc(func() error {
		return doGetAllDynamicConfigs(vc)
	})
}

func doGetAllDynamicConfigs(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	brokersData, err := admin.Brokers().GetAllDynamicConfigurations()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), brokersData)
	}
	return err
}
