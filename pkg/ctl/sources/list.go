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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func listSourcesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "List all running Pulsar IO source connectors"
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc: "List all running Pulsar IO source connectors",
		Command: "pulsarctl source list \n" +
			"\t--tenant public\n" +
			"\t--namespace default",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "+--------------------+\n" +
			"|   Source Name    |\n" +
			"+--------------------+\n" +
			"| test_source_name |\n" +
			"+--------------------+",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all running Pulsar IO source connectors",
		desc.ToString(),
		desc.ExampleToString(),
		"list",
	)

	sourceData := &utils.SourceData{}

	// set the run source
	vc.SetRunFunc(func() error {
		return doListSources(vc, sourceData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("SourcesConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&sourceData.Tenant,
			"tenant",
			"",
			"The source's tenant")

		flagSet.StringVar(
			&sourceData.Namespace,
			"namespace",
			"",
			"The source's namespace")
	})
	vc.EnableOutputFlagSet()
}

func doListSources(vc *cmdutils.VerbCmd, sourceData *utils.SourceData) error {
	processNamespaceCmd(sourceData)

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	sources, err := admin.Sources().ListSources(sourceData.Tenant, sourceData.Namespace)
	if err != nil {
		return err
	}

	oc := cmdutils.NewOutputContent().
		WithObject(sources).
		WithTextFunc(func(w io.Writer) error {
			table := tablewriter.NewWriter(w)
			table.SetHeader([]string{"Pulsar Sources Name"})

			for _, f := range sources {
				table.Append([]string{f})
			}

			table.Render()
			return nil
		})
	err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)

	return err
}
