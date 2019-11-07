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

func getNsIsolationPolicies(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "List all namespace isolation policies of a cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc:    "List all namespace isolation policies of a cluster.",
		Command: "pulsarctl ns-isolation-policy list (cluster-name)",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"test-policiy-1\" : {\n" +
			"    \"namespaces\" : [ \"default\" ],\n" +
			"    \"primary\" : [ \"aaa\" ],\n" +
			"    \"secondary\" : [ ],\n" +
			"    \"auto_failover_policy\" : {\n" +
			"      \"policy_type\" : \"min_available\",\n" +
			"      \"parameters\" : {\n" +
			"        \"min_limit\" : \"3\",\n" +
			"        \"usage_threshold\" : \"100\"\n" +
			"      }\n" +
			"    }\n" +
			"  }\n" +
			"}",
	}

	clusterNameErr := cmdutils.Output{
		Desc: "Reason: Cluster name does not exist, please check cluster name.",
		Out:  "Reason: Cluster name does not exist.",
	}

	paramsErr := cmdutils.Output{
		Desc: "the cluster name is not specified or the cluster name is specified more than one, " +
			"please check cluster name",
		Out: "the cluster name is not specified or the cluster name is specified more than one",
	}
	out = append(out, successOut, clusterNameErr, paramsErr)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all namespace isolation policies of a cluster.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetNsIsolationPolicies(vc)
	}, "the cluster name is not specified or the cluster name is specified more than one")
}

func doGetNsIsolationPolicies(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	nsIsolationData, err := admin.NsIsolationPolicy().GetNamespaceIsolationPolicies(clusterName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), nsIsolationData)
	}
	return err
}
