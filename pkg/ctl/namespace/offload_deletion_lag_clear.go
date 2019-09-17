package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func ClearOffloadDeletionLagCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for clearing offload deletion for a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []Example
	clear := Example{
		Desc:    "Clear offload deletion for a namespace <namespace-name>",
		Command: "pulsarctl namespace clear-offload-deletion-lag <namespace-name>",
	}
	desc.CommandExamples = append(examples, clear)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Clear offload deletion lag for the namespace <namespace-name> successfully",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"clear-offload-deletion-lag",
		"Clear offload deletion lag for a namespace",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doClearOffloadDeletionLag(vc)
	})
}

func doClearOffloadDeletionLag(vc *cmdutils.VerbCmd) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().ClearOffloadDeleteLag(*ns)
	if err == nil {
		vc.Command.Printf("Clear offload deletion lag for the namespace %s successfully", ns.String())
	}

	return err
}
