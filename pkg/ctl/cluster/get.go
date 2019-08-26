package cluster

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getClusterConfiguration(vc *cmdutils.VerbCmd)  {
	vc.SetDescription(
		"get",
		"Get the configuration data for the specified cluster",
		"This command is used for getting the configuration data of the specified cluster.",
		"get")

	var clusterName string

	vc.SetRunFuncWithNameArg(func() error {
		return doGetClusterConfiguration(vc, clusterName)
	})

	vc.FlagSetGroup.InFlagSet("ClusterName", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(&clusterName, "cluster-name", "", "Pulsar cluster name, e.g. pulsar-cluster")
		cobra.MarkFlagRequired(flagSet, "cluster-name")
	})

}

func doGetClusterConfiguration(vc *cmdutils.VerbCmd, clusterName string) error {
	admin := cmdutils.NewPulsarClient()
	clusterData, err := admin.Clusters().Get(clusterName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		s, err  := json.MarshalIndent(clusterData,  "", "    ")
		if err != nil {
			return err
		}
		vc.Command.Println(string(s))
	}

	return err
}
