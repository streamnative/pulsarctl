package messages

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func ResetCursorCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for resetting position for " +
		"subscription to position closest to timestamp or messageId"
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions"

	var examples []Example
	resetCursorTime := Example{
		Desc:    "Reset position for subscription <subscription-name> to position closest to timestamp <time>",
		Command: "pulsarctl reset --time <time> <topic-name> <subscription-name>",
	}

	resetCursorMessageId := Example{
		Desc:    "Reset position for subscription <subscription-name> to position closest to message id <message-id>",
		Command: "pulsarctl reset --message-id <message-id> <topic-name> <subscription-name>",
	}
	desc.CommandExamples = append(examples, resetCursorTime, resetCursorMessageId)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Reset the cursor to <time>/<message-id> successfully",
	}

	resetFlagError := Output{
		Desc: "the time is not specified or the message id is not specified",
		Out:  "[✖]  The reset position must be specified",
	}

	out = append(out, successOut, ArgsError, resetFlagError, TopicNotFoundError, SubNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"reset-cursor",
		"Reset the cursor to position closest to timestamp or messageId",
		desc.ToString(),
		"reset")

	var t string
	var mId string

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doResetCursor(vc, t, mId)
	}, CheckSubscriptionNameTwoArgs)

	vc.FlagSetGroup.InFlagSet("ResetCursor", func(set *pflag.FlagSet) {
		set.StringVarP(&t, "time", "t", "",
			"time to reset back to (e.g. 1s, 1m, 1h, 1d, 1w, 1y)")
		set.StringVarP(&mId, "message-id", "m", "",
			"message id to reset back to (e.g. ledgerId:entryId)")
	})
}

func doResetCursor(vc *cmdutils.VerbCmd, t, mId string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	sName := vc.NameArgs[1]

	// TODO: use isPersistent
	if topic.GetDomain().String() != "persistent" {
		return errors.New("")
	}

	admin := cmdutils.NewPulsarClient()
	if t != "" {
		d, err := parseRelativeTimeInSeconds(t)
		if err != nil {
			return err
		}
		resetTime := time.Now().Add(-d).UnixNano() / 1e6
		err = admin.Subscriptions().ResetCursorWithTimestamp(*topic, sName, resetTime)
		if err == nil {
			vc.Command.Printf("Reset the cursor to %s successfully", t)
		}

		return err
	} else if mId != "" {
		id, err := ParseMessageId(mId)
		if err != nil {
			return err
		}
		err = admin.Subscriptions().ResetCursorWithMessageId(*topic, sName, *id)
		if err == nil {
			vc.Command.Printf("Reset the cursor to %s successfully", mId)
		}
		return err
	} else {
		return errors.New("The reset position must be specified")
	}
}

func parseRelativeTimeInSeconds(relativeTime string) (time.Duration, error) {
	if relativeTime == "" {
		return -1, errors.New("Time can not be empty.")
	}

	unitTime := relativeTime[len(relativeTime)-1:]
	t := relativeTime[:len(relativeTime)-1]
	timeValue, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return -1, errors.Errorf("Invalid time '%s'", t)
	}

	switch strings.ToLower(unitTime) {
	case "s":
		return time.Duration(timeValue) * time.Second, nil
	case "m":
		return time.Duration(timeValue) * time.Minute, nil
	case "h":
		return time.Duration(timeValue) * time.Hour, nil
	case "d":
		return time.Duration(timeValue) * time.Hour * 24, nil
	case "w":
		return time.Duration(timeValue) * time.Hour * 24 * 7, nil
	case "y":
		return time.Duration(timeValue) * time.Hour * 24 * 7 * 365, nil
	default:
		return -1, errors.Errorf("Invalid time unit '%s'", unitTime)
	}
}
