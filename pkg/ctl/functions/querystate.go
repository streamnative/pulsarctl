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
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"time"
)

func querystateFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Fetch the current state associated with a Pulsar Function."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	querystate := pulsar.Example{
		Desc: "Fetch the current state associated with a Pulsar Function",
		Command: "pulsarctl functions querystate \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name <the name of Pulsar Function> \n" +
			"\t--key <the name of key> \n" +
			"\t--watch",
	}
	examples = append(examples, querystate)

	querystateWithFQFN := pulsar.Example{
		Desc: "Fetch the current state associated with a Pulsar Function with FQFN",
		Command: "pulsarctl functions querystate \n" +
			"\t--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]\n" +
			"\t--key <the name of key> \n" +
			"\t--watch",
	}
	examples = append(examples, querystateWithFQFN)

	querystateNoWatch := pulsar.Example{
		Desc: "Fetch the current state associated with a Pulsar Function",
		Command: "pulsarctl functions querystate \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name <the name of Pulsar Function> \n" +
			"\t--key <the name of key> ",
	}
	examples = append(examples, querystateNoWatch)

	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"key\": \"pulsar\",\n" +
			"  \"stringValue\": \"hello\",\n" +
			"  \"byteValue\": null,\n" +
			"  \"numberValue\": 0,\n" +
			"  \"version\": 6\n" +
			"}",
	}

	failOut := pulsar.Output{
		Desc: "You must specify a name for the Pulsar Functions or a FQFN, please check the --name args",
		Out:  "[✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN)",
	}

	failOutWithNameNotExist := pulsar.Output{
		Desc: "The name of Pulsar Functions doesn't exist, please check the --name args",
		Out:  "[✖]  code: 404 reason: Function <your function name> doesn't exist",
	}

	failOutWithKeyNotExist := pulsar.Output{
		Desc: "key <the name of key> doesn't exist, please check --key args",
		Out:  "error: key <the name of key> doesn't exist",
	}

	out = append(out, successOut, failOut, failOutWithNameNotExist, failOutWithKeyNotExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"querystate",
		"Fetch the current state associated with a Pulsar Function",
		desc.ToString(),
		"querystate",
	)

	functionData := &pulsar.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doQueryStateFunction(vc, functionData)
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

		flagSet.StringVarP(
			&functionData.Key,
			"key",
			"k",
			"",
			"key")

		flagSet.BoolVarP(
			&functionData.Watch,
			"watch",
			"w",
			false,
			"Watch for changes in the value associated with a key for a Pulsar Function")
	})
}

func doQueryStateFunction(vc *cmdutils.VerbCmd, funcData *pulsar.FunctionData) error {
	err := processBaseArguments(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithApiVersion(pulsar.V3)

	for {
		functionState, err := admin.Functions().GetFunctionState(funcData.Tenant, funcData.Namespace, funcData.FuncName, funcData.Key)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			cmdutils.PrintJson(vc.Command.OutOrStdout(), functionState)
		}

		if funcData.Watch {
			time.Sleep(time.Millisecond * 1000)
		}

		if !funcData.Watch {
			break
		}
	}
	return err
}
