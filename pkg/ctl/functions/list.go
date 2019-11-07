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

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/pflag"
)

func listFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "List all Pulsar Functions running under a specific tenant and namespace."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc: "List all Pulsar Functions running under a specific tenant and namespace",
		Command: "pulsarctl functions list \n" +
			"\t--tenant public\n" +
			"\t--namespace default",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "+--------------------+\n" +
			"|   Function Name    |\n" +
			"+--------------------+\n" +
			"| test_function_name |\n" +
			"+--------------------+",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all Pulsar Functions running under a specific tenant and namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"list",
	)

	functionData := &utils.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doListFunctions(vc, functionData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("FunctionsConfig", func(flagSet *pflag.FlagSet) {
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
	})
}

func doListFunctions(vc *cmdutils.VerbCmd, funcData *utils.FunctionData) error {
	processNamespaceCmd(funcData)

	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	functions, err := admin.Functions().GetFunctions(funcData.Tenant, funcData.Namespace)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Pulsar Function Name"})

		for _, f := range functions {
			table.Append([]string{f})
		}

		table.Render()
	}
	return err
}
