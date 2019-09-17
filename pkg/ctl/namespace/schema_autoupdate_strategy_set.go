package namespace

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func SetSchemaAutoUpdateStrategyCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for setting the schema auto-update strategy for a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []Example
	set := Example{
		Desc:    "Set the schema auto-update strategy <strategy>",
		Command: "pulsarctl namespace set-schema-autoupdate-strategy --compatibility <strategy> <namespace-name>",
	}
	desc.CommandExamples = append(examples, set)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Set the schema auto-update strategy <strategy> for a namespace <namespace-name>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-schema-autoupdate-strategy",
		"Set the schema auto-update strategy for a namespace",
		desc.ToString())

	var s string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSchemaAutoUpdateStrategy(vc, s)
	})

	vc.FlagSetGroup.InFlagSet("Schema Auto Update Strategy", func(set *pflag.FlagSet) {
		set.StringVarP(&s, "compatibility", "c", "",
			"Compatibility level required for new schemas created via a Producer. Possible values "+
				"(AutoUpdateDisabled, Backward, Forward, Full, AlwaysCompatible, BackwardTransitive, ForwardTransitive, FullTransitive)")
	})
}

func doSetSchemaAutoUpdateStrategy(vc *cmdutils.VerbCmd, strategy string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	s := AutoUpdateDisabled
	if strategy != "" {
		s, err = ParseSchemaAutoUpdateCompatibilityStrategy(strategy)
		if err != nil {
			return err
		}
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSchemaAutoUpdateCompatibilityStrategy(*ns, s)
	if err == nil {
		vc.Command.Printf("Set the schema auto-update strategy %s for the namespace %s", s.String(), ns.String())
	}

	return err
}
