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
)

func deleteNs(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Deletes a namespace. The namespace needs to be empty"
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example
	del := pulsar.Example{
		Desc:    "Deletes a namespace",
		Command: "pulsarctl namespaces delete <namespace-name>",
	}
	examples = append(examples, del)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Created <namespace-name> successfully",
	}

	notTenantName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}

	notExistTenantName := pulsar.Output{
		Desc: "the tenant name not exist, please check the tenant name",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	out = append(out, successOut, notExistTenantName, notTenantName)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Deletes a namespace. The namespace needs to be empty",
		desc.ToString(),
		"delete",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteNs(vc)
	})
}

func doDeleteNs(vc *cmdutils.VerbCmd) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Namespaces().DeleteNamespace(ns)
	if err == nil {
		vc.Command.Printf("Deleted %s successfully", ns)
	}
	return err
}
