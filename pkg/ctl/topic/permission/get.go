package permission

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetPermissionsCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting the permissions of a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Get the permissions of a topic <topic-name>",
		Command: "pulsarctl topic get-permissions <topic-name>",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"<role>\": [\n" +
			"    \"<action>\"\n" +
			"  ]" +
			"\n}",
	}
	out = append(out, successOut, e.ArgError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-permissions",
		"Get the permissions of a topic",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetPermissions(vc)
	})
}

func doGetPermissions(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	permissions, err := admin.Topics().GetPermissions(*topic)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), permissions)
	}

	return err
}
