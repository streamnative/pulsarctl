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

package packages

import (
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func listPackagesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "List all specified type packages under a specific tenant and namespace"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc: "List all the specified type packages under a namespace",
		Command: "pulsarctl packages list \n" +
			"\t--type function\n" +
			"\tpublic/default",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "+--------------------+\n" +
			"|   Package Name    |\n" +
			"+--------------------+\n" +
			"| function://public/default/example@v0.1 |\n" +
			"+--------------------+",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all the specified type packages under a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"list",
	)

	var packageTypeName string

	// set the run function
	vc.SetRunFuncWithNameArg(func() error {
		return doListPackages(vc, packageTypeName)
	}, "the package URL is not provided")

	// register the params
	vc.FlagSetGroup.InFlagSet("List packages", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&packageTypeName,
			"type",
			"",
			"",
			"function, source, sink")
	})
	vc.EnableOutputFlagSet()
}

func doListPackages(vc *cmdutils.VerbCmd, packageType string) error {
	namespace, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	packages, err := admin.Packages().List(packageType, namespace.String())
	if err != nil {
		return err
	}

	oc := cmdutils.NewOutputContent().
		WithObject(packages).
		WithTextFunc(func(w io.Writer) error {
			table := tablewriter.NewWriter(w)
			table.SetHeader([]string{"Pulsar Package Name"})

			for _, f := range packages {
				table.Append([]string{f})
			}

			table.Render()
			return nil
		})
	err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)

	return err
}
