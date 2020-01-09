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

package context

import (
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/context/internal"
)

func getContextsCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Describe one or many contexts"
	desc.CommandPermission = "This command does not need any permission"

	var examples []cmdutils.Example
	listAllContexts := cmdutils.Example{
		Desc:    "List all the contexts in your pulsarconfig file",
		Command: "pulsarctl config get-contexts",
	}

	getOneContext := cmdutils.Example{
		Desc:    "Describe one context in your pulsarconfig file",
		Command: "pulsarctl context get",
	}

	examples = append(examples, listAllContexts, getOneContext)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "+---------------+\n" +
			"|    NAME       |\n" +
			"+---------------+\n" +
			"| test-pulsar |\n" +
			"+---------------+",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"get",
		"Describe one or many contexts",
		desc.ToString(),
		desc.ExampleToString(),
		"get")

	ops := new(getContextOptions)
	ops.access = internal.NewDefaultPathOptions()

	// set the run function with name argument
	vc.SetRunFunc(func() error {
		return doRunGetContext(vc, ops)
	})

}

type getContextOptions struct {
	access internal.ConfigAccess
}

func doRunGetContext(vc *cmdutils.VerbCmd, ops *getContextOptions) error {
	config, err := ops.access.GetStartingConfig()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(vc.Command.OutOrStdout())
	columnNames := []string{"CURRENT", "NAME", "BROKER SERVICE URL", "BOOKIE SERVICE URL"}
	table.SetHeader(columnNames)

	for name, ctx := range config.Contexts {
		if name == config.CurrentContext {
			table.Append([]string{"*", name, ctx.BrokerServiceURL, ctx.BookieServiceURL})
		} else {
			table.Append([]string{"", name, ctx.BrokerServiceURL, ctx.BookieServiceURL})
		}
	}

	table.Render()
	return nil
}
