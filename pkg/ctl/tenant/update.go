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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func UpdateTenantCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for updating the configuration of a tenant."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	empty := cmdutils.Example{
		Desc:    "clear the tenant configuration of a tenant",
		Command: "pulsarctl tenant update (tenant-name)",
	}
	examples = append(examples, empty)

	updateAdminRole := cmdutils.Example{
		Desc:    "update the admin roles for tenant (tenant-name)",
		Command: "pulsarctl tenants update --admin-roles (admin-A)--admin-roles (admin-B) (tenant-name)",
	}
	examples = append(examples, updateAdminRole)

	updateClusters := cmdutils.Example{
		Desc:    "update the allowed cluster list for tenant (tenant-name)",
		Command: "pulsarctl tenants update --allowed-clusters (cluster-A) --allowed-clusters (cluster-B) (tenant-name)",
	}
	examples = append(examples, updateClusters)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Update tenant [%s] successfully",
	}
	out = append(out, successOut)

	notExist := cmdutils.Output{
		Desc: "the specified tenant does not exist in",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}
	out = append(out, tenantNameArgsError, notExist)

	flagErrorOut := cmdutils.Output{
		Desc: "the flag --admin-roles or --allowed-clusters are not specified",
		Out:  "[✖]  the admin roles or the allowed clusters is not specified",
	}
	out = append(out, flagErrorOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"update",
		"update the configuration for a tenant",
		desc.ToString(),
		desc.ExampleToString(),
		"u")

	var data utils.TenantData

	vc.SetRunFuncWithNameArg(func() error {
		return doUpdateTenant(vc, &data)
	}, "the tenant name is not specified or the tenant name is specified more than one")

	vc.FlagSetGroup.InFlagSet("TenantData", func(set *pflag.FlagSet) {
		set.StringSliceVarP(
			&data.AdminRoles,
			"admin-roles",
			"r",
			nil,
			"Allowed admins to access the tenant")
		set.StringSliceVarP(
			&data.AllowedClusters,
			"allowed-clusters",
			"c",
			nil,
			"Allowed clusters")
	})
}

func doUpdateTenant(vc *cmdutils.VerbCmd, data *utils.TenantData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	data.Name = vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Tenants().Update(*data)
	if err == nil {
		vc.Command.Printf("Update tenant %s successfully\n", data.Name)
	}
	return err
}
