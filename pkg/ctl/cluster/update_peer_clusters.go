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

	"github.com/spf13/pflag"
)

func updatePeerClustersCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for updating peer clusters."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	update := pulsar.Example{
		Desc:    "updating the <cluster-name> peer clusters",
		Command: "pulsarctl clusters update-peer-clusters -p cluster-a -p cluster-b (cluster-name)",
	}
	examples = append(examples, update)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "output example",
		Out:  "(cluster-name) peer clusters updated",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"update-peer-clusters",
		"Update the peer clusters",
		desc.ToString(),
		desc.ExampleToString(),
		"upc")

	clusterData := &pulsar.ClusterData{}

	vc.SetRunFuncWithNameArg(func() error {
		return doUpdatePeerClusters(vc, clusterData)
	}, "the cluster name is not specified or the cluster name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Update peer clusters", func(set *pflag.FlagSet) {
		set.StringSliceVarP(
			&clusterData.PeerClusterNames,
			"peer-cluster",
			"p",
			[]string{""},
			"Cluster to be registered as a peer-cluster of this cluster")
	})

}

func doUpdatePeerClusters(vc *cmdutils.VerbCmd, clusterData *pulsar.ClusterData) error {
	clusterData.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().UpdatePeerClusters(clusterData.Name, clusterData.PeerClusterNames)
	if err == nil {
		vc.Command.Printf("%s peer clusters updated\n", clusterData.Name)
	}
	return err
}
