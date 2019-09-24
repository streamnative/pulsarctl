package cluster

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func CreateClusterCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for adding the configuration data for a cluster.\n"+
			"The configuration data is mainly used for geo-replication between clusters, so please make sure the service urls provided in this command are reachable between clusters.\n"+
			"This operation requires Pulsar super-user privileges."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	create := pulsar.Example{
		Desc:    "Provisions a new cluster",
		Command: "pulsarctl clusters create (cluster-name)",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Cluster (cluster-name) added",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"add",
		"Add a pulsar cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"create")

	clusterData := &pulsar.ClusterData{}

	// set the run function with name argument
	vc.SetRunFuncWithNameArg(func() error {
		return doCreateCluster(vc, clusterData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("ClusterData", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&clusterData.ServiceURL,
			"url",
			"",
			"Pulsar cluster web service url, e.g. http://example.pulsar.io:8080")
		flagSet.StringVar(
			&clusterData.ServiceURLTls,
			"url-tls",
			"",
			"Pulsar cluster tls secured web service url, e.g. https://example.pulsar.io:8443")
		flagSet.StringVar(
			&clusterData.BrokerServiceURL,
			"broker-url",
			"",
			"Pulsar cluster broker service url, e.g. pulsar://example.pulsar.io:6650")
		flagSet.StringVar(
			&clusterData.BrokerServiceURLTls,
			"broker-url-tls",
			"",
			"Pulsar cluster tls secured broker service url, e.g. pulsar+ssl://example.pulsar.io:6651")
		flagSet.StringArrayVarP(
			&clusterData.PeerClusterNames,
			"peer-cluster",
			"p",
			[]string{""},
			"Cluster to be registered as a peer-cluster of this cluster.")
	})
}

func doCreateCluster(vc *cmdutils.VerbCmd, clusterData *pulsar.ClusterData) error {
	clusterData.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().Create(*clusterData)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Cluster %s added\n", clusterData.Name)
	}
	return err
}
