package tenant

import (
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func listTenantCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for listing all exist tenants."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	listSuccess := pulsar.Example{
		Desc:    "list all exist tenants",
		Command: "pulsarctl tenants list",
	}
	examples = append(examples, listSuccess)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "+-------------+\n" +
			"| TENANT NAME |\n" +
			"+-------------+\n" +
			"| public      |\n" +
			"| sample      |\n" +
			"+-------------+",
	}
	out = append(out, successOut)
	out = append(out, tenantNameArgsError)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all exist tenants",
		desc.ToString(),
		"l")

	vc.SetRunFunc(func() error {
		return doListTenant(vc)
	})
}

func doListTenant(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	admin := cmdutils.NewPulsarClient()
	tenants, err := admin.Tenants().List()
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Tenant Name"})

		for _, t := range tenants {
			table.Append([]string{t})
		}

		table.Render()
	}
	return err
}
