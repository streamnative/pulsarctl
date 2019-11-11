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

package nsisolationpolicy

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getAllBrokersWithPolicies(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "List all brokers with namespace-isolation policies attached to it."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc:    "List all brokers with namespace-isolation policies attached to it",
		Command: "pulsarctl ns-isolation-policy brokers (cluster-name)",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "[\n" +
			"  {\n" +
			"    \"brokerName\": \"127.0.0.1:8080\",\n" +
			"    \"policyName\": \"\",\n" +
			"    \"isPrimary\": false,\n" +
			"    \"namespaceRegex\": null\n" +
			"  }\n" +
			"]",
	}

	clusterNameErr := cmdutils.Output{
		Desc: "Reason: Cluster name does not exist, please check cluster name.",
		Out:  "Reason: Cluster name does not exist.",
	}

	paramsErr := cmdutils.Output{
		Desc: "need to specified the cluster name and the policy name, please add cluster name and policy name",
		Out:  "need to specified the cluster name and the policy name",
	}

	noPolicies := cmdutils.Output{
		Desc: "namespace-isolation policies not found for standalone",
		Out:  "[✖]  code: 404 reason: namespace-isolation policies not found for standalone",
	}

	out = append(out, successOut, clusterNameErr, paramsErr, noPolicies)
	desc.CommandOutput = out

	vc.SetDescription(
		"brokers",
		"List all brokers with namespace-isolation policies attached to it",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetAllBrokersWithPolicies(vc)
	}, "the cluster name is not specified or the cluster name is specified more than one")
}

func doGetAllBrokersWithPolicies(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	nsIsolationData, err := admin.NsIsolationPolicy().GetBrokersWithNamespaceIsolationPolicy(clusterName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), nsIsolationData)
	}
	return err
}
