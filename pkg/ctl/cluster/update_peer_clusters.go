package cluster

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func updatePeerClustersCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for updating peer clusters."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	update := pulsar.Example{
		Desc:    "updating the <cluster-name> peer clusters",
		Command: "pulsarctl clusters update-peer-clusters -p cluster-a -p cluster-b <cluster-name>",
	}
	examples = append(examples, update)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "<cluster-name> peer clusters updated",
	}
	out = append(out, successOut)

	failOut := pulsar.Output{
		Desc: "output of doesn't specified a cluster name",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}
	out = append(out, failOut)

	clusterNotExist := pulsar.Output{
		Desc: "output of cluster doesn't exist",
		Out:  "[✖]  code: 404 reason: Cluster does not exist",
	}
	out = append(out, clusterNotExist)

	desc.CommandOutput = out

	vc.SetDescription(
		"update-peer-clusters",
		"Update the peer clusters",
		desc.ToString(),
		"upc")

	clusterData := &pulsar.ClusterData{}

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

func doUpdatePeerClusters(vc *cmdutils.VerbCmd, clusterData *pulsar.ClusterData) error {
	clusterData.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().UpdatePeerClusters(clusterData.Name, clusterData.PeerClusterNames)
	if err == nil {
		vc.Command.Printf("%s peer clusters updated", clusterData.Name)
	}
	return err
}
