package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/olekukonko/tablewriter"
)

func listClustersCmd(vc *cmdutils.VerbCmd) {
	// update the description
	vc.SetDescription(
		"list",
		"List the available pulsar clusters",
		"This command is used for listing the list of available pulsar clusters.")

	// set the run function
	vc.SetRunFunc(func() error {
		return doListClusters(vc)
	})
}

func doListClusters(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	clusters, err := admin.Clusters().List()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Cluster Name"})

		for _, c := range clusters {
			table.Append([]string{c})
		}

		table.Render()
	}
	return err
}
