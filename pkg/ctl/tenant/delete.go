package tenant

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func deleteTenantCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription

	desc.CommandUsedFor = "This command is used for deleting an existing tenant."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	deleteExample := pulsar.Example{
		Desc:    "delete a tenant named (tenant-name)",
		Command: "pulsarctl tenants delete (tenant-name)",
	}
	examples = append(examples, deleteExample)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Delete tenant <tenant-name> successfully",
	}
	out = append(out, successOut)

	NonEmptyError := pulsar.Output{
		Desc: "there has namespace(s) under the tenant (tenant-name)",
		Out: "code: 409 reason: The tenant still has active namespaces",
	}
	out = append(out, tenantNameArgsError, tenantNotExistError, NonEmptyError)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"d",
		desc.ToString(),
		desc.ExampleToString(),
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
		vc.Command.Printf("Delete tenant %s successfully\n", vc.NameArg)
	}
	return err
}
