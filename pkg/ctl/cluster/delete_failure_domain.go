package cluster

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func deleteFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for deleting the failure domain <domain-name> of the cluster <cluster-name>"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	delete := pulsar.Example{
		Desc:    "deleting the failure domain",
		Command: "pulsarctl clusters delete-failure-domain --domain-name <domain-name> <cluster-name>",
	}
	examples = append(examples, delete)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "output example",
		Out:  "Delete failure domain [<domain-name>] for cluster [<cluster-name>] succeed",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete-failure-domain",
		"Delete a failure domain",
		desc.ToString(),
		"dfd")

	var failureDomainData pulsar.FailureDomainData

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteFailureDomain(vc, &failureDomainData)
	})

	vc.FlagSetGroup.InFlagSet("FailureDomainData", func(set *pflag.FlagSet) {
		set.StringVarP(
			&failureDomainData.DomainName,
			"domain-name",
			"n",
			"",
			"The failure domain name")
	})
}

func doDeleteFailureDomain(vc *cmdutils.VerbCmd, data *pulsar.FailureDomainData) error {
	data.ClusterName = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().DeleteFailureDomain(*data)
	if err == nil {
		vc.Command.Printf("Delete failure domain [%s] for cluster [%s] succeed\n", data.DomainName, data.ClusterName)
	}

	return err
}
