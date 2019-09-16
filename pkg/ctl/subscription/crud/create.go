package crud

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func CreateCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for creating a subscription on a topic."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []Example
	create := Example{
		Desc:    "Create a subscription <subscription-name> on a topic <topic-name> from latest",
		Command: "pulsarctl subscription create <topic-name> <subscription-name>",
	}

	createWithFlag := Example{
		Desc:    "Create a subscription <subscription-name> on a topic <topic-name> from the specified position <position>",
		Command: "pulsarctl subscription create --messageId <position> <topic-name> <subscription-name>",
	}
	desc.CommandExamples = append(examples, create, createWithFlag)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Create subscription <subscription-name> on topic <topic-name> from <position> successfully",
	}
	out = append(out, successOut, ArgsError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	out = append(out, MessageIdErrors...)
	desc.CommandOutput = out

	var messageId string

	vc.SetDescription(
		"create",
		"Create subscription on a topic from latest",
		desc.ToString())

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doCreate(vc, messageId)
	}, CheckSubscriptionNameTwoArgs)

	vc.FlagSetGroup.InFlagSet("Create Subscription", func(set *pflag.FlagSet) {
		set.StringVarP(&messageId, "messageId", "m", "latest",
			"message id where to create the subscription. It can be either 'latest', "+
				"'earliest or (ledgerId:entryId)")
	})
}

func doCreate(vc *cmdutils.VerbCmd, id string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	sName := vc.NameArgs[1]

	var messageId MessageId
	switch id {
	case "latest":
		messageId = Latest
	case "earliest":
		messageId = Earliest
	default:
		i, err := ParseMessageId(id)
		if err != nil {
			return err
		}
		messageId = *i
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Subscriptions().Create(*topic, sName, messageId)
	if err == nil {
		vc.Command.Printf("Create subscription %s on topic %s from %s successfully", sName, topic.String(), id)
	}

	return err
}
