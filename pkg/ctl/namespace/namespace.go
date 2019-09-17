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
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"namespaces",
		"Operations about namespaces",
		"",
		"namespace",
	)

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getNamespacesFromTenant)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getTopics)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getPolicies)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, createNs)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteNs)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, setMessageTTL)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getMessageTTL)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getRetention)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, setRetention)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getBacklogQuota)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, setBacklogQuota)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, removeBacklogQuota)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getAntiAffinityGroup)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, setAntiAffinityGroup)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteAntiAffinityGroup)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getAntiAffinityNamespaces)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getPersistence)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, setPersistence)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, setDeduplication)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, setReplicationClusters)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getReplicationClusters)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, unload)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, splitBundle)

	return resourceCmd
}
