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

package topic

import (
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func ListTopicsCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for listing all exist topics under the specified namespace."
	desc.CommandPermission = "This command requires admin permissions."
	desc.CommandScope = "non-partitioned topic, partitioned topic"

	listTopics := cmdutils.Example{
		Desc:    "List all exist topics under the namespace(tenant/namespace)",
		Command: "pulsarctl topics list (tenant/namespace)",
	}
	desc.CommandExamples = []cmdutils.Example{listTopics}

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: `+----------------------------------------------------------+---------------+
|                        TOPIC NAME                        | PARTITIONED ? |
+----------------------------------------------------------+---------------+
+----------------------------------------------------------+---------------+`,
	}

	argError := cmdutils.Output{
		Desc: "the namespace is not specified",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}
	out = append(out, successOut, argError, TenantNotExistError, NamespaceNotExistError)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all exist topics under the specified namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"lp")

	vc.SetRunFuncWithNameArg(func() error {
		return doListTopics(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.EnableOutputFlagSet()
}

func doListTopics(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	namespace, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	partitionedTopics, nonPartitionedTopics, err := admin.Topics().List(*namespace)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		return err
	}

	oc := cmdutils.NewOutputContent().
		WithObject(listOutput{partitionedTopics, nonPartitionedTopics}).
		WithTextFunc(func(w io.Writer) error {
			table := tablewriter.NewWriter(w)
			table.SetHeader([]string{"topic name", "partitioned ?"})

			for _, v := range partitionedTopics {
				table.Append([]string{v, "Y"})
			}

			for _, v := range nonPartitionedTopics {
				table.Append([]string{v, "N"})
			}
			table.Render()
			return nil
		})
	err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)

	return err
}

type listOutput struct {
	PartitionedTopics    []string `json:"partitionedTopics"`
	NonPartitionedTopics []string `json:"nonPartitionedTopics"`
}
