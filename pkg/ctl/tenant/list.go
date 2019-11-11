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

package tenant

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/olekukonko/tablewriter"
)

func listTenantCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for listing all the existing tenants."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	listSuccess := cmdutils.Example{
		Desc:    "list all the existing tenants",
		Command: "pulsarctl tenants list",
	}
	examples = append(examples, listSuccess)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "+-------------+\n" +
			"| TENANT NAME |\n" +
			"+-------------+\n" +
			"| public      |\n" +
			"| sample      |\n" +
			"+-------------+",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all exist tenants",
		desc.ToString(),
		desc.ExampleToString(),
		"l")

	vc.SetRunFunc(func() error {
		return doListTenant(vc)
	})
}

func doListTenant(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	tenants, err := admin.Tenants().List()
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Tenant Name"})

		for _, t := range tenants {
			table.Append([]string{t})
		}

		table.Render()
	}
	return err
}
