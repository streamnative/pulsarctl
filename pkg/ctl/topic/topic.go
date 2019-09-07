package topic

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
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
	}

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd, commands...)

	return resourceCmd
}
