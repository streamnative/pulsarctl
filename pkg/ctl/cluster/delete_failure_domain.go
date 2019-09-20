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
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func deleteFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for deleting the failure domain <domain-name> of the cluster <cluster-name>"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	delete := pulsar.Example{
		Desc:    "delete the failure domain",
		Command: "pulsarctl clusters delete-failure-domain <cluster-name> <domain-name>",
	}
	examples = append(examples, delete)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "output example",
		Out:  "Delete failure domain [<domain-name>] for cluster [<cluster-name>] succeed",
	}
	out = append(out, successOut, failureDomainArgsError)

	failureDomainNonExist := pulsar.Output{
		Desc: "the specified failure domain is not exist",
		Out: "code: 404 reason: Domain-name non-existent-failure-domain" +
			" or cluster standalone does not exist",
	}
	out = append(out, failureDomainNonExist)

	clusterNotExist := pulsar.Output{
		Desc: "the specified cluster is not exist",
		Out:  "code: 412 reason: Cluster non-existent-cluster does not exist.",
	}
	out = append(out, clusterNotExist)

	desc.CommandOutput = out

	vc.SetDescription(
		"delete-failure-domain",
		"Delete a failure domain",
		desc.ToString(),
		"dfd")

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doDeleteFailureDomain(vc)
	}, checkFailureDomainArgs)
}

func doDeleteFailureDomain(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	var failureDomain pulsar.FailureDomainData
	failureDomain.ClusterName = vc.NameArgs[0]
	failureDomain.DomainName = vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().DeleteFailureDomain(failureDomain)
	if err == nil {
		vc.Command.Printf("Delete failure domain [%s] for cluster [%s] succeed\n", failureDomain.DomainName, failureDomain.ClusterName)
	}

	return err
}
