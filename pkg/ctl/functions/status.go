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

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func statusFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Check the current status of a Pulsar Function."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example
	status := cmdutils.Example{
		Desc: "Check the current status of a Pulsar Function",
		Command: "pulsarctl functions status \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Function)",
	}
	examples = append(examples, status)

	statusWithFQFN := cmdutils.Example{
		Desc: "Check the current status of a Pulsar Function with FQFN",
		Command: "pulsarctl functions status \n" +
			"\t--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]",
	}
	examples = append(examples, statusWithFQFN)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"numInstances\": 1,\n" +
			"  \"numRunning\": 1,\n" +
			"  \"instances\": [\n" +
			"    {\n" +
			"      \"instanceId\": 0,\n" +
			"      \"status\": {\n" +
			"        \"running\": true,\n" +
			"        \"error\": \"\",\n" +
			"        \"numRestarts\": 0,\n" +
			"        \"numReceived\": 0,\n" +
			"        \"numSuccessfullyProcessed\": 0,\n" +
			"        \"numUserExceptions\": 0,\n" +
			"        \"latestUserExceptions\": [],\n" +
			"        \"numSystemExceptions\": 0,\n" +
			"        \"latestSystemExceptions\": [],\n" +
			"        \"averageLatency\": 0,\n" +
			"        \"lastInvocationTime\": 0,\n" +
			"        \"workerId\": \"c-standalone-fw-127.0.0.1-8080\"\n" +
			"      }\n" +
			"    }\n" +
			"  ]\n" +
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

	failOutWithWrongInstanceID := cmdutils.Output{
		Desc: "Used an instanceID that does not exist or other impermissible actions",
		Out:  "[✖]  code: 400 reason: Operation not permitted",
	}

	out = append(out, successOut, failOut, failOutWithNameNotExist, failOutWithWrongInstanceID)
	desc.CommandOutput = out

	vc.SetDescription(
		"status",
		"Check the current status of a Pulsar Function",
		desc.ToString(),
		desc.ExampleToString(),
		"getstatus",
	)

	functionData := &utils.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doStatusFunction(vc, functionData)
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
			"The function instanceId (Get-status of all instances if instance-id is not provided)")
	})
}

func doStatusFunction(vc *cmdutils.VerbCmd, funcData *utils.FunctionData) error {
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
		functionInstanceStatusData, err := admin.Functions().GetFunctionStatusWithInstanceID(
			funcData.Tenant, funcData.Namespace, funcData.FuncName, instanceID)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		}
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), functionInstanceStatusData)
	} else {
		functionStatus, err := admin.Functions().GetFunctionStatus(funcData.Tenant, funcData.Namespace, funcData.FuncName)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		}
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), functionStatus)
	}

	return err
}
