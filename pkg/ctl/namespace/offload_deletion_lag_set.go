package namespace

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func SetOffloadDeletionLagCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for setting offload deletion for a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []Example
	set := Example{
		Desc:    "Set offload deletion <duration> for a namespace <namespace-name>",
		Command: "pulsarctl namespace set-offload-deletion-lag --lag <duration> <namespace-name>",
	}
	desc.CommandExamples = append(examples, set)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Set offload deletion lag <duration> for the namespace <namespace-name>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-offload-deletion-lag",
		"Set offload deletion lag for a namespace",
		desc.ToString())

	var d string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetOffloadDeletionLag(vc, d)
	})

	vc.FlagSetGroup.InFlagSet("Offload Deletion Lag", func(set *pflag.FlagSet) {
		set.StringVarP(&d, "lag", "l", "",
			"Duration  to wait after offloading a ledger segment, before deleting the copy of that segment "+
				"from cluster local storage. (e.g. 1s, 1m, 1h, 1d, 1w, 1y)")
		cobra.MarkFlagRequired(set, "lag")
	})
}

func doSetOffloadDeletionLag(vc *cmdutils.VerbCmd, d string) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	t, err := parseRelativeTimeInSeconds(d)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetOffloadDeleteLag(*ns, t.Nanoseconds()/1e6)
	if err == nil {
		vc.Command.Printf("Set offload deletion lag %s for the namespace %s", d, ns.String())
	}

	return err
}
