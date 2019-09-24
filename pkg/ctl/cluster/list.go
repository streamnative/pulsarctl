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
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func listClustersCmd(vc *cmdutils.VerbCmd) {
	// update the description
	vc.SetDescription(
		"list",
		"List the available pulsar clusters",
		"This command is used for listing the list of available pulsar clusters.")

	// set the run function
	vc.SetRunFunc(func() error {
		return doListClusters(vc)
	})
}

func doListClusters(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	clusters, err := admin.Clusters().List()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{ "Cluster Name" })

		for _, c := range clusters {
			table.Append([]string { c })
		}

		table.Render()
	}
	return err
}