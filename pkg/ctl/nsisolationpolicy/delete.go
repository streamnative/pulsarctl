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

func deleteNsIsolationPolicy(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Delete namespace isolation policy of a cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc:    "Delete namespace isolation policy of a cluster.",
		Command: "pulsarctl ns-isolation-policy delete (cluster-name) (policy-name)",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Delete namespaces isolation policy:(policy name) successful.",
	}

	policyNameErr := cmdutils.Output{
		Desc: "NamespaceIsolationPolicies for cluster standalone does not exist, please check policy name.",
		Out:  "NamespaceIsolationPolicies for cluster standalone does not exist",
	}

	clusterNameErr := cmdutils.Output{
		Desc: "Reason: Cluster name does not exist, please check cluster name.",
		Out:  "Reason: Cluster name does not exist.",
	}

	paramsErr := cmdutils.Output{
		Desc: "need to specified the cluster name and the policy name, please add cluster name and policy name",
		Out:  "need to specified the cluster name and the policy name",
	}
	out = append(out, successOut, policyNameErr, clusterNameErr, paramsErr)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete namespace isolation policy of a cluster.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doDeleteNsIsolationPolicy(vc)
	}, checkNsIsolationPolicyArgs)
}

func doDeleteNsIsolationPolicy(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArgs[0]
	policyName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	err := admin.NsIsolationPolicy().DeleteNamespaceIsolationPolicy(clusterName, policyName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Delete namespaces isolation policy: %s successful\n", policyName)
	}
	return err
}
