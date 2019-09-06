package cluster

import (
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getPeerClustersCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the peer clusters of the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	get := Example{
		Desc:    "getting the <cluster-name> peer clusters",
		Command: "pulsarctl clusters get-peer-clusters <cluster-name>",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: "+-------------------+\n" +
			"|   PEER CLUSTERS   |\n" +
			"+-------------------+\n" +
			"| test_peer_cluster |\n" +
			"+-------------------+",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)

	desc.CommandOutput = out

	vc.SetDescription(
		"get-peer-clusters",
		"Getting list of peer clusters",
		desc.ToString(),
		"gpc")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetPeerClusters(vc)
	})
}

func doGetPeerClusters(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	peerClusters, err := admin.Clusters().GetPeerClusters(clusterName)
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Peer clusters"})

		for _, c := range peerClusters {
			table.Append([]string{c})
		}

		table.Render()
	}
	return err
}
