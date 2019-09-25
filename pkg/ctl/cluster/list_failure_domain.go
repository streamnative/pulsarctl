package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func listFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting all failure domain under the cluster <cluster-name>."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	list := pulsar.Example{
		Desc:    "listing all the failure domains under the specified cluster",
		Command: "pulsarctl clusters list-failure-domains <cluster-name>",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "output example",
		Out: "{\n" +
			"  \"failure-domain\": {\n" +
			"    \"brokers\": [\n" +
			"      \"failure-broker-A\",\n" +
			"      \"failure-broker-B\"\n" +
			"    ]\n" +
			"  }\n" +
			"}",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"list-failure-domains",
		"List the existing failure domains for a cluster",
		desc.ToString(),
		"lfd")

	vc.SetRunFuncWithNameArg(func() error {
		return doListFailureDomain(vc)
	})
}

func doListFailureDomain(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	clusterName := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	domainData, err := admin.Clusters().ListFailureDomains(clusterName)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), domainData)
	}
	return err
}
