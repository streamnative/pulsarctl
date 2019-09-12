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

package sinks

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func listSinksCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "List all running Pulsar IO sink connectors"
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []pulsar.Example

	list := pulsar.Example{
		Desc: "List all running Pulsar IO sink connectors",
		Command: "pulsarctl sink list \n" +
			"\t--tenant public\n" +
			"\t--namespace default",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "+--------------------+\n" +
			"|   Sink Name    |\n" +
			"+--------------------+\n" +
			"| test_sink_name |\n" +
			"+--------------------+",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all running Pulsar IO sink connectors",
		desc.ToString(),
		"list",
	)

	sinkData := &pulsar.SinkData{}

	// set the run sink
	vc.SetRunFunc(func() error {
		return doListSinks(vc, sinkData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("SinksConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&sinkData.Tenant,
			"tenant",
			"",
			"The sink's tenant")

		flagSet.StringVar(
			&sinkData.Namespace,
			"namespace",
			"",
			"The sink's namespace")
	})
}

func doListSinks(vc *cmdutils.VerbCmd, sinkData *pulsar.SinkData) error {
	processNamespaceCmd(sinkData)

	admin := cmdutils.NewPulsarClientWithApiVersion(pulsar.V3)
	sinks, err := admin.Sinks().ListSinks(sinkData.Tenant, sinkData.Namespace)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Pulsar Sinks Name"})

		for _, f := range sinks {
			table.Append([]string{f})
		}

		table.Render()
	}
	return err
}
