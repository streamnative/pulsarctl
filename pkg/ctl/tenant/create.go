package tenant

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func createTenantCmd(vc *cmdutils.VerbCmd)  {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for creating a new tenant."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	create := pulsar.Example{
		Desc: "create a tenant named <tenant-name>",
		Command: "pulsarctl tenants create <tenant-name>",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "Create tenant <tenant-name> successfully",
	}
	out = append(out, successOut)
	out = append(out, tenantNameArgsError, tenantAlreadyExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Create a tenant",
		desc.ToString(),
		"create")

	var tenantData pulsar.TenantData

	vc.SetRunFuncWithNameArg(func() error {
		return doCreateTenant(vc, &tenantData)
	})

	vc.FlagSetGroup.InFlagSet("TenantData", func(set *pflag.FlagSet) {
		set.StringSliceVarP(
			&tenantData.AdminRoles,
			"admin-roles",
			"r",
			[]string{""},
			"Allowed admins to access the tenant")
		set.StringSliceVarP(
			&tenantData.AllowedClusters,
			"allowed-clusters",
			"c",
			[]string{""},
			"Allowed clusters")
	})
}

func doCreateTenant(vc *cmdutils.VerbCmd, data *pulsar.TenantData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	data.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Tenants().Create(*data)
	if err == nil {
		vc.Command.Printf("Create tenant %s successfully\n", data.Name)
	}

	return err
}
