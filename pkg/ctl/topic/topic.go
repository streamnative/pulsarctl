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
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/info"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/lookup"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/permission"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/stats"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/unload"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"topics",
		"Operations about topic(s)",
		"",
		"topic")

	commands := []func(*cmdutils.VerbCmd){
		CreateTopicCmd,
		DeleteTopicCmd,
		GetTopicCmd,
		ListTopicsCmd,
		UpdateTopicCmd,
		GetInternalInfoCmd,
		GrantPermissionCmd,
		RevokePermissions,
		GetPermissionsCmd,
		LookupTopicCmd,
		GetBundleRangeCmd,
		GetLastMessageIdCmd,
		GetStatsCmd,
		GetInternalStatsCmd,
		unload.UnloadCmd,
	}

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd, commands...)

	return resourceCmd
}
