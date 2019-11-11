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

func CreateClusterCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for adding the configuration data for a cluster. " +
		"The configuration data is mainly used for geo-replication between clusters, so please " +
		"make sure the service urls provided in this command are reachable between clusters. " +
		"This operation requires Pulsar super-user privileges."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc:    "Provisions a new cluster",
		Command: "pulsarctl clusters create (cluster-name)",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Cluster (cluster-name) added",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"add",
		"Add a pulsar cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"create")

	clusterData := &utils.ClusterData{}

	// set the run function with name argument
	vc.SetRunFuncWithNameArg(func() error {
		return doCreateCluster(vc, clusterData)
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
		flagSet.StringArrayVarP(
			&clusterData.PeerClusterNames,
			"peer-cluster",
			"p",
			[]string{""},
			"Cluster to be registered as a peer-cluster of this cluster.")
	})
}

func doCreateCluster(vc *cmdutils.VerbCmd, clusterData *utils.ClusterData) error {
	clusterData.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().Create(*clusterData)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Cluster %s added\n", clusterData.Name)
	}
	return err
}
