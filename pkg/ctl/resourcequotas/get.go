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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/pkg/errors"
)

func getResourceQuota(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Get the resource quota for a specified namespace bundle, " +
		"or default quota if no namespace/bundle is specified."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the resource quota use default namespace/bundle",
		Command: "pulsarctl resource-quotas get",
	}
	getWithArgs := cmdutils.Example{
		Desc:    "Get the resource quota for a specified namespace bundle",
		Command: "pulsarctl resource-quotas get (namespace name) (bundle range)",
	}
	examples = append(examples, get, getWithArgs)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
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
		"Get the resource quota for a specified namespace bundle, "+
			"or default quota if no namespace/bundle is specified.",
		desc.ToString(),
		desc.ExampleToString(),
		"get")

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doGetResourceQuota(vc)
	}, func(args []string) error {
		if len(args) > 2 && len(args) == 1 {
			return errors.New("need two arguments or zero arguments apply to the command")
		}
		return nil
	})
}

func doGetResourceQuota(vc *cmdutils.VerbCmd) error {
	var namespace, bundle string
	if len(vc.NameArgs) > 0 {
		namespace = vc.NameArgs[0]
		bundle = vc.NameArgs[1]
	}
	admin := cmdutils.NewPulsarClient()

	var err error

	switch {
	case bundle == "" && namespace == "":
		resourceQuotaData, err := admin.ResourceQuotas().GetDefaultResourceQuota()
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			cmdutils.PrintJSON(vc.Command.OutOrStdout(), resourceQuotaData)
		}
	case bundle != "" && namespace != "":
		nsName, err := utils.GetNamespaceName(namespace)
		if err != nil {
			return err
		}
		resourceQuotaData, err := admin.ResourceQuotas().GetNamespaceBundleResourceQuota(
			nsName.String(), bundle)
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
