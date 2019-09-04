package tenant

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func deleteTenantCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription

	desc.CommandUsedFor = "This command is used for deleting a tenant and all namespaces and topics under it will be deleted."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	deleteExample := pulsar.Example{
		Desc:    "delete the tenant named <tenant-name>",
		Command: "pulsarctl tenants delete <tenant-name>",
	}
	examples = append(examples, deleteExample)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Delete tenant <tenant-name> successfully",
	}
	out = append(out, successOut)
	out = append(out, tenantNameArgsError, tenantNotExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"d",
		desc.ToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteTenant(vc)
	})
}

func doDeleteTenant(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	admin := cmdutils.NewPulsarClient()
	err := admin.Tenants().Delete(vc.NameArg)
	if err == nil {
		vc.Command.Printf("Delete tenant [%s] successfully\n", vc.NameArg)
	}
	return err
}
