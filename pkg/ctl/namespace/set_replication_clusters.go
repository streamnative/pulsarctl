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

package namespace

import (
	"strings"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func setReplicationClusters(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Set the replicated clusters for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	setClusters := pulsar.Example{
		Desc:    "Set the replicated clusters for a namespace",
		Command: "pulsarctl namespaces set-clusters tenant/namespace --clusters (cluster name)",
	}

	examples = append(examples, setClusters)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Set replication clusters successfully for tenant/namespace",
	}

	noNamespaceName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}

	tenantNotExistError := pulsar.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := pulsar.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	invalidClustersName := pulsar.Output{
		Desc: "Invalid cluster name, please check if your cluster name has the appropriate " +
			"permissions under the current tenant",
		Out: "[✖]  code: 403 reason: Cluster name is not in the list of allowed clusters list for tenant [public]",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError, invalidClustersName)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-clusters",
		"Set the replicated clusters for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"set-clusters",
	)

	var data pulsar.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSetReplicationClusters(vc, data)
	}, "the cluster name is not specified or the cluster name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&data.ClusterIds,
			"clusters",
			"c",
			"",
			"Replication Cluster Ids list (comma separated values)")

		cobra.MarkFlagRequired(flagSet, "clusters")
	})
}

func doSetReplicationClusters(vc *cmdutils.VerbCmd, data pulsar.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()

	clusters := strings.Split(data.ClusterIds, ",")
	err := admin.Namespaces().SetNamespaceReplicationClusters(ns, clusters)
	if err == nil {
		vc.Command.Printf("Set replication clusters successfully for %s\n", ns)
	}
	return err
}
