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
	"errors"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func triggerFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Trigger the specified Pulsar Function with a supplied value."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []pulsar.Example
	trigger := pulsar.Example{
		Desc: "Trigger the specified Pulsar Function with a supplied value",
		Command: "pulsarctl functions trigger \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Function) \n" +
			"\t--topic (the name of input topic)\n" +
			"\t--trigger-value \"hello pulsar\"",
	}
	examples = append(examples, trigger)

	triggerWithFQFN := pulsar.Example{
		Desc: "Trigger the specified Pulsar Function with a supplied value",
		Command: "pulsarctl functions trigger \n" +
			"\t--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]\n" +
			"\t--topic (the name of input topic)\n" +
			"\t--trigger-value \"hello pulsar\"",
	}
	examples = append(examples, triggerWithFQFN)

	triggerWithFile := pulsar.Example{
		Desc: "Trigger the specified Pulsar Function with a supplied value",
		Command: "pulsarctl functions trigger \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Function) \n" +
			"\t--topic (the name of input topic)\n" +
			"\t--trigger-file (the path of trigger file)",
	}
	examples = append(examples, triggerWithFile)

	desc.CommandExamples = examples

	var out []pulsar.Output
	failOut := pulsar.Output{
		Desc: "You must specify a name for the Pulsar Functions or a FQFN, please check the --name args",
		Out:  "[✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN)",
	}

	failOutWithNameNotExist := pulsar.Output{
		Desc: "The name of Pulsar Functions doesn't exist, please check the --name args",
		Out:  "[✖]  code: 404 reason: Function (your function name) doesn't exist",
	}

	failOutWithWrongInstanceID := pulsar.Output{
		Desc: "Used an instanceID that does not exist or other impermissible actions",
		Out:  "[✖]  code: 400 reason: Operation not permitted",
	}

	failOutWithTopic := pulsar.Output{
		Desc: "Function in trigger function has unidentified topic",
		Out:  "[✖]  code: 400 reason: Function in trigger function has unidentified topic",
	}

	failOutWithTimeout := pulsar.Output{
		Desc: "Request Timed Out",
		Out:  "[✖]  code: 408 reason: Request Timed Out",
	}

	out = append(out, failOut, failOutWithNameNotExist, failOutWithWrongInstanceID, failOutWithTopic, failOutWithTimeout)
	desc.CommandOutput = out

	vc.SetDescription(
		"trigger",
		"Trigger the specified Pulsar Function with a supplied value",
		desc.ToString(),
		desc.ExampleToString(),
		"trigger",
	)

	functionData := &pulsar.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doTriggerFunction(vc, functionData)
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
			&functionData.Topic,
			"topic",
			"",
			"The specific topic name that the function consumes from that you want to inject the data to")

		flagSet.StringVar(
			&functionData.TriggerFile,
			"trigger-file",
			"",
			"The path to the file that contains the data with which you want to trigger the function")

		flagSet.StringVar(
			&functionData.TriggerValue,
			"trigger-value",
			"",
			"The value with which you want to trigger the function")
	})
}

func doTriggerFunction(vc *cmdutils.VerbCmd, funcData *pulsar.FunctionData) error {
	err := processBaseArguments(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V3)

	if funcData.TriggerValue == "" && funcData.TriggerFile == "" {
		return errors.New("either a trigger value or a trigger filepath needs to be specified")
	}

	if funcData.TriggerValue != "" && funcData.TriggerFile != "" {
		return errors.New("either a triggerValue or a triggerFile needs to specified for the" +
			" function, cannot specify both")
	}

	retval, err := admin.Functions().TriggerFunction(funcData.Tenant, funcData.Namespace,
		funcData.FuncName, funcData.Topic, funcData.TriggerValue, funcData.TriggerFile)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Println(retval)
	}
	return err
}
