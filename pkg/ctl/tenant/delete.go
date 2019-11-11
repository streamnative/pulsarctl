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

package tenant

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func deleteTenantCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription

	desc.CommandUsedFor = "This command is used for deleting an existing tenant."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	deleteExample := cmdutils.Example{
		Desc:    "delete a tenant named (tenant-name)",
		Command: "pulsarctl tenants delete (tenant-name)",
	}
	examples = append(examples, deleteExample)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Delete tenant <tenant-name> successfully",
	}
	out = append(out, successOut)

	NonEmptyError := cmdutils.Output{
		Desc: "there has namespace(s) under the tenant (tenant-name)",
		Out:  "code: 409 reason: The tenant still has active namespaces",
	}
	out = append(out, tenantNameArgsError, tenantNotExistError, NonEmptyError)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"d",
		desc.ToString(),
		desc.ExampleToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteTenant(vc)
	}, "the tenant name is not specified or the tenant name is specified more than one")
}

func doDeleteTenant(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	admin := cmdutils.NewPulsarClient()
	err := admin.Tenants().Delete(vc.NameArg)
	if err == nil {
		vc.Command.Printf("Delete tenant %s successfully\n", vc.NameArg)
	}
	return err
}
