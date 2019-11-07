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

func createTenantCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for creating a new tenant."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc:    "create a tenant named (tenant-name)",
		Command: "pulsarctl tenants create (tenant-name)",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Create tenant (tenant-name) successfully",
	}
	out = append(out, successOut)
	out = append(out, tenantNameArgsError, tenantAlreadyExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Create a tenant",
		desc.ToString(),
		desc.ExampleToString(),
		"create")

	var tenantData utils.TenantData

	vc.SetRunFuncWithNameArg(func() error {
		return doCreateTenant(vc, &tenantData)
	}, "the tenant name is not specified or the tenant name is specified more than one")

	vc.FlagSetGroup.InFlagSet("TenantData", func(set *pflag.FlagSet) {
		set.StringSliceVarP(
			&tenantData.AdminRoles,
			"admin-roles",
			"r",
			[]string{""},
			"Allowed admins to access the tenant")
		set.StringSliceVarP(
			&tenantData.AllowedClusters,
			"allowed-clusters",
			"c",
			[]string{""},
			"Allowed clusters")
	})
}

func doCreateTenant(vc *cmdutils.VerbCmd, data *utils.TenantData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	data.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Tenants().Create(*data)
	if err == nil {
		vc.Command.Printf("Create tenant %s successfully\n", data.Name)
	}

	return err
}
