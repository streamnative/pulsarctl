package subscription

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/messages"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"subscriptions",
		"Operations about subscription(s)",
		"",
		"subscription")

	command := []func(cmd *cmdutils.VerbCmd){
		CreateCmd,
		DeleteCmd,
		ListCmd,

		ExpireCmd,
		ResetCursorCmd,
		SkipCmd,
	}

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd, command...)

	return resourceCmd
}
