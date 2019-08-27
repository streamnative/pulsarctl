package cluster

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, createClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, listClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getClusterDataCmd)
<<<<<<< HEAD
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteClusterCmd)
=======
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updateClusterCmd)
>>>>>>> Add command cluster update

	return resourceCmd
}
