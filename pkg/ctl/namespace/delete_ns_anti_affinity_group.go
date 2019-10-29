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
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
)

func deleteAntiAffinityGroup(vc *cmdutils.VerbCmd) {
	desc := common.LongDescription{}
	desc.CommandUsedFor = "Delete an anti-affinity group of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []common.Example
	delAntiAffinity := common.Example{
		Desc:    "Delete an anti-affinity group of a namespace",
		Command: "pulsarctl namespaces delete-anti-affinity-group tenant/namespace",
	}
	examples = append(examples, delAntiAffinity)
	desc.CommandExamples = examples

	var out []common.Output
	successOut := common.Output{
		Desc: "normal output",
		Out:  "Delete the anti-affinity group successfully for [tenant/namespace]",
	}

	noNamespaceName := common.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}

	tenantNotExistError := common.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := common.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete-anti-affinity-group",
		"Delete an anti-affinity group of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"delete-anti-affinity-group",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteAntiAffinityGroup(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doDeleteAntiAffinityGroup(vc *cmdutils.VerbCmd) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Namespaces().DeleteNamespaceAntiAffinityGroup(ns)
	if err == nil {
		vc.Command.Printf("Delete the anti-affinity group successfully for [%s]", ns)
	}
	return err
}
