package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetOffloadThresholdCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting offload threshold for a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	set := Example{
		Desc:    "Get offload threshold for a namespace <namespace-name>",
		Command: "pulsarctl namespace get-offload-threshold <namespace-name>",
	}
	desc.CommandExamples = append(examples, set)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "The offload threshold is <size> byte(s) for the namespace <namespace-name>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-offload-threshold",
		"Get offload threshold for a namespace",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetOffloadThreshold(vc)
	})
}

func doGetOffloadThreshold(vc *cmdutils.VerbCmd) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	threshold, err := admin.Namespaces().GetOffloadThreshold(*ns)
	if err == nil {
		vc.Command.Printf("The offload threshold is %d byte(s) for the namespace %s", threshold, ns.String())
	}

	return err
}
