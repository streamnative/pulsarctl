package permission

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func RevokePermissions(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for revoking permissions on a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	revoke := Example{
		Desc:    "Revoke permissions on a topic <topic-name>",
		Command: "pulsarctl topic revoke-permissions --role <role> <topic-name>",
	}
	examples = append(examples, revoke)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Revoke permissions for the role <role> to the topic <topic-name> successfully\n",
	}

	flagError := Output{
		Desc: "the specified role is empty",
		Out:  "Invalid role name",
	}
	out = append(out, successOut, flagError, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"revoke-permissions",
		"Revoke permissions on topic",
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
	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().RevokePermission(*topic, role)
	if err == nil {
		vc.Command.Printf("Revoke permissions for the role %s to "+
			"the topic %s successfully\n", role, topic.String())
	}

	return err
}
