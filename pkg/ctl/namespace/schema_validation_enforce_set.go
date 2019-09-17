package namespace

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func SetSchemaValidationEnforcedCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for setting the schema whether open schema validation enforced."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []Example
	enable := Example{
		Desc:    "Enable schema validation enforced",
		Command: "pulsarctl namespace set-schema-validation-enforced <namespace-name>",
	}

	disable := Example{
		Desc:    "Disable schema validation enforced",
		Command: "pulsarctl namespace set-schema-validation-enforced --disable <namespace-name>",
	}
	desc.CommandExamples = append(examples, enable, disable)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Enable/Disable schema validation enforced",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-schema-validation-enforced",
		"Enable/Disable schema validation enforced",
		desc.ToString())

	var d bool

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSchemaValidationEnforced(vc, d)
	})

	vc.FlagSetGroup.InFlagSet("Schema Validation Enforced", func(set *pflag.FlagSet) {
		set.BoolVarP(&d, "disable", "d", false,
			"Disable schema validation enforced")
	})
}

func doSetSchemaValidationEnforced(vc *cmdutils.VerbCmd, disable bool) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSchemaValidationEnforced(*ns, !disable)
	if err == nil {
		var out string
		if disable {
			out += "Disable "
		} else {
			out += "Enable "
		}
		vc.Command.Println(out + "schema validation enforced")
	}

	return err
}
