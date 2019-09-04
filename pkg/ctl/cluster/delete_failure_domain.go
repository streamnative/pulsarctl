package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func deleteFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for deleting the failure domain <domain-name> of the cluster <cluster-name>"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	delete := pulsar.Example{
		Desc:    "delete the failure domain",
		Command: "pulsarctl clusters delete-failure-domain <cluster-name> <domain-name>",
	}
	examples = append(examples, delete)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "output example",
		Out:  "Delete failure domain [<domain-name>] for cluster [<cluster-name>] succeed",
	}
	out = append(out, successOut, failureDomainArgsError, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete-failure-domain",
		"Delete a failure domain",
		desc.ToString(),
		"dfd")


	vc.SetRunFuncWithNameArgs(func() error {
		return doDeleteFailureDomain(vc)
	}, checkFailureDomainArgs)
}

func doDeleteFailureDomain(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	var failureDomain pulsar.FailureDomainData
	failureDomain.ClusterName = vc.NameArgs[0]
	failureDomain.DomainName = vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().DeleteFailureDomain(failureDomain)
	if err == nil {
		vc.Command.Printf("Delete failure domain [%s] for cluster [%s] succeed\n", failureDomain.DomainName, failureDomain.ClusterName)
	}

	return err
}
