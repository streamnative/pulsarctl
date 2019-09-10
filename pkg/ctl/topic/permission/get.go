package permission

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetPermissionsCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the permissions on a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get the permissions on a topic <topic-name>",
		Command: "pulsarctl topic get-permissions <topic-name>",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"<role>\": [\n" +
			"    \"<action>\"\n" +
			"  ]" +
			"\n}",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-permissions",
		"Get the permissions on a topic",
		desc.ToString(),
		"get")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetPermissions(vc)
	})
}

func doGetPermissions(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	permissions, err := admin.Topics().GetPermissions(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), permissions)
	}

	return err
}
