package info

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetLastMessageIdCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the last message id of a topic (partition)."
	desc.CommandPermission= "This command requires tenant admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic"

	var examples []Example
	get := Example{
		Desc:    "Get the last message id of a topic <persistent-topic-name>",
		Command: "pulsarctl topic last-message-id <persistent-topic-name>",
	}

	getPartitionedTopic := Example{
		Desc: "Get the last message id of a partition of a partitioned topic <topic-name>",
		Command: "pulsarctl topic last-message-id --partition <partition> <topic-name>",
	}
	desc.CommandExamples = append(examples, get, getPartitionedTopic)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "{\n" +
			"  \"LedgerId\": 0,\n" +
			"  \"EntryId\": 0,\n" +
			"  \"PartitionedIndex\": 0" +
			"\n}",
	}
	out = append(out, successOut, ArgError)

	topicNotFoundError := Output{
		Desc: "the topic <persistent-topic-name> does not exist in the cluster",
		Out: "[✖]  code: 404 reason: Topic not found",
	}
	out = append(out, topicNotFoundError)

	notAllowedError := Output{
		Desc: "the topic <persistent-topic-name> does not a persistent topic",
		Out: "[✖]  code: 405 reason: GetLastMessageId on a non-persistent topic is not allowed",
	}
	out = append(out, notAllowedError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	var partition int

	vc.SetDescription(
		"last-message-id",
		"Get the last message id of a topic",
		desc.ToString(),
		"lmi")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetLastMessageId(vc, partition)
	})

	vc.FlagSetGroup.InFlagSet("LastMessageId", func(set *pflag.FlagSet) {
		set.IntVarP(&partition, "partition", "p", -1,
			"The partitioned topic index value")
	})
}

func doGetLastMessageId(vc *cmdutils.VerbCmd, partition int) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if partition >= 0 {
		topic, err = topic.GetPartition(partition)
		if err != nil {
			return err
		}
	}

	admin := cmdutils.NewPulsarClient()
	messageId, err := admin.Topics().GetLastMessageId(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), messageId)
	}

	return err
}
