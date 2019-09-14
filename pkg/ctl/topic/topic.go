package topic

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/compact"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/offload"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/stats"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/teminate"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/unload"
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

		GetStatsCmd,
		GetInternalStatsCmd,

		UnloadCmd,
		OffloadCmd,
		OffloadStatusCmd,
		CompactCmd,
		CompactStatusCmd,
		TerminateCmd,
	}

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd, commands...)

	return resourceCmd
}
