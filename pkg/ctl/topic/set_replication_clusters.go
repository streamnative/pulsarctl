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

package topic

import (
	"strings"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func SetReplicationClustersCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set the replication clusters for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	setClusters := cmdutils.Example{
		Desc:    "Set the replication clusters for a topic",
		Command: "pulsarctl topics set-replication-clusters persistent://tenant/namespace/topic --clusters cluster1,cluster2",
	}

	examples = append(examples, setClusters)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set replication clusters successfully for [topic]",
	}

	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)

	invalidClustersName := cmdutils.Output{
		Desc: "Invalid cluster name, please check if your cluster name has the appropriate " +
			"permissions under the current tenant",
		Out: "[✖]  code: 403 reason: Invalid cluster id: <cluster-name>",
	}

	out = append(out, invalidClustersName, TopicLevelPolicyNotEnabledError)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-replication-clusters",
		"Set the replication clusters for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-replication-clusters",
	)

	var clusterIDs string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetReplicationClusters(vc, clusterIDs)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("ReplicationClusters", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&clusterIDs,
			"clusters",
			"c",
			"",
			"Replication Cluster Ids list (comma separated values)")

		_ = cobra.MarkFlagRequired(flagSet, "clusters")
	})
	vc.EnableOutputFlagSet()
}

func doSetReplicationClusters(vc *cmdutils.VerbCmd, clusterIDs string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	clusters := strings.Split(clusterIDs, ",")
	err = admin.Topics().SetReplicationClusters(*topic, clusters)
	if err == nil {
		vc.Command.Printf("Set replication clusters successfully for [%s]\n", topic.String())
	}
	return err
}
