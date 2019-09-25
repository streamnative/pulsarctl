package permission

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func RevokePermissions(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for revoking a client role permissions on a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example
	revoke := pulsar.Example{
		Desc:    "Revoke permissions of a topic <topic-name>",
		Command: "pulsarctl topic revoke-permissions --role <role> <topic-name>",
	}
	examples = append(examples, revoke)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Revoke permissions for the role <role> of the topic <topic-name> successfully\n",
	}

	flagError := pulsar.Output{
		Desc: "the specified role is empty",
		Out:  "Invalid role name",
	}
	out = append(out, successOut, flagError, e.ArgError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"revoke-permissions",
		"Revoke a client role permissions on a topic",
		desc.ToString(),
		"revoke")

	var role string

	vc.SetRunFuncWithNameArg(func() error {
		return doRevokePermissions(vc, role)
	})

	vc.FlagSetGroup.InFlagSet("RevokePermissions", func(set *pflag.FlagSet) {
		set.StringVar(&role, "role", "", "Client role to which revoke permissions")
		cobra.MarkFlagRequired(set, "role")
	})
}

func doRevokePermissions(vc *cmdutils.VerbCmd, role string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	if role == "" {
		return errors.New("Invalid role name")
	}
	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().RevokePermission(*topic, role)
	if err == nil {
		vc.Command.Printf("Revoke permissions for the role %s of "+
			"the topic %s successfully\n", role, topic.String())
	}

	return err
}
