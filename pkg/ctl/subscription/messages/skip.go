package messages

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func SkipCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for skipping messages for the subscription."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []Example
	skip := Example{
		Desc:    "Skip <count> messages for the subscription <subscription-name> under the topic <topic-name>",
		Command: "pulsarctl subscription skip-messages --count <count> <topic-name> <subscription-name>",
	}

	skipAll := Example{
		Desc:    "Skip all messages for  the subscription <subscription-name> under the topic <topic-name> (clear-backlog)",
		Command: "pulsarctl subscription skip-messages --count -1 <topic-name> <subscription-name>",
	}
	desc.CommandExamples = append(examples, skip, skipAll)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Subscription <subscription-name> skip <count> messages on topic <topic-name> successfully",
	}
	out = append(out, successOut, ArgsError, TopicNotFoundError, SubNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"skip-messages",
		"Skip messages for the subscription under a topic",
		desc.ToString(),
		"skip")

	var count int64

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doSkip(vc, count)
	}, CheckSubscriptionNameTwoArgs)

	vc.FlagSetGroup.InFlagSet("Skip Messages", func(set *pflag.FlagSet) {
		set.Int64VarP(&count, "count", "n", -1,
			"number of messages to skip")
		cobra.MarkFlagRequired(set, "count")
	})
}

func doSkip(vc *cmdutils.VerbCmd, count int64) error {
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
	if count < 0 {
		err = admin.Subscriptions().ClearBacklog(*topic, sName)
	} else {
		err = admin.Subscriptions().SkipMessages(*topic, sName, count)
	}

	if err == nil {
		vc.Command.Printf("Subscription %s skip %d messages on topic %s successfully",
			sName, count, topic.String())
	}

	return err
}
