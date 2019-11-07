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

	"github.com/olekukonko/tablewriter"
)

func getPeerClustersCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the peer clusters of the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "getting the (cluster-name) peer clusters",
		Command: "pulsarctl clusters get-peer-clusters (cluster-name)",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "+-------------------+\n" +
			"|   PEER CLUSTERS   |\n" +
			"+-------------------+\n" +
			"| test_peer_cluster |\n" +
			"+-------------------+",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)

	desc.CommandOutput = out

	vc.SetDescription(
		"get-peer-clusters",
		"Getting list of peer clusters",
		desc.ToString(),
		desc.ExampleToString(),
		"gpc")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetPeerClusters(vc)
	}, "the cluster name is not specified or the cluster name is specified more than one")
}

func doGetPeerClusters(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	peerClusters, err := admin.Clusters().GetPeerClusters(clusterName)
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Peer clusters"})

		for _, c := range peerClusters {
			table.Append([]string{c})
		}

		table.Render()
	}
	return err
}
