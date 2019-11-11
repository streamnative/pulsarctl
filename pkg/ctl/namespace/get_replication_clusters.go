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
)

func getReplicationClusters(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the replicated clusters of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	getReplication := cmdutils.Example{
		Desc:    "Get the replicated clusters of a namespace",
		Command: "pulsarctl namespaces get-clusters tenant/namespace",
	}
	examples = append(examples, getReplication)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "(cluster name)",
	}
	noNamespaceName := cmdutils.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}

	tenantNotExistError := cmdutils.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := cmdutils.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-clusters",
		"Get the replicated clusters of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"get-clusters",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetReplicationClusters(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetReplicationClusters(vc *cmdutils.VerbCmd) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	clusters, err := admin.Namespaces().GetNamespaceReplicationClusters(ns)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), &clusters)
	}
	return err
}
