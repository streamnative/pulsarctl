package cluster

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func updatePeerClustersCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for updating peer clusters."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	update := Example{
		Desc:    "updating the <cluster-name> peer clusters",
		Command: "pulsarctl clusters update-peer-clusters -p cluster-a -p cluster-b <cluster-name>",
	}
	examples = append(examples, update)
	desc.CommandExamples = examples

	var out []Output
	successOut := Output{
		Desc: "output example",
		Out:  "<cluster-name> peer clusters updated",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"update-peer-clusters",
		"Update the peer clusters",
		desc.ToString(),
		"upc")

	clusterData := &ClusterData{}

	vc.SetRunFuncWithNameArg(func() error {
		return doUpdatePeerClusters(vc, clusterData)
	})

	vc.FlagSetGroup.InFlagSet("Update peer clusters", func(set *pflag.FlagSet) {
		set.StringSliceVarP(
			&clusterData.PeerClusterNames,
			"peer-cluster",
			"p",
			[]string{""},
			"Cluster to be registered as a peer-cluster of this cluster")
	})

}

func doUpdatePeerClusters(vc *cmdutils.VerbCmd, clusterData *ClusterData) error {
	clusterData.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().UpdatePeerClusters(clusterData.Name, clusterData.PeerClusterNames)
	if err == nil {
		vc.Command.Printf("%s peer clusters updated", clusterData.Name)
	}
	return err
}
