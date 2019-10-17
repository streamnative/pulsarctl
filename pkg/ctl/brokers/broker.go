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

package brokers

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"brokers",
		"Operations about broker(s)",
		"",
		"broker")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getBrokerListCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getDynamicConfigListNameCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getOwnedNamespacesCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updateDynamicConfig)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteDynamicConfigCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getAllDynamicConfigsCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getInternalConfigCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getRuntimeConfigCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, healthCheckCmd)

	return resourceCmd
}
