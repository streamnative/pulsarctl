package tenant

import (
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

var tenantNameArgsError = pulsar.Output{
	Desc: "the tenant name is not specified or the tenant name is specified more than one",
	Out:  "[✖]  only one argument is allowed to be used as a name",
}

var tenantNotExistError = pulsar.Output{
	Desc: "the specified tenant does not exist in the broker",
	Out:  "[✖]  code: 404 reason: The tenant does not exist",
}

var tenantAlreadyExistError = pulsar.Output{
	Desc: "the specified tenant has been created",
	Out:  "[✖]  code: 409 reason: Tenant already exists",
}

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"tenants",
		"Operations about tenant(s)",
		"",
		"tenant")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, createTenantCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteTenantCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, updateTenantCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, listTenantCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getTenantCmd)

	return resourceCmd
}
