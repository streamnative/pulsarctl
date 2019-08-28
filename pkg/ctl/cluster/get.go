package cluster

import (
	"encoding/json"
	"errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getClusterDataCmd(vc *cmdutils.VerbCmd)  {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "This command is used for getting the cluster data of the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc: "getting the <cluster-name> data",
		Command: "pulsarctl clusters get <cluster-name>",
	}
	examples = append(examples, get)

	desc.CommandExamples = examples
	desc.CommandOutput = "The configuration data of the specified cluster"

	vc.SetDescription(
		"get",
		"Get the configuration data for the specified cluster",
		desc.ToString(),
		"get")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetClusterData(vc)
	})
}

func doGetClusterData(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg
	if clusterName == "" {
		return errors.New("Should specified a cluster name")
	}

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
