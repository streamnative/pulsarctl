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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func UpdateClusterCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for updating the cluster data of the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	updateURL := cmdutils.Example{
		Desc:    "updating the web service url of the (cluster-name)",
		Command: "pulsarctl clusters update --url http://example:8080 (cluster-name)",
	}
	examples = append(examples, updateURL)

	updateURLTLS := cmdutils.Example{
		Desc:    "updating the tls secured web service url of the (cluster-name)",
		Command: "pulsarctl clusters update --url-tls https://example:8080 (cluster-name)",
	}
	examples = append(examples, updateURLTLS)

	updateBrokerURL := cmdutils.Example{
		Desc:    "updating the broker service url of the (cluster-name)",
		Command: "pulsarctl clusters update --broker-url pulsar://example:6650 (cluster-name)",
	}
	examples = append(examples, updateBrokerURL)

	updateBrokerURLTLS := cmdutils.Example{
		Desc:    "updating the tls secured web service url of the (cluster-name)",
		Command: "pulsarctl clusters update --broker-url-tls pulsar+ssl://example:6650 (cluster-name)",
	}
	examples = append(examples, updateBrokerURLTLS)

	updatePeerCluster := cmdutils.Example{
		Desc:    "registered as a peer-cluster of the (cluster-name) clusters",
		Command: "pulsarctl clusters update -p (cluster-a) -p (cluster-b) (cluster)",
	}
	examples = append(examples, updatePeerCluster)

	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Cluster (cluster-name) updated",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"update",
		"Update a pulsar cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"update")

	clusterData := &utils.ClusterData{}

	vc.SetRunFuncWithNameArg(func() error {
		return doUpdateCluster(vc, clusterData)
	}, "the cluster name is not specified or the cluster name is specified more than one")

	// register the params
	vc.FlagSetGroup.InFlagSet("ClusterData", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&clusterData.ServiceURL,
			"url",
			"",
			"Pulsar cluster web service url, e.g. http://example.pulsar.io:8080")
		flagSet.StringVar(
			&clusterData.ServiceURLTls,
			"url-tls",
			"",
			"Pulsar cluster tls secured web service url, e.g. https://example.pulsar.io:8443")
		flagSet.StringVar(
			&clusterData.BrokerServiceURL,
			"broker-url",
			"",
			"Pulsar cluster broker service url, e.g. pulsar://example.pulsar.io:6650")
		flagSet.StringVar(
			&clusterData.BrokerServiceURLTls,
			"broker-url-tls",
			"",
			"Pulsar cluster tls secured broker service url, e.g. pulsar+ssl://example.pulsar.io:6651")
		flagSet.StringSliceVarP(
			&clusterData.PeerClusterNames,
			"peer-cluster",
			"p",
			[]string{""},
			"Cluster to be registered as a peer-cluster of this cluster.")
	})

}

func doUpdateCluster(vc *cmdutils.VerbCmd, clusterData *utils.ClusterData) error {
	clusterData.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().Update(*clusterData)
	if err == nil {
		vc.Command.Printf("Cluster %s updated\n", clusterData.Name)
	}
	return err
}
