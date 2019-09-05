package topic

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/partitioned"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"topics",
		"Operations about topic(s)",
		"",
		"topic")

	commands := []func(*cmdutils.VerbCmd) {
		partitioned.CreateTopicCmd,
		partitioned.DeleteTopicCmd,
		partitioned.GetTopicCmd,
		partitioned.ListTopicsCmd,
		partitioned.UpdateTopicCmd,
	}

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd, commands...)

	return resourceCmd
}

func Hello()  {
	fmt.Print("hello")
}