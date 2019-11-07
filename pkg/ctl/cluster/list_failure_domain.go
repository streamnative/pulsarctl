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

package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func listFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting all failure domain under the cluster (cluster-name)."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "listing all the failure domains under the specified cluster",
		Command: "pulsarctl clusters list-failure-domains (cluster-name)",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "output example",
		Out: "{\n" +
			"  \"failure-domain\": {\n" +
			"    \"brokers\": [\n" +
			"      \"failure-broker-A\",\n" +
			"      \"failure-broker-B\"\n" +
			"    ]\n" +
			"  }\n" +
			"}",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"list-failure-domains",
		"List the existing failure domains for a cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"lfd")

	vc.SetRunFuncWithNameArg(func() error {
		return doListFailureDomain(vc)
	}, "the cluster name is not specified or the cluster name is specified more than one")
}

func doListFailureDomain(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	clusterName := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	domainData, err := admin.Clusters().ListFailureDomains(clusterName)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), domainData)
	}
	return err
}
