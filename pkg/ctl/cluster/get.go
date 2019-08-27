package cluster

import (
	"encoding/json"
	"errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)


var commandUsedFor = "This command is used for getting the cluster data of the specified cluster."
var commandExample =
		"{\n" +
		"    serviceUrl : http://localhost:8080, \n" +
		"    serviceUrlTls : https://localhost:8080, \n" +
		"    brokerServiceUrl: pulsar://localhost:6650, \n" +
		"    brokerServiceUrlTls: pulsar+ssl://localhost:6650, \n" +
		"    peerClusterNames: \"\" \n" +
		"}\n"
var commandPermission = "This command only admin can use."

func getClusterDataCmd(vc *cmdutils.VerbCmd)  {
	vc.SetDescription(
		"get",
		"Get the configuration data for the specified cluster",
		concat("\n"),
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

func concat(join string) string {
	return "USED FOR:" + join + "\t" + commandUsedFor  + join +
		"PERMISSION:" + join + "\t" + commandPermission + join +
		"EXAMPLE:" + join + commandExample
}
