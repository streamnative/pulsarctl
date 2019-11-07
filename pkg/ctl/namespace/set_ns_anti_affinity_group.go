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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func setAntiAffinityGroup(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set the anti-affinity group for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	setAntiAffinityName := cmdutils.Example{
		Desc: "Set the anti-affinity group for a namespace",
		Command: "pulsarctl namespaces set-anti-affinity-group tenant/namespace \n" +
			"\t--group (anti-affinity group name)",
	}

	examples = append(examples, setAntiAffinityName)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set the anti-affinity group: (anti-affinity group name) successfully for <tenant/namespace>",
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
		"set-anti-affinity-group",
		"Set the anti-affinity group for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"set-anti-affinity-group",
	)

	var data utils.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSetAntiAffinityGroup(vc, data)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&data.AntiAffinityGroup,
			"group",
			"g",
			"",
			"Anti-affinity group name")

		cobra.MarkFlagRequired(flagSet, "group")
	})
}

func doSetAntiAffinityGroup(vc *cmdutils.VerbCmd, data utils.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Namespaces().SetNamespaceAntiAffinityGroup(ns, data.AntiAffinityGroup)
	if err == nil {
		vc.Command.Printf("Set the anti-affinity group: %s successfully for %s\n", data.AntiAffinityGroup, ns)
	}
	return err
}
