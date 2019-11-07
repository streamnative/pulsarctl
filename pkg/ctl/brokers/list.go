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

package brokers

import (
	"errors"

	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getBrokerListCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "List active brokers of the cluster"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "List active brokers of the cluster",
		Command: "pulsarctl brokers list (cluster-name)",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "127.0.0.1:8080",
	}

	var argsError = cmdutils.Output{
		Desc: "the cluster name is not specified or the cluster name is specified more than one",
		Out:  "[✖]  the cluster name is not specified or the cluster name is specified more than one",
	}

	out = append(out, successOut, argsError)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List active brokers of the cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"list")

	vc.SetRunFuncWithNameArg(func() error {
		return doListCluster(vc)
	}, "the cluster name is not specified or the cluster name is specified more than one")
}

func doListCluster(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg
	if clusterName == "" {
		return errors.New("should specified a cluster name")
	}

	admin := cmdutils.NewPulsarClient()
	brokersData, err := admin.Brokers().GetActiveBrokers(clusterName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Brokers List"})

		for _, c := range brokersData {
			table.Append([]string{c})
		}

		table.Render()
	}
	return err
}
