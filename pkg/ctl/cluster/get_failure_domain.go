package cluster

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting the specified failure domain on the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc:    "getting the broker list in the <cluster-name> cluster failure domain <domain-name>",
		Command: "pulsarctl clusters get-failure-domain -n <domain-name> <cluster-name>",
	}
	examples = append(examples, get)

	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "output example",
		Out: "{\n  " +
			"\"brokers\" : [\n" +
			"    \"failure-broker-A\",\n" +
			"    \"failure-broker-B\",\n" +
			"  ]\n" +
			"}",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	var failureDomainData pulsar.FailureDomainData

	vc.SetDescription(
		"get-failure-domain",
		"Get the failure domain",
		desc.ToString(),
		"gfd")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetFailureDomain(vc, &failureDomainData)
	})

	vc.FlagSetGroup.InFlagSet("FailureDomain", func(set *pflag.FlagSet) {
		set.StringVarP(
			&failureDomainData.DomainName,
			"domain-name",
			"n",
			"",
			"The failure domain name")
	})

}

func doGetFailureDomain(vc *cmdutils.VerbCmd, data *pulsar.FailureDomainData) error {
	data.ClusterName = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	resFailureDomain, err := admin.Clusters().GetFailureDomain(data.ClusterName, data.DomainName)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), resFailureDomain)
	}

	return err
}
