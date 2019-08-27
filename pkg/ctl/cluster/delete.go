package cluster

import (
	"errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func deleteClusterCmd(vc *cmdutils.VerbCmd)  {
	vc.SetDescription(
		"delete",
		"Delete a pulsar cluster",
		"Delete a pulsar cluster",
		"delete")

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteCluster(vc)
	})
}

func doDeleteCluster(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg
	if clusterName == "" {
		return errors.New("Should specified a cluster ")
	}

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().Delete(clusterName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Cluster %s delete successfully", clusterName)
	}
	return err
}
