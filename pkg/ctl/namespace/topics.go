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

package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/olekukonko/tablewriter"
)

func getTopics(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Get the list of topics for a namespace"
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example

	topics := pulsar.Example{
		Desc:    "Get the list of topics for a namespace",
		Command: "pulsarctl namespaces topics (tenant/namespace)",
	}

	examples = append(examples, topics)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "+-------------+\n" +
			"| TOPICS NAME |\n" +
			"+-------------+\n" +
			"+-------------+",
	}

	noNamespaceName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}

	tenantNotExistError := pulsar.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := pulsar.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"topics",
		"Get the list of topics for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"topics",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doListTopics(vc)
	})
}

func doListTopics(vc *cmdutils.VerbCmd) error {
	tenantAndNamespace := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	listTopics, err := admin.Namespaces().GetTopics(tenantAndNamespace)
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Topics Name"})
		for _, topic := range listTopics {
			table.Append([]string{topic})
		}
		table.Render()
	}
	return err
}
