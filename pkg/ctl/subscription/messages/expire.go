package messages

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func ExpireCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for expiring messages that older than given expiry time (in seconds)" +
		" for the subscription."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []Example
	expire := Example{
		Desc: "Expire messages that older than given expire time (in seconds) <expire-time> for the subscription " +
			"<subscription-name> under a topic <topic-name>",
		Command: "pulsarctl subscription expire-messages --expire-time <expire-time> <topic-name> <subscription-name>",
	}

	expireAllSub := Example{
		Desc: "Expire message that older than given expire time (in second) <expire-time> for all subscriptions " +
			"under a topic <topic-name>",
		Command: "pulsarctl subscriptions expire-messages --expire-time <expire-time> <topic-name> all",
	}
	desc.CommandExamples = append(examples, expire, expireAllSub)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Expire messages after <time>(s) for subscription <subscription-name> of topic <topic-name> successfully",
	}
	out = append(out, successOut, ArgsError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"expire-messages",
		"Expiring messages that older than given expire time (in seconds)",
		desc.ToString(),
		"expire")

	var time int64

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doExpire(vc, time)
	}, CheckSubscriptionNameTwoArgs)

	vc.FlagSetGroup.InFlagSet("ExpireMessages", func(set *pflag.FlagSet) {
		set.Int64VarP(&time, "expire-time", "t", 0,
			"Expire messages older than time in seconds")
		cobra.MarkFlagRequired(set, "expire-time")
	})
}

func doExpire(vc *cmdutils.VerbCmd, time int64) error {
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
	if sName == "all" {
		err = admin.Subscriptions().ExpireAllMessages(*topic, time)
	} else {
		err = admin.Subscriptions().ExpireMessages(*topic, sName, time)
	}

	if err == nil {
		vc.Command.Printf("Expire messages after %d(s) for subscription %s of topic %s successfully",
			time, sName, topic.String())
	}

	return err
}
