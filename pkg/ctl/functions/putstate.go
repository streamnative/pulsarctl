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
	"encoding/json"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func putstateFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Put a key/value pair to the state associated with a Pulsar Function."
	desc.CommandPermission = "This command requires user permissions."

	var examples []pulsar.Example
	putstate := pulsar.Example{
		Desc: "Put a key/value pair to the state associated with a Pulsar Function",
		Command: "pulsarctl functions putstate \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name <the name of Pulsar Function> \n" +
			"\t--state \"{\"key\":\"pulsar\", \"stringValue\":\"hello\"}\" ",
	}
	examples = append(examples, putstate)

	putstateWithFQFN := pulsar.Example{
		Desc: "Put a key/value pair to the state associated with a Pulsar Function with FQFN",
		Command: "pulsarctl functions putstate \n" +
			"\t--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions] \n" +
			"\t--state \"{\"key\":\"pulsar\", \"stringValue\":\"hello\"}\"",
	}
	examples = append(examples, putstateWithFQFN)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Put state <the function state> successfully",
	}

	failOut := pulsar.Output{
		Desc: "You must specify a name for the Pulsar Functions or a FQFN, please check the --name args",
		Out:  "[✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN)",
	}

	failOutWithNameNotExist := pulsar.Output{
		Desc: "The name of Pulsar Functions doesn't exist, please check the `--name` arg",
		Out:  "[✖]  code: 404 reason: Function <your function name> doesn't exist",
	}

	failOutWithWrongJson := pulsar.Output{
		Desc: "unexpected end of JSON input, please check the `--state` arg",
		Out:  "[✖]  unexpected end of JSON input",
	}

	out = append(out, successOut, failOut, failOutWithNameNotExist, failOutWithWrongJson)
	desc.CommandOutput = out

	vc.SetDescription(
		"putstate",
		"Put a key/value pair to the state associated with a Pulsar Function",
		desc.ToString(),
		"putstate",
	)

	functionData := &pulsar.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doPutStateFunction(vc, functionData)
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
			&functionData.State,
			"state",
			"t",
			"",
			"The FunctionState that needs to be put")
	})
}

func doPutStateFunction(vc *cmdutils.VerbCmd, funcData *pulsar.FunctionData) error {
	err := processBaseArguments(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithApiVersion(pulsar.V3)

	var state pulsar.FunctionState
	err = json.Unmarshal([]byte(funcData.State), &state)
	if err != nil {
		return err
	}

	err = admin.Functions().PutFunctionState(funcData.Tenant, funcData.Namespace, funcData.FuncName, state)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Put state %+v successfully", state)
	}

	return err
}
