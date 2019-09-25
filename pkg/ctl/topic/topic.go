package topic

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/info"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/lookup"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/permission"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/stats"

	"github.com/spf13/cobra"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"topics",
		"Operations about topic(s)",
		"",
		"topic")

	commands := []func(*cmdutils.VerbCmd){
		crud.CreateTopicCmd,
		crud.DeleteTopicCmd,
		crud.GetTopicCmd,
		crud.ListTopicsCmd,
		crud.UpdateTopicCmd,
		permission.GrantPermissionCmd,
		permission.RevokePermissions,
		permission.GetPermissionsCmd,
		lookup.TopicCmd,
		lookup.GetBundleRangeCmd,
		info.GetLastMessageIDCmd,
		stats.GetStatsCmd,
		stats.GetInternalStatsCmd,
	}

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd, commands...)

	return resourceCmd
}
