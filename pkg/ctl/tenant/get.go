package tenant

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getTenantCmd(vc *cmdutils.VerbCmd)  {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting a tenant info."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	getSuccess := pulsar.Example{
		Desc: "get the tenant named <tenant-name>",
		Command:"pulsarctl tenants get <tenant-name>",
	}
	examples = append(examples, getSuccess)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out :  "{\n" +
			"  \"adminRoles\": [],\n" +
			"  \"allowedClusters\": [\n" +
			"    \"standalone\"\n" +
			"  ]\n" +
			"}",
	}
	out = append(out, successOut)
	out = append(out, tenantNameArgsError)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"get the tenant info for the specified tenant",
		desc.ToString(),
		"g")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetTenant(vc)
	})
}

func doGetTenant(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	data, err := admin.Tenants().Get(vc.NameArg)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), data)
	}
	return err
}
