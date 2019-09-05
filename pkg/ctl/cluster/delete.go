package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar/common"
)

func deleteClusterCmd(vc *cmdutils.VerbCmd) {
	desc := LongDescription{}
	desc.CommandUsedFor = "This command is used for deleting an existing cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	delete := Example{
		Desc:    "deleting the cluster named <cluster-name>",
		Command: "pulsarctl clusters delete <cluster-name>",
	}
	examples = append(examples, delete)
	desc.CommandExamples = examples

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Cluster <cluster-name> delete successfully.",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
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
