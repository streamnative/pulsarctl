package cluster

import (
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	`github.com/streamnative/pulsarctl/pkg/pulsar`
)

func listClustersCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "List the existing clusters"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	create := pulsar.Example{
		Desc:    "List the existing clusters",
		Command: "pulsarctl clusters create-failure-domain (cluster-name) (domain-name)",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "+--------------+\n" +
				"| CLUSTER NAME |\n" +
				"+--------------+\n" +
				"| standalone   |\n" +
				"| test-a       |\n" +
				"+--------------+",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"list",
		"List the available pulsar clusters",
		"This command is used for listing the list of available pulsar clusters.",
		desc.ExampleToString(),
		"",
	)

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
		table.SetHeader([]string{ "Cluster Name" })

		for _, c := range clusters {
			table.Append([]string { c })
		}

		table.Render()
	}
	return err
}