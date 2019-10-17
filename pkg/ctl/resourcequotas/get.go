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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func getResourceQuota(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "Get the resource quota for specified namespace bundle, " +
		"or default quota if no namespace/bundle specified."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Get the resource quota use default namespace/bundle",
		Command: "pulsarctl resource-quotas get",
	}
	getWithArgs := pulsar.Example{
		Desc:    "Get the resource quota for specified namespace bundle",
		Command: "pulsarctl resource-quotas get --namespace (namespace name) --bundle (bundle range)",
	}
	examples = append(examples, get, getWithArgs)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"msgRateIn\" : 40.0,\n" +
			"  \"msgRateOut\" : 120.0,\n" +
			"  \"bandwidthIn\" : 100000.0,\n" +
			"  \"bandwidthOut\" : 300000.0,\n" +
			"  \"memory\" : 80.0,\n" +
			"  \"dynamic\" : true\n" +
			"}",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"Get the resource quota for specified namespace bundle, "+
			"or default quota if no namespace/bundle specified.",
		desc.ToString(),
		desc.ExampleToString(),
		"get")

	quotaData := &pulsar.ResourceQuotaData{}

	vc.SetRunFunc(func() error {
		return doGetResourceQuota(vc, quotaData)
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
	})
}

func doGetResourceQuota(vc *cmdutils.VerbCmd, quotaData *pulsar.ResourceQuotaData) error {
	admin := cmdutils.NewPulsarClient()

	var err error

	switch {
	case quotaData.Bundle == "" && quotaData.Names == "":
		resourceQuotaData, err := admin.ResourceQuotas().GetDefaultResourceQuota()
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			cmdutils.PrintJSON(vc.Command.OutOrStdout(), resourceQuotaData)
		}
	case quotaData.Bundle != "" && quotaData.Names != "":
		nsName, err := pulsar.GetNamespaceName(quotaData.Names)
		if err != nil {
			return err
		}
		resourceQuotaData, err := admin.ResourceQuotas().GetNamespaceBundleResourceQuota(
			nsName.String(), quotaData.Bundle)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			cmdutils.PrintJSON(vc.Command.OutOrStdout(), resourceQuotaData)
		}
	default:
		return errors.New("Namespace and bundle must be provided together")
	}

	return err
}
