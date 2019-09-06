package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the specified failure domain on the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	get := Example{
		Desc:    "getting the broker list in the <cluster-name> cluster failure domain <domain-name>",
		Command: "pulsarctl clusters get-failure-domain <cluster-name> <domain-name>",
	}
	examples = append(examples, get)

	desc.CommandExamples = examples

	var out []Output
	successOut := Output{
		Desc: "output example",
		Out: "{\n  " +
			"\"brokers\" : [\n" +
			"    \"failure-broker-A\",\n" +
			"    \"failure-broker-B\",\n" +
			"  ]\n" +
			"}",
	}
	out = append(out, successOut, failureDomainArgsError, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-failure-domain",
		"Get the failure domain",
		desc.ToString(),
		"gfd")

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doGetFailureDomain(vc)
	}, checkFailureDomainArgs)
}

func doGetFailureDomain(vc *cmdutils.VerbCmd) error {
	// fot testing
	if vc.NameError != nil {
		return vc.NameError
	}

	clusterName := vc.NameArgs[0]
	domainName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	resFailureDomain, err := admin.Clusters().GetFailureDomain(clusterName, domainName)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), resFailureDomain)
	}

	return err
}
