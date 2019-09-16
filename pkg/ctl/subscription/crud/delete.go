package crud

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func DeleteCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for deleting a durable subscription from a topic."
	desc.CommandPermission = "This command requires tenant admin and namespace consume permissions."

	var examples []Example
	deleteSub := Example{
		Desc:    "Delete a subscription <subscription-name> from a topic <topic-name>",
		Command: "pulsarctl subscription delete <topic-name> <subscription-name>",
	}
	desc.CommandExamples = append(examples, deleteSub)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Delete subscription %s on the topic %s  successfully",
	}
	out = append(out, successOut, ArgsError, SubNotFoundError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete a subscription on a topic",
		desc.ToString())

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doDelete(vc)
	}, CheckSubscriptionNameTwoArgs)
}

func doDelete(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	sName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	err = admin.Subscriptions().Delete(*topic, sName)
	if err == nil {
		vc.Command.Printf("Delete subscription %s on the topic %s successfully", sName, topic.String())
	}

	return err
}
