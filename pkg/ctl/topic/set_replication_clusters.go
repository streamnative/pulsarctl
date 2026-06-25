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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func SetReplicationClustersCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set the replication clusters for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	setReplication := cmdutils.Example{
		Desc:    "Set the replication clusters for a topic",
		Command: "pulsarctl topics set-replication-clusters tenant/namespace/topic --clusters cluster1,cluster2",
	}
	examples = append(examples, setReplication)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set the replication clusters for [topic] successfully",
	}

	noTopicName := cmdutils.Output{
		Desc: "you must specify a tenant/namespace/topic name, please check if the tenant/namespace/topic name is provided",
		Out:  "[✖]  the topic name is not specified or the topic name is specified more than one",
	}

	tenantNotExistError := cmdutils.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := cmdutils.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	out = append(out, successOut, noTopicName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-replication-clusters",
		"Set the replication clusters for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-replication-clusters",
	)

	var clusters []string

	vc.FlagSetGroup.InFlagSet("Set replication clusters", func(flagSet *pflag.FlagSet) {
		flagSet.StringSliceVarP(&clusters, "clusters", "c", nil,
			"Replication cluster ids.")
		_ = cobra.MarkFlagRequired(flagSet, "clusters")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doSetReplicationClusters(vc, clusters)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doSetReplicationClusters(vc *cmdutils.VerbCmd, clusters []string) error {
	topic := vc.NameArg

	if len(clusters) == 0 {
		return errors.New("clusters cannot be empty")
	}

	admin := cmdutils.NewPulsarClient()
	topicName, err := utils.GetTopicName(topic)
	if err != nil {
		return err
	}

	err = admin.Topics().SetReplicationClusters(*topicName, clusters)
	if err == nil {
		vc.Command.Printf("Set the replication clusters successfully on [%s]\n", topic)
	}

	return err
}
