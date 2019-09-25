package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var argsError = pulsar.Output{
	Desc: "the cluster name is not specified or the cluster name is specified more than one",
	Out:  "[✖]  only one argument is allowed to be used as a name",
}

var clusterNonExist = pulsar.Output{
	Desc: "the specified cluster does not exist in the broker",
	Out:  "[✖]  code: 412 reason: Cluster <cluster-name> does not exist.",
}

var failureDomainArgsError = pulsar.Output{
	Desc: "the cluster name and(or) failure domain name is not specified or the name is specified more than one",
	Out:  "[✖]  need to specified the cluster name and the failure domain name",
}

var checkFailureDomainArgs = func(args []string) error {
	if len(args) != 2 {
		return errors.New("need to specified the cluster name and the failure domain name")
	}
	return nil
}

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, CreateClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, listClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getClusterDataCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, UpdateClusterCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updatePeerClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getPeerClustersCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, createFailureDomainCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getFailureDomainCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, listFailureDomainCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteFailureDomainCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updateFailureDomainCmd)

	return resourceCmd
}
