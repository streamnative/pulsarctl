package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetSchemaValidationEnforcedCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the schema validation enforced."
	desc.CommandPermission = "This command requires super-user and tenant admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get schema validation status",
		Command: "pulsarctl namespace get-schema-validation-enforced <namespace-name>",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Schema validation enforced is enabled/disabled",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-schema-validation-enforced",
		"Enable/Disable schema validation enforced",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetSchemaValidationEnforced(vc)
	})
}

func doGetSchemaValidationEnforced(vc *cmdutils.VerbCmd) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	s, err := admin.Namespaces().GetSchemaValidationEnforced(*ns)
	if err == nil {
		out := "Schema validation enforced is "
		if s {
			out += "enabled"
		} else {
			out += "disabled"
		}
		vc.Command.Println(out)
	}

	return err
}
