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

func getRuntimeConfigCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get runtime configuration values"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "Get runtime configuration values",
		Command: "pulsarctl brokers get-runtime-config",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  “activeConsumerFailoverDelayTimeMillis”: “1000\",\n" +
			"  “advertisedAddress”: “127.0.0.1\",\n" +
			"  “allowAutoTopicCreation”: “true”,\n" +
			"  “anonymousUserRole”: “”,\n" +
			"  “authenticateOriginalAuthData”: “false”,\n" +
			"  “authenticationEnabled”: “false”,\n" +
			"  ...\n" +
			"}",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-runtime-config",
		"Get runtime configuration values",
		desc.ToString(),
		desc.ExampleToString(),
		"get-runtime-config")

	vc.SetRunFunc(func() error {
		return doGetRuntimeConfig(vc)
	})
}

func doGetRuntimeConfig(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	brokersData, err := admin.Brokers().GetRuntimeConfigurations()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), brokersData)
	}
	return err
}
