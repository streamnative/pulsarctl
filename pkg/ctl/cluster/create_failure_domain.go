package cluster

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func createFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for creating a failure domain of the <cluster-name>."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	create := pulsar.Example{
		Desc:    "creating the failure domain",
		Command: "pulsarctl clusters create-failure-domain --domain-name <domain-name> <cluster-name>",
	}
	examples = append(examples, create)

	createWithBrokers := pulsar.Example{
		Desc:    "creating the failure domain with brokers",
		Command: "pulsarctl clusters create-failure-domain --domain-name <domain-name> --broker-list <cluster-A>,<cluster-B> <cluster-name>",
	}
	examples = append(examples, createWithBrokers)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Create failure domain <domain-name> for cluster <cluster-name> succeed",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"create-failure-domain",
		"Create a failure domain",
		desc.ToString(),
		"cfd")

	var failureDomainData pulsar.FailureDomainData

	vc.SetRunFuncWithNameArg(func() error {
		return doCreateFailureDomain(vc, &failureDomainData)
	})

	vc.FlagSetGroup.InFlagSet("FailureDomainData", func(set *pflag.FlagSet) {
		set.StringVar(
			&failureDomainData.DomainName,
			"domain-name",
			"",
			"The failure domain name")
		set.StringArrayVarP(
			&failureDomainData.BrokerList,
			"broker-list",
			"b",
			[]string{""},
			"Set the failure domain clusters")
	})
}

func doCreateFailureDomain(vc *cmdutils.VerbCmd, failureDomain *pulsar.FailureDomainData) error {
	failureDomain.ClusterName = vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().CreateFailureDomain(*failureDomain)
	if err == nil {
		vc.Command.Printf(
			"Create failure domain [%s] for cluster [%s] succeed\n",
			failureDomain.DomainName, failureDomain.ClusterName)
	}
	return err
}
