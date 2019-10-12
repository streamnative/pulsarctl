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

package crud

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/olekukonko/tablewriter"
)

func ListTopicsCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for listing all exist topics under the specified namespace."
	desc.CommandPermission = "This command requires admin permissions."
	desc.CommandScope = "non-partitioned topic, partitioned topic"

	listTopics := pulsar.Example{
		Desc:    "List all exist topics under the namespace(tenant/namespace)",
		Command: "pulsarctl topics list (tenant/namespace)",
	}
	desc.CommandExamples = []pulsar.Example{listTopics}

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: `+----------------------------------------------------------+---------------+
|                        TOPIC NAME                        | PARTITIONED ? |
+----------------------------------------------------------+---------------+
+----------------------------------------------------------+---------------+`,
	}

	argError := pulsar.Output{
		Desc: "the namespace is not specified",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}
	out = append(out, successOut, argError, e.TenantNotExistError, e.NamespaceNotExistError)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all exist topics under the specified namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"lp")

	vc.SetRunFuncWithNameArg(func() error {
		return doListTopics(vc)
	})
}

func doListTopics(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	namespace, err := pulsar.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	partitionedTopics, nonPartitionedTopics, err := admin.Topics().List(*namespace)
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"topic name", "partitioned ?"})

		for _, v := range partitionedTopics {
			table.Append([]string{v, "Y"})
		}

		for _, v := range nonPartitionedTopics {
			table.Append([]string{v, "N"})
		}
		table.Render()
	}

	return err
}
