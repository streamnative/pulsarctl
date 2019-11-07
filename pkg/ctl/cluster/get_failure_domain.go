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

func getFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the specified failure domain on the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "getting the broker list in the (cluster-name) cluster failure domain (domain-name)",
		Command: "pulsarctl clusters get-failure-domain (cluster-name) (domain-name)",
	}
	examples = append(examples, get)

	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "output example",
		Out: "{\n  " +
			"\"brokers\" : [\n" +
			"    \"failure-broker-A\",\n" +
			"    \"failure-broker-B\",\n" +
			"  ]\n" +
			"}",
	}
	out = append(out, successOut, failureDomainArgsError, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-failure-domain",
		"Get the failure domain",
		desc.ToString(),
		desc.ExampleToString(),
		"gfd")

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doGetFailureDomain(vc)
	}, checkFailureDomainArgs)
}

func doGetFailureDomain(vc *cmdutils.VerbCmd) error {
	// fot testing
	if vc.NameError != nil {
		return vc.NameError
	}

	clusterName := vc.NameArgs[0]
	domainName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	resFailureDomain, err := admin.Clusters().GetFailureDomain(clusterName, domainName)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), resFailureDomain)
	}

	return err
}
