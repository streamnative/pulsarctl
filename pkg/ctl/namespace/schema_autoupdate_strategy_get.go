package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetSchemaAutoUpdateStrategyCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the schema auto-update strategy for a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []Example
	get := Example{
		Desc:    "Get the schema auto-update strategy for a namespace <namespace-name>",
		Command: "pulsarctl namespace get-schema-autoupdate-strategy <namespace-name>",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "The schema auto-update strategy is <strategy> for the namespace <namespace-name>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-schema-autoupdate-strategy",
		"Get the schema auto-update strategy for a namespace",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetSchemaAutoUpdateStrategy(vc)
	})
}

func doGetSchemaAutoUpdateStrategy(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	s, err := admin.Namespaces().GetSchemaAutoUpdateCompatibilityStrategy(*ns)
	if err == nil {
		vc.Command.Printf("The schema auto-update strategy is %s for the namespace %s", s.String(), ns.String())
	}

	return err
}
