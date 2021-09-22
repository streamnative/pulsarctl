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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
)

func listPackageVersionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "List all versions of a package"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc: "List all versions of a package",
		Command: "pulsarctl packages list-versions \n" +
			"\tfunction://public/default/example",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "+--------------------+\n" +
			"|   Package Version    |\n" +
			"+--------------------+\n" +
			"| function://public/default/example@v0.1 |\n" +
			"+--------------------+",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list-versions",
		"List all versions of a package",
		desc.ToString(),
		desc.ExampleToString(),
		"list-versions",
	)

	// set the run function
	vc.SetRunFuncWithNameArg(func() error {
		return doListPackageVersions(vc)
	}, "the package URL is not provided")

	vc.EnableOutputFlagSet()
}

func doListPackageVersions(vc *cmdutils.VerbCmd) error {
	_, err := utils.GetPackageName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	packages, err := admin.Packages().ListVersions(vc.NameArg)
	if err != nil {
		return err
	}

	oc := cmdutils.NewOutputContent().
		WithObject(packages).
		WithTextFunc(func(w io.Writer) error {
			table := tablewriter.NewWriter(w)
			table.SetHeader([]string{"Pulsar Package Version"})

			for _, f := range packages {
				table.Append([]string{f})
			}

			table.Render()
			return nil
		})
	err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)

	return err
}
