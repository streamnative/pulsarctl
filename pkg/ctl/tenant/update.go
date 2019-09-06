package tenant

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func updateTenantCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for updating the configuration of a tenant."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	empty := pulsar.Example{
		Desc:    "clear the tenant configuration of a tenant",
		Command: "pulsarctl tenant update <tenant-name>",
	}
	examples = append(examples, empty)

	updateAdminRole := pulsar.Example{
		Desc:    "update the admin roles for tenant <tenant-name>",
		Command: "pulsarctl tenants update --admin-roles <admin-A> --admin-roles <admin-B> <tenant-name>",
	}
	examples = append(examples, updateAdminRole)

	updateClusters := pulsar.Example{
		Desc:    "update the allowed cluster access list for tenant <tenant-name>",
		Command: "pulsarctl tenants update --allowed-clusters <cluster-A> --allowed-clusters <cluster-B> <tenant-name>",
	}
	examples = append(examples, updateClusters)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Update tenant [%s] successfully",
	}
	out = append(out, successOut)

	notExist := pulsar.Output{
		Desc: "the specified tenant does not exist in",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}
	out = append(out, tenantNameArgsError, notExist)

	flagErrorOut := pulsar.Output{
		Desc: "the flag --admin-roles and --allowed-clusters are not specified",
		Out:  "[✖]  the admin roles or the allowed clusters is not specified",
	}
	out = append(out, flagErrorOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"update",
		"update the tenant admin roles and cluster access list",
		desc.ToString(),
		"u")

	var data pulsar.TenantData

	vc.SetRunFuncWithNameArg(func() error {
		return doUpdateTenant(vc, &data)
	})

	vc.FlagSetGroup.InFlagSet("TenantData", func(set *pflag.FlagSet) {
		set.StringSliceVarP(
			&data.AdminRoles,
			"admin-roles",
			"r",
			nil,
			"Allowed admins to access the tenant")
		set.StringSliceVarP(
			&data.AllowedClusters,
			"allowed-clusters",
			"c",
			nil,
			"Allowed clusters")
	})
}

func doUpdateTenant(vc *cmdutils.VerbCmd, data *pulsar.TenantData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	data.Name = vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Tenants().Update(*data)
	if err == nil {
		vc.Command.Printf("Update tenant %s successfully\n", data.Name)
	}
	return err
}
