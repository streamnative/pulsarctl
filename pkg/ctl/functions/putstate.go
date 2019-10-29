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
	"io/ioutil"
	"strings"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func putstateFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Put a key/value pair to the state associated with a Pulsar Function."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []pulsar.Example
	putstate := pulsar.Example{
		Desc: "Put a key/(string value) pair to the state associated with a Pulsar Function",
		Command: "pulsarctl functions putstate \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Function) \n" +
			"\t(key name) - (string value) ",
	}
	examples = append(examples, putstate)

	putstateWithByte := pulsar.Example{
		Desc: "Put a key/(file path) pair to the state associated with a Pulsar Function",
		Command: "pulsarctl functions putstate \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Function) \n" +
			"\t(key name) = (file path) ",
	}
	examples = append(examples, putstateWithByte)

	putstateWithFQFN := pulsar.Example{
		Desc: "Put a key/value pair to the state associated with a Pulsar Function with FQFN",
		Command: "pulsarctl functions putstate \n" +
			"\t--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions] \n" +
			"\t(key name) - (string value) ",
	}
	examples = append(examples, putstateWithFQFN)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Put state (the function state) successfully",
	}

	failOut := pulsar.Output{
		Desc: "You must specify a name for the Pulsar Functions or a FQFN, please check the --name args",
		Out:  "[✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN)",
	}

	failOutWithNameNotExist := pulsar.Output{
		Desc: "The name of Pulsar Functions doesn't exist, please check the `--name` arg",
		Out:  "[✖]  code: 404 reason: Function (your function name) doesn't exist",
	}

	failOutWithKeyOrValueNotExist := pulsar.Output{
		Desc: "The state key and state value not specified, please check your input format",
		Out:  "[✖]  need to specified the state key and state value",
	}

	fileOutErrInputFormat := pulsar.Output{
		Desc: "The format of the input is incorrect, please check.",
		Out:  "[✖]  error input format",
	}

	out = append(out, successOut, failOut, failOutWithNameNotExist, failOutWithKeyOrValueNotExist, fileOutErrInputFormat)
	desc.CommandOutput = out

	vc.SetDescription(
		"putstate",
		"Put a key/value pair to the state associated with a Pulsar Function",
		desc.ToString(),
		desc.ExampleToString(),
		"putstate",
	)

	functionData := &pulsar.FunctionData{}

	// set the run function
	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doPutStateFunction(vc, functionData)
	}, checkPutStateArgs)

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

func doPutStateFunction(vc *cmdutils.VerbCmd, funcData *pulsar.FunctionData) error {
	err := processBaseArguments(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V3)

	var state pulsar.FunctionState

	state.Key = vc.NameArgs[0]
	value := vc.NameArgs[1]

	switch value {
	case "-":
		state.StringValue = strings.Join(vc.NameArgs[2:], " ")
	case "=":
		contents, err := ioutil.ReadFile(vc.NameArgs[2])
		if err != nil {
			return err
		}
		state.ByteValue = contents
	default:
		return errors.New("error input format")
	}

	err = admin.Functions().PutFunctionState(funcData.Tenant, funcData.Namespace, funcData.FuncName, state)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Put state %+v successfully\n", state)
	}

	return err
}
