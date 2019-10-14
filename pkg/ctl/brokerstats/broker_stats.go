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

package brokerstats

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"broker-stats",
		"Operations to collect broker statistics",
		"",
		"broker-stats")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, dumpMonitoringMetrics)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, dumpMBeans)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, dumpTopics)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, dumpAllocatorStats)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, dumpLoadReport)

	return resourceCmd
}
