package cluster

import (
	"encoding/json"
	"errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getClusterDataCmd(vc *cmdutils.VerbCmd) {
	desc := LongDescription{}
	desc.CommandUsedFor = "This command is used for getting the cluster data of the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	get := Example{
		Desc:    "getting the <cluster-name> data",
		Command: "pulsarctl clusters get <cluster-name>",
	}
	examples = append(examples, get)

	desc.CommandExamples = examples

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: "{\n  " +
			"\"serviceUrl\": \"http://localhost:8080\",\n  " +
			"\"serviceUrlTls\": \"\",\n  " +
			"\"brokerServiceUrl\": \"pulsar://localhost:6650\",\n  " +
			"\"brokerServiceUrlTls\": \"\",\n  " +
			"\"peerClusterNames\": null\n" +
			"}",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

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
		return errors.New("should specified a cluster name")
	}

	admin := cmdutils.NewPulsarClient()
	clusterData, err := admin.Clusters().Get(clusterName)
	if err == nil {
		s, err := json.MarshalIndent(clusterData, "", "    ")
		if err != nil {
			return err
		}
		vc.Command.Println(string(s))
	}

	return err
}
