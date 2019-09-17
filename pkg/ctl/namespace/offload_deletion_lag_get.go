package namespace

import (
	"fmt"
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetOffloadDeletionLagCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting offload deletion for a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get offload deletion for a namespace <namespace-name>",
		Command: "pulsarctl namespace get-offload-deletion-lag <namespace-name>",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "The offload deletion lag is <n> minute(s) for the namespace <namespace-name>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-offload-deletion-lag",
		"Get offload deletion lag for a namespace",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetOffloadDeletionLag(vc)
	})
}

func doGetOffloadDeletionLag(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	ms, err := admin.Namespaces().GetOffloadDeleteLag(*ns)
	if err == nil {
		t, _ := time.ParseDuration(fmt.Sprintf("%dms", ms))
		vc.Command.Printf("The offload deletion lag is %f minute(s) for the namespace %s", t.Minutes(), ns.String())
	} else if err != nil && ms == 0 {
		vc.Command.Printf("The offload deletion lag is not set for the namespace %s", ns.String())
		err = nil
	}

	return err
}
