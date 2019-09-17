package namespace

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func SetOffloadThresholdCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for setting offload threshold for a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []Example
	set := Example{
		Desc:    "Set offload threshold <size> for a namespace <namespace-name>",
		Command: "pulsarctl namespace set-offload-threshold --size <size> <namespace-name>",
	}
	desc.CommandExamples = append(examples, set)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Set the offload threshold <threshold> for the namespace <namespace-name>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-offload-threshold",
		"Set offload threshold for a namespace",
		desc.ToString())

	var threshold string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetOffloadThreshold(vc, threshold)
	})

	vc.FlagSetGroup.InFlagSet("Offload Threshold", func(set *pflag.FlagSet) {
		set.StringVarP(&threshold, "size", "s", "-1",
			"Maximum number of bytes stored in the pulsar cluster for a topic before data will  "+
				"start being automatically offloaded to longterm  storage (e.g. 10m, 16g, 3t, 100)\n"+
				"Negative values disable automatic offload.\n"+
				"0 triggers offloading as soon as possible.")
		cobra.MarkFlagRequired(set, "lag")
	})
}

func doSetOffloadThreshold(vc *cmdutils.VerbCmd, threshold string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	size, err := validateSizeString(threshold)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetOffloadThreshold(*ns, size)
	if err == nil {
		vc.Command.Printf("Set the offload threshold %d for the namespace %s", size, ns.String())
	}

	return err
}
