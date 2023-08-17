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

package sources

import (
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func listBuiltInSourcesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the list of Pulsar IO connector sources supported by Pulsar cluster"
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc:    "Get the list of Pulsar IO connector sources supported by Pulsar cluster",
		Command: "pulsarctl source available-sources",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "+---------+----------+-----------+\n" +
			"|   Name   |   Desc   |   Class   |\n" +
			"+----------+----------+-----------+\n" +
			"| source_name | example source | aaa.bbb |\n" +
			"+----------+----------+-----------+",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"available-sources",
		"List Pulsar IO connector sources supported by Pulsar cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"available-sources",
	)

	// set the run source
	vc.SetRunFunc(func() error {
		return doListBuiltInSources(vc)
	})

	vc.EnableOutputFlagSet()
}

func doListBuiltInSources(vc *cmdutils.VerbCmd) error {

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	connectorDefinition, err := admin.Sinks().GetBuiltInSinks()
	if err != nil {
		return err
	}

	oc := cmdutils.NewOutputContent().
		WithObject(connectorDefinition).
		WithTextFunc(func(w io.Writer) error {
			table := tablewriter.NewWriter(w)
			table.SetHeader([]string{"Name", "Description", "ClassName"})

			for _, f := range connectorDefinition {
				if f.SourceClass != "" {
					table.Append([]string{f.Name, f.Description, f.SourceClass})
				}
			}

			table.Render()
			return nil
		})
	err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)

	return err
}
