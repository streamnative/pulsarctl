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
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func getAntiAffinityNamespaces(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the list of namespaces in the same anti-affinity group."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	getRetention := cmdutils.Example{
		Desc:    "Get the list of namespaces in the same anti-affinity group.",
		Command: "pulsarctl namespaces get-anti-affinity-namespaces tenant/namespace",
	}
	examples = append(examples, getRetention)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "(anti-affinity name list)",
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

	var data utils.NamespacesData

	vc.SetDescription(
		"get-anti-affinity-namespaces",
		"Get the list of namespaces in the same anti-affinity group.",
		desc.ToString(),
		desc.ExampleToString(),
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
			"t",
			"",
			"tenant is only used for authorization. \n"+
				"Client has to be admin of any of the tenant to access this api")

		cobra.MarkFlagRequired(flagSet, "group")
	})
}

func doGetAntiAffinityNamespaces(vc *cmdutils.VerbCmd, data utils.NamespacesData) error {
	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V1)
	strList, err := admin.Namespaces().GetAntiAffinityNamespaces(data.Tenant, data.Cluster, data.AntiAffinityGroup)
	if err == nil {
		vc.Command.Println(strList)
	}
	return err
}
