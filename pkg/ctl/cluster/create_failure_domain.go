package cluster

import (
	"github.com/pkg/errors"
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
		Desc:    "create the failure domain",
		Command: "pulsarctl clusters create-failure-domain <cluster-name> <domain-name>",
	}
	examples = append(examples, create)

	createWithBrokers := pulsar.Example{
		Desc:    "create the failure domain with brokers",
		Command: "pulsarctl clusters create-failure-domain" +
			" --broker-list <cluster-A> --broker-list <cluster-B> <cluster-name> <domain-name>",
	}
	examples = append(examples, createWithBrokers)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Create failure domain <domain-name> for cluster <cluster-name> succeed",
	}
	out = append(out, successOut)

	argsErrorOut := pulsar.Output{
		Desc:"the args need to be specified as <cluster-name> <domain-name>",
		Out: "[✖]  need specified two names for cluster and failure domain",
	}
	out =append(out, argsErrorOut)
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

	checkArgs := func(args []string) error {
		if len(args) != 2 {
			return errors.New("need to specified two names for cluster and failure domain")
		}
		return nil
	}

	vc.SetRunFuncWithNameArgs(func() error {
		return doCreateFailureDomain(vc, &failureDomainData)
	}, checkArgs)

	vc.FlagSetGroup.InFlagSet("FailureDomainData", func(set *pflag.FlagSet) {
		set.StringSliceVarP(
			&failureDomainData.BrokerList,
			"broker-list",
			"b",
			[]string{""},
			"Set the failure domain clusters")
	})
}

func doCreateFailureDomain(vc *cmdutils.VerbCmd, failureDomain *pulsar.FailureDomainData) error {
	failureDomain.ClusterName = vc.NameArgs[0]
	failureDomain.DomainName = vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().CreateFailureDomain(*failureDomain)
	if err == nil {
		vc.Command.Printf(
			"Create failure domain [%s] for cluster [%s] succeed\n",
			failureDomain.DomainName, failureDomain.ClusterName)
	}
	return err
}
