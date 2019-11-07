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

package functions

import (
	"strconv"

	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func statsFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the current stats of a Pulsar Function."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example
	stats := cmdutils.Example{
		Desc: "Get the current stats of a Pulsar Function",
		Command: "pulsarctl functions stats \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Function)",
	}
	examples = append(examples, stats)

	statsWithFQFN := cmdutils.Example{
		Desc: "Get the current stats of a Pulsar Function with FQFN",
		Command: "pulsarctl functions stats \n" +
			"\t--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]",
	}
	examples = append(examples, statsWithFQFN)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"receivedTotal\": 0,\n" +
			"  \"processedSuccessfullyTotal\": 0,\n" +
			"  \"systemExceptionsTotal\": 0,\n" +
			"  \"userExceptionsTotal\": 0,\n" +
			"  \"avgProcessLatency\": 0,\n" +
			"  \"lastInvocation\": 0,\n" +
			"  \"oneMin\": {\n" +
			"    \"receivedTotal\": 0,\n" +
			"    \"processedSuccessfullyTotal\": 0,\n" +
			"    \"systemExceptionsTotal\": 0,\n" +
			"    \"userExceptionsTotal\": 0,\n" +
			"    \"avgProcessLatency\": 0\n" +
			"  },\n" +
			"  \"instances\": [\n" +
			"    {\n" +
			"      \"receivedTotal\": 0,\n" +
			"      \"processedSuccessfullyTotal\": 0,\n" +
			"      \"systemExceptionsTotal\": 0,\n" +
			"      \"userExceptionsTotal\": 0,\n" +
			"      \"avgProcessLatency\": 0,\n" +
			"      \"instanceId\": 0,\n" +
			"      \"metrics\": {\n" +
			"        \"oneMin\": {\n" +
			"          \"receivedTotal\": 0,\n" +
			"          \"processedSuccessfullyTotal\": 0,\n" +
			"          \"systemExceptionsTotal\": 0,\n" +
			"          \"userExceptionsTotal\": 0,\n" +
			"          \"avgProcessLatency\": 0\n" +
			"        },\n" +
			"        \"lastInvocation\": 0,\n" +
			"        \"userMetrics\": {},\n" +
			"        \"receivedTotal\": 0,\n" +
			"        \"processedSuccessfullyTotal\": 0,\n" +
			"        \"systemExceptionsTotal\": 0,\n" +
			"        \"userExceptionsTotal\": 0,\n" +
			"        \"avgProcessLatency\": 0\n" +
			"      }\n" +
			"    }\n" +
			"  ],\n" +
			"  \"instanceId\": 0,\n" +
			"  \"metrics\": {\n" +
			"    \"oneMin\": {\n" +
			"      \"receivedTotal\": 0,\n" +
			"      \"processedSuccessfullyTotal\": 0,\n" +
			"      \"systemExceptionsTotal\": 0,\n" +
			"      \"userExceptionsTotal\": 0,\n" +
			"      \"avgProcessLatency\": 0\n" +
			"    },\n" +
			"    \"lastInvocation\": 0,\n" +
			"    \"userMetrics\": null,\n" +
			"    \"receivedTotal\": 0,\n" +
			"    \"processedSuccessfullyTotal\": 0,\n" +
			"    \"systemExceptionsTotal\": 0,\n" +
			"    \"userExceptionsTotal\": 0,\n" +
			"    \"avgProcessLatency\": 0\n" +
			"  }\n" +
			"}",
	}

	failOut := cmdutils.Output{
		Desc: "You must specify a name for the Pulsar Functions or a FQFN, please check the --name args",
		Out:  "[✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN)",
	}

	failOutWithNameNotExist := cmdutils.Output{
		Desc: "The name of Pulsar Functions doesn't exist, please check the --name args",
		Out:  "[✖]  code: 404 reason: Function (your function name) doesn't exist",
	}

	out = append(out, successOut, failOut, failOutWithNameNotExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"stats",
		"Get the current stats of a Pulsar Function",
		desc.ToString(),
		desc.ExampleToString(),
		"stats",
	)

	functionData := &utils.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doStatsFunction(vc, functionData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("FunctionsConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&functionData.FQFN,
			"fqfn",
			"",
			"The Fully Qualified Function Name (FQFN) for the function")

		flagSet.StringVar(
			&functionData.Tenant,
			"tenant",
			"",
			"The tenant of a Pulsar Function")

		flagSet.StringVar(
			&functionData.Namespace,
			"namespace",
			"",
			"The namespace of a Pulsar Function")

		flagSet.StringVar(
			&functionData.FuncName,
			"name",
			"",
			"The name of a Pulsar Function")

		flagSet.StringVar(
			&functionData.InstanceID,
			"instance-id",
			"",
			"The function instanceId (Get-stats of all instances if instance-id is not provided)")
	})
}

func doStatsFunction(vc *cmdutils.VerbCmd, funcData *utils.FunctionData) error {
	err := processBaseArguments(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	if funcData.InstanceID != "" {
		instanceID, err := strconv.Atoi(funcData.InstanceID)
		if err != nil {
			return err
		}
		functionInstanceStatsData, err := admin.Functions().GetFunctionStatsWithInstanceID(
			funcData.Tenant, funcData.Namespace, funcData.FuncName, instanceID)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		}
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), functionInstanceStatsData)
	} else {
		functionStats, err := admin.Functions().GetFunctionStats(funcData.Tenant, funcData.Namespace, funcData.FuncName)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		}

		cmdutils.PrintJSON(vc.Command.OutOrStdout(), functionStats)
	}

	return err
}
