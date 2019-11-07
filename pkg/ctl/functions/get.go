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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func getFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Fetch information about a Pulsar Function"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	get := cmdutils.Example{
		Desc: "Fetch information about a Pulsar Function",
		Command: "pulsarctl functions get \n" +
			"\t--tenant public\n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Function)",
	}

	getWithFqfn := cmdutils.Example{
		Desc: "Fetch information about a Pulsar Function with FQFN",
		Command: "pulsarctl functions get \n" +
			"\t--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]",
	}
	examples = append(examples, get, getWithFqfn)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n  " +
			"\"tenant\": \"public\",\n  " +
			"\"namespace\": \"default\",\n  " +
			"\"name\": \"test-functions\",\n  " +
			"\"className\": \"org.apache.pulsar.functions.api.examples.ExclamationFunction\",\n  " +
			"\"inputSpecs\": {\n    \"persistent://public/default/test-topic-1\": {\n      " +
			"\"isRegexPattern\": false\n    " +
			"}\n  " +
			"},\n  " +
			"\"output\": \"persistent://public/default/test-topic-2\",\n  " +
			"\"processingGuarantees\": \"ATLEAST_ONCE\",\n  " +
			"\"retainOrdering\": false,\n  \"userConfig\": {},\n  " +
			"\"runtime\": \"JAVA\",\n  \"autoAck\": true,\n  " +
			"\"parallelism\": 1,\n  \"resources\": {\n    " +
			"\"cpu\": 1.0,\n    \"ram\": 1073741824,\n    \"disk\": 10737418240\n  },\n  " +
			"\"cleanupSubscription\": true\n}",
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
		"get",
		"Fetch information about a Pulsar Function",
		desc.ToString(),
		desc.ExampleToString(),
		"get",
	)

	functionData := &utils.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doGetFunctions(vc, functionData)
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
	})
}

func doGetFunctions(vc *cmdutils.VerbCmd, funcData *utils.FunctionData) error {
	err := processBaseArguments(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	functionConfig, err := admin.Functions().GetFunction(funcData.Tenant, funcData.Namespace, funcData.FuncName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), functionConfig)
	}

	return err
}
