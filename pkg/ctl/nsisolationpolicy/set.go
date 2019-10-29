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
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func setPolicy(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "Create/Update a namespace isolation policy for a cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	create := pulsar.Example{
		Desc: "Create/Update a namespace isolation policy for a cluster",
		Command: "pulsarctl ns-isolation-policy set (cluster-name) (policy name) " +
			"--auto-failover-policy-params min_limit=3,usage_threshold=100 " +
			"--auto-failover-policy-type min_available " +
			"--namespaces default " +
			"--primary test-primary " +
			"--secondary test-secondary",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Create/Update namespaces isolation policy: (policy name) successful.",
	}

	clusterNameErr := pulsar.Output{
		Desc: "Reason: Cluster name does not exist, please check cluster name.",
		Out:  "Reason: Cluster name does not exist.",
	}

	policyNameErr := pulsar.Output{
		Desc: "NamespaceIsolationPolicies for cluster standalone does not exist, please check policy name.",
		Out:  "NamespaceIsolationPolicies for cluster standalone does not exist",
	}

	paramsErr := pulsar.Output{
		Desc: "the cluster name is not specified or the cluster name is specified more than one, " +
			"please check cluster name",
		Out: "the cluster name is not specified or the cluster name is specified more than one",
	}
	out = append(out, successOut, clusterNameErr, policyNameErr, paramsErr)
	desc.CommandOutput = out

	vc.SetDescription(
		"set",
		"Create/Update a namespace isolation policy for a cluster",
		desc.ToString(),
		desc.ExampleToString())

	nsData := &pulsar.NsIsolationPoliciesData{}

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doSetPolicy(vc, nsData)
	}, checkNsIsolationPolicyArgs)

	vc.FlagSetGroup.InFlagSet("NsIsolationPoliciesData", func(flagSet *pflag.FlagSet) {
		flagSet.StringSliceVar(
			&nsData.Namespaces,
			"namespaces",
			nil,
			"Broker-name to get namespace-isolation policies attached to it")
		flagSet.StringSliceVar(
			&nsData.Primary,
			"primary",
			nil,
			"Broker-name to get namespace-isolation policies attached to it")
		flagSet.StringSliceVar(
			&nsData.Secondary,
			"secondary",
			nil,
			"Broker-name to get namespace-isolation policies attached to it")
		flagSet.StringVar(
			&nsData.AutoFailoverPolicyTypeName,
			"auto-failover-policy-type",
			"",
			"Broker-name to get namespace-isolation policies attached to it")
		flagSet.StringVar(
			&nsData.AutoFailoverPolicyParams,
			"auto-failover-policy-params",
			"",
			"comma separated name=value auto failover policy parameters")

		cobra.MarkFlagRequired(flagSet, "namespaces")
		cobra.MarkFlagRequired(flagSet, "primary")
		cobra.MarkFlagRequired(flagSet, "auto-failover-policy-type")
		cobra.MarkFlagRequired(flagSet, "auto-failover-policy-params")
	},
	)
}

func doSetPolicy(vc *cmdutils.VerbCmd, nsData *pulsar.NsIsolationPoliciesData) error {
	clusterName := vc.NameArgs[0]
	policyName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()

	policyParams, err := utils.Convert(nsData.AutoFailoverPolicyParams)
	if err != nil {
		return err
	}

	namespaceIsolationData, err := pulsar.CreateNamespaceIsolationData(nsData.Namespaces, nsData.Primary,
		nsData.Secondary, nsData.AutoFailoverPolicyTypeName, policyParams)
	if err != nil {
		return err
	}
	err = admin.NsIsolationPolicy().CreateNamespaceIsolationPolicy(clusterName, policyName, *namespaceIsolationData)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Create/Update namespaces isolation policy:%s successful\n", policyName)
	}
	return err
}
