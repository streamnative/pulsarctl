package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetCompactionThresholdCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting compaction threshold for a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	set := Example{
		Desc:    "Get compaction threshold for a namespace <namespace-name>",
		Command: "pulsarctl namespace get-compaction-threshold <namespace-name>",
	}
	desc.CommandExamples = append(examples, set)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "The compaction threshold is <size> byte(s) for the namespace <namespace-name>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-compaction-threshold",
		"Get compaction threshold for a namespace",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetCompactionThreshold(vc)
	})
}

func doGetCompactionThreshold(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	threshold, err := admin.Namespaces().GetCompactionThreshold(*ns)
	if err == nil {
		vc.Command.Printf("The compaction threshold is %d byte(s) for the namespace %s", threshold, ns.String())
	}

	return err
}
