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
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func restartFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "This command is used for restarting function instance."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example

	restart := pulsar.Example{
		Desc: "Restart function instance",
		Command: "pulsarctl functions restart \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Function)",
	}
	examples = append(examples, restart)

	restartWithInstanceID := pulsar.Example{
		Desc: "Restart function instance with instance ID",
		Command: "pulsarctl functions restart \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Function)\n" +
			"\t--instance-id 1",
	}
	examples = append(examples, restartWithInstanceID)

	restartWithFQFN := pulsar.Example{
		Desc: "Restart function instance with FQFN",
		Command: "pulsarctl functions restart \n" +
			"\t--fqfn tenant/namespace/name [eg: public/default/ExampleFunctions]",
	}
	examples = append(examples, restartWithFQFN)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Restarted (the name of a Pulsar Function) successfully",
	}

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

	out = append(out, successOut, failOut, failOutWithNameNotExist, failOutWithWrongInstanceID)
	desc.CommandOutput = out

	vc.SetDescription(
		"restart",
		"Restart function instance",
		desc.ToString(),
		desc.ExampleToString(),
		"restart",
	)

	functionData := &pulsar.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doRestartFunctions(vc, functionData)
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
			"The function instanceId (restart all instances if instance-id is not provided)")
	})
}

func doRestartFunctions(vc *cmdutils.VerbCmd, funcData *pulsar.FunctionData) error {
	err := processBaseArguments(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V3)
	if funcData.InstanceID != "" {
		instanceID, err := strconv.Atoi(funcData.InstanceID)
		if err != nil {
			return err
		}
		err = admin.Functions().RestartFunctionWithID(funcData.Tenant, funcData.Namespace, funcData.FuncName, instanceID)
		if err != nil {
			return err
		}
		vc.Command.Printf("Restarted %s successfully\n", funcData.FuncName)
	} else {
		err = admin.Functions().RestartFunction(funcData.Tenant, funcData.Namespace, funcData.FuncName)
		if err != nil {
			return err
		}

		vc.Command.Printf("Restarted %s successfully\n", funcData.FuncName)
	}

	return nil
}
