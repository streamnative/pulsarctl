package info

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func GetLastMessageIDCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting the last message id of a topic (partition)."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Get the last message id of a topic <persistent-topic-name>",
		Command: "pulsarctl topic last-message-id <persistent-topic-name>",
	}

	getPartitionedTopic := pulsar.Example{
		Desc:    "Get the last message id of a partition of a partitioned topic <topic-name>",
		Command: "pulsarctl topic last-message-id --partition <partition> <topic-name>",
	}
	examples = append(examples, get, getPartitionedTopic)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"LedgerID\": 0,\n" +
			"  \"EntryId\": 0,\n" +
			"  \"PartitionedIndex\": 0" +
			"\n}",
	}
	out = append(out, successOut, e.ArgError)

	topicNotFoundError := pulsar.Output{
		Desc: "the topic <persistent-topic-name> does not exist in the cluster",
		Out:  "[✖]  code: 404 reason: Topic not found",
	}
	out = append(out, topicNotFoundError)

	notAllowedError := pulsar.Output{
		Desc: "the topic <persistent-topic-name> does not a persistent topic",
		Out:  "[✖]  code: 405 reason: GetLastMessageID on a non-persistent topic is not allowed",
	}
	out = append(out, notAllowedError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	var partition int

	vc.SetDescription(
		"last-message-id",
		"Get the last message id of a topic",
		desc.ToString(),
		"lmi")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetLastMessageID(vc, partition)
	})

	vc.FlagSetGroup.InFlagSet("LastMessageId", func(set *pflag.FlagSet) {
		set.IntVarP(&partition, "partition", "p", -1,
			"The partitioned topic index value")
	})
}

func doGetLastMessageID(vc *cmdutils.VerbCmd, partition int) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArg)
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
	messageID, err := admin.Topics().GetLastMessageID(*topic)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), messageID)
	}

	return err
}
