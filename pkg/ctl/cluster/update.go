package cluster

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func UpdateClusterCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for updating the cluster data of the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example

	updateUrl := Example{
		Desc:    "updating the web service url of the <cluster-name>",
		Command: "pulsarctl clusters update --url http://example:8080 <cluster-name>",
	}
	examples = append(examples, updateUrl)

	updateUrlTls := Example{
		Desc:    "updating the tls secured web service url of the <cluster-name>",
		Command: "pulsarctl clusters update --url-tls https://example:8080 <cluster-name>",
	}
	examples = append(examples, updateUrlTls)

	updateBrokerUrl := Example{
		Desc:    "updating the broker service url of the <cluster-name>",
		Command: "pulsarctl clusters update --broker-url pulsar://example:6650 <cluster-name>",
	}
	examples = append(examples, updateBrokerUrl)

	updateBrokerUrlTls := Example{
		Desc:    "updating the tls secured web service url of the <cluster-name>",
		Command: "pulsarctl clusters update --broker-url-tls pulsar+ssl://example:6650 <cluster-name>",
	}
	examples = append(examples, updateBrokerUrlTls)

	updatePeerCluster := Example{
		Desc:    "registered as a peer-cluster of the <cluster-name> clusters",
		Command: "pulsarctl clusters update -p <cluster-a> -p <cluster-b> <cluster>",
	}
	examples = append(examples, updatePeerCluster)

	desc.CommandExamples = examples

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Cluster <cluster-name> updated",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"update",
		"Update a pulsar cluster",
		desc.ToString(),
		"update")

	clusterData := &ClusterData{}

	vc.SetRunFuncWithNameArg(func() error {
		return doUpdateCluster(vc, clusterData)
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
		flagSet.StringSliceVarP(
			&clusterData.PeerClusterNames,
			"peer-cluster",
			"p",
			[]string{""},
			"Cluster to be registered as a peer-cluster of this cluster.")
	})

}

func doUpdateCluster(vc *cmdutils.VerbCmd, clusterData *ClusterData) error {
	clusterData.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().Update(*clusterData)
	if err == nil {
		vc.Command.Printf("Cluster %s updated", clusterData.Name)
	}
	return err
}
