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

package resourcequotas

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func resetNamespaceBundleResourceQuota(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "Reset the specified namespace bundle's resource quota to default value."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	reset := pulsar.Example{
		Desc:    "Reset the specified namespace bundle's resource quota to default value",
		Command: "pulsarctl resource-quotas get",
	}

	examples = append(examples, reset)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Reset resource quota successful",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"reset-namespace-bundle-quota",
		"Reset the specified namespace bundle's resource quota to default value.",
		desc.ToString(),
		desc.ExampleToString(),
		"delete")

	quotaData := &pulsar.ResourceQuotaData{}

	vc.SetRunFunc(func() error {
		return doResetNamespaceBundleResourceQuota(vc, quotaData)
	})

	vc.FlagSetGroup.InFlagSet("SchemaConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&quotaData.Names,
			"namespace",
			"n",
			"",
			"cluster/namespace, must be specified together with '--bundle'")
		flagSet.StringVarP(
			&quotaData.Bundle,
			"bundle",
			"b",
			"",
			"{start-boundary}_{end-boundary}, must be specified together with '--namespace'")

		cobra.MarkFlagRequired(flagSet, "namespace")
		cobra.MarkFlagRequired(flagSet, "bundle")
	})
}

func doResetNamespaceBundleResourceQuota(vc *cmdutils.VerbCmd, quotaData *pulsar.ResourceQuotaData) error {
	admin := cmdutils.NewPulsarClient()

	nsName, err := pulsar.GetNamespaceName(quotaData.Names)
	if err != nil {
		return err
	}
	err = admin.ResourceQuotas().ResetNamespaceBundleResourceQuota(nsName.String(), quotaData.Bundle)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Println("Reset resource quota successful")
	}

	return err
}
