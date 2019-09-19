package permission

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GrantPermissionCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for granting permissions to a client role on a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	grant := Example{
		Desc:    "Grant permissions to a client on a single topic <topic-name>",
		Command: "pulsarctl topic grant-permissions --role <role> --actions <action-1> --actions <action-2> <topic-name>",
	}
	desc.CommandExamples = append(examples, grant)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Grant role %s and actions %v to the topic %s successfully",
	}

	flagError := Output{
		Desc: "the specified role is empty",
		Out:  "Invalid role name",
	}

	actionsError := Output{
		Desc: "the specified actions is not allowed.",
		Out: "The auth action  only can be specified as 'produce', " +
			"'consume', or 'functions'. Invalid auth action '<actions>'",
	}
	out = append(out, successOut, ArgError, flagError, actionsError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"grant-permissions",
		"Grant permissions to a client on a topic",
		desc.ToString(),
		"grant")

	var role string
	var actions []string

	vc.SetRunFuncWithNameArg(func() error {
		return doGrantPermission(vc, role, actions)
	})

	vc.FlagSetGroup.InFlagSet("GrantPermissions", func(set *pflag.FlagSet) {
		set.StringVar(&role, "role", "",
			"Client role to which grant permissions")
		set.StringSliceVar(&actions, "actions", []string{},
			"Actions to be granted (produce,consume,functions)")
		cobra.MarkFlagRequired(set, "role")
		cobra.MarkFlagRequired(set, "actions")
	})
}

func doGrantPermission(vc *cmdutils.VerbCmd, role string, actions []string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if role == "" {
		return errors.New("Invalid role name")
	}

	authActions, err := getAuthActions(actions)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().GrantPermission(*topic, role, authActions)
	if err == nil {
		vc.Command.Printf(
			"Grant permissions for the role %s and the actions %v to "+
				"the topic %s successfully\n", role, actions, topic.String())
	}

	return err
}

func getAuthActions(actions []string) ([]AuthAction, error) {
	var authActions []AuthAction
	for _, v := range actions {
		a, err := ParseAuthAction(v)
		if err != nil {
			return nil, err
		}
		authActions = append(authActions, a)
	}
	return authActions, nil
}
