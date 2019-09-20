package topic

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/info"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/permission"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/info"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/lookup"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/stats"
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
	}

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd, commands...)

	return resourceCmd
}
