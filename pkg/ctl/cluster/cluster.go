package cluster

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

var argsError = pulsar.Output{
	Desc: "the cluster name is not specified or the cluster name is specified more than one",
	Out:  "[✖]  only one argument is allowed to be used as a name",
}

var clusterNonExist = pulsar.Output{
	Desc: "the specified cluster does not exist in the broker",
	Out:  "[✖]  code: 404 reason: Cluster does not exist",
}

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, createClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, listClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getClusterDataCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updateClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updatePeerClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getPeerClustersCmd)

	return resourceCmd
}
