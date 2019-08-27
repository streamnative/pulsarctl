package cluster

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func updateClusterCmd(vc *cmdutils.VerbCmd)  {
	vc.SetDescription(
		"update",
		"Update a pulsar cluster",
		"",
		"update")

	clusterData := &pulsar.ClusterData{}

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
		flagSet.StringArrayVarP(
			&clusterData.PeerClusterNames,
			"peer-cluster",
			"p",
			[]string{""},
			"Cluster to be registered as a peer-cluster of this cluster.")
	})

}

func doUpdateCluster(vc *cmdutils.VerbCmd, clusterData *pulsar.ClusterData) error {
	clusterData.Name = vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().Update(*clusterData)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Cluster %s updated\n", clusterData.Name)
	}
	return err
}