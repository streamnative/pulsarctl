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
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getAntiAffinityNamespaces(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Get Anti-affinity namespaces grouped with the given anti-affinity group name"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	getRetention := pulsar.Example{
		Desc:    "Get Anti-affinity namespaces grouped with the given anti-affinity group name",
		Command: "pulsarctl namespaces get-anti-affinity-namespaces tenant/namespace",
	}
	examples = append(examples, getRetention)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "<anti-affinity name list>",
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

	var data pulsar.NamespacesData

	vc.SetDescription(
		"get-anti-affinity-namespaces",
		"Get Anti-affinity namespaces grouped with the given anti-affinity group name",
		desc.ToString(),
		"get-anti-affinity-namespaces",
	)

	vc.SetRunFunc(func() error {
		return doGetAntiAffinityNamespaces(vc, data)
	})

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&data.AntiAffinityGroup,
			"group",
			"g",
			"",
			"Anti-affinity group name")

		flagSet.StringVarP(
			&data.Cluster,
			"cluster",
			"c",
			"",
			"Cluster name")

		flagSet.StringVarP(
			&data.Tenant,
			"tenant",
			"p",
			"",
			"tenant is only used for authorization. \n"+
				"Client has to be admin of any of the tenant to access this api")

		cobra.MarkFlagRequired(flagSet, "group")
	})
}

func doGetAntiAffinityNamespaces(vc *cmdutils.VerbCmd, data pulsar.NamespacesData) error {
	admin := cmdutils.NewPulsarClientWithApiVersion(pulsar.V1)
	strList, err := admin.Namespaces().GetAntiAffinityNamespaces(data.Tenant, data.Cluster, data.AntiAffinityGroup)
	if err == nil {
		vc.Command.Println(strList)
	}
	return err
}
