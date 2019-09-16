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

func getRetention(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Get the retention policy for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	getRetention := pulsar.Example{
		Desc:    "Get the retention policy for a namespace",
		Command: "pulsarctl namespaces get-retention tenant/namespace",
	}
	examples = append(examples, getRetention)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"RetentionTimeInMinutes\": 0,\n" +
			"  \"RetentionSizeInMB\": 0\n" +
			"}",
	}

	notTenantName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}

	notExistTenantName := pulsar.Output{
		Desc: "the tenant name not exist, please check the tenant name",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	notExistNsName := pulsar.Output{
		Desc: "the namespace not exist, please check namespace name",
		Out:  "[✖]  code: 404 reason: Namespace <tenant/namespace> does not exist",
	}

	out = append(out, successOut, notTenantName, notExistTenantName, notExistNsName)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-retention",
		"Get the retention policy for a namespace",
		desc.ToString(),
		"get-retention",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetRetention(vc)
	})
}

func doGetRetention(vc *cmdutils.VerbCmd) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	policy, err := admin.Namespaces().GetRetention(ns)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), &policy)
	}
	return err
}
