package namespace

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func SetCompactionThresholdCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for setting compaction threshold for a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []Example
	set := Example{
		Desc:    "Set compaction threshold <size> for a namespace <namespace-name>",
		Command: "pulsarctl namespace set-compaction-threshold --size <size> <namespace-name>",
	}
	desc.CommandExamples = append(examples, set)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Set the compaction threshold <threshold> for the namespace <namespace-name>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-compaction-threshold",
		"Set compaction threshold for a namespace",
		desc.ToString())

	var threshold string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetCompactionThreshold(vc, threshold)
	})

	vc.FlagSetGroup.InFlagSet("Compaction Threshold", func(set *pflag.FlagSet) {
		set.StringVar(&threshold, "size", "0",
			"Maximum number of bytes in a topic backlog before compaction is triggered "+
				"(e.g. 10M, 16G, 3T). 0 disable automatic compaction")
		cobra.MarkFlagRequired(set, "size")
	})
}

func doSetCompactionThreshold(vc *cmdutils.VerbCmd, threshold string) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	size, err := validateSizeString(threshold)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetCompactionThreshold(*ns, size)
	if err == nil {
		vc.Command.Printf("Set the compaction threshold %d for the namespace %s", size, ns.String())
	}

	return err
}
