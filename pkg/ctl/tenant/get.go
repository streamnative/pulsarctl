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
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getTenantCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting the configuration of a tenant."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	getSuccess := pulsar.Example{
		Desc:    "get the configuration of tenant (tenant-name)",
		Command: "pulsarctl tenants get (tenant-name)",
	}
	examples = append(examples, getSuccess)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"adminRoles\": [],\n" +
			"  \"allowedClusters\": [\n" +
			"    \"standalone\"\n" +
			"  ]\n" +
			"}",
	}
	out = append(out, successOut)
	notExist := pulsar.Output{
		Desc: "the specified tenant does not exist in the cluster",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}
	out = append(out, tenantNameArgsError, notExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"get the configuration of a tenant",
		desc.ToString(),
		desc.ExampleToString(),
		"g")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetTenant(vc)
	})
}

func doGetTenant(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	admin := cmdutils.NewPulsarClient()
	data, err := admin.Tenants().Get(vc.NameArg)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), data)
	}
	return err
}
