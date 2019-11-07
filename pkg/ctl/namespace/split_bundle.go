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

func splitBundle(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Split a namespace-bundle from the current serving broker"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	splitBundle := cmdutils.Example{
		Desc:    "Split a namespace-bundle from the current serving broker",
		Command: "pulsarctl namespaces split-bundle tenant/namespace --bundle ({start-boundary}_{end-boundary})",
	}

	splitBundleWithUnload := cmdutils.Example{
		Desc: "Split a namespace-bundle from the current serving broker",
		Command: "pulsarctl namespaces split-bundle tenant/namespace \n" +
			"\t--bundle ({start-boundary}_{end-boundary})\n" +
			"\t--unload",
	}

	examples = append(examples, splitBundle, splitBundleWithUnload)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Split a namespace bundle: ({start-boundary}_{end-boundary}) successfully",
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

	ownershipFail := cmdutils.Output{
		Desc: "Please check if there is an active topic under the current split bundle.",
		Out:  "[✖]  code: 412 reason: Failed to find ownership for ServiceUnit:public/default/(bundle range)",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError, ownershipFail)
	desc.CommandOutput = out

	vc.SetDescription(
		"split-bundle",
		"Split a namespace-bundle from the current serving broker",
		desc.ToString(),
		desc.ExampleToString(),
		"split-bundle",
	)

	var data utils.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSplitBundle(vc, data)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&data.Bundle,
			"bundle",
			"b",
			"",
			"{start-boundary}_{end-boundary}")

		flagSet.BoolVarP(
			&data.Unload,
			"unload",
			"u",
			false,
			"Unload newly split bundles after splitting old bundle")

		cobra.MarkFlagRequired(flagSet, "bundle")
	})
}

func doSplitBundle(vc *cmdutils.VerbCmd, data utils.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Namespaces().SplitNamespaceBundle(ns, data.Bundle, data.Unload)
	if err == nil {
		vc.Command.Printf("Split a namespace bundle: %s successfully\n", data.Bundle)
	}
	return err
}
