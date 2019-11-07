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

func deleteClusterCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for deleting an existing cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	delete := cmdutils.Example{
		Desc:    "deleting the cluster named (cluster-name)",
		Command: "pulsarctl clusters delete (cluster-name)",
	}
	examples = append(examples, delete)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Cluster (cluster-name) delete successfully.",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete an existing cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"delete")

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteCluster(vc)
	}, "the cluster name is not specified or the cluster name is specified more than one")
}

func doDeleteCluster(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().Delete(clusterName)
	if err == nil {
		vc.Command.Printf("Cluster %s delete successfully\n", clusterName)
	}
	return err
}
