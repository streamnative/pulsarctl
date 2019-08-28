package cluster

import (
	"errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func deleteClusterCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "This command is used for deleting an existing cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	delete := pulsar.Example{
		Desc:    "deleting the cluster named <cluster-name>",
		Command: "pulsarctl clusters delete <cluster-name>",
	}
	examples = append(examples, delete)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Cluster <cluster-name> delete successfully.",
	}
	out = append(out, successOut)

	failOut := pulsar.Output{
		Desc: "output of doesn't specified a cluster name",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}
	out = append(out, failOut)

	clusterNotExist := pulsar.Output{
		Desc: "output of cluster doesn't exist",
		Out:  "[✖]  code: 404 reason: Cluster does not exist",
	}
	out = append(out, clusterNotExist)

	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete an existing cluster",
		desc.ToString(),
		"delete")

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteCluster(vc)
	})
}

func doDeleteCluster(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().Delete(clusterName)
	if err == nil {
		vc.Command.Printf("Cluster %s delete successfully", clusterName)
	}
	return err
}
