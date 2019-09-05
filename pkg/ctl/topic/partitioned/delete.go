package partitioned

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar/common"
)

func DeleteTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for deleting an exist partitioned topic."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	deleteTopic := Example{
		Desc:    "Delete a partitioned topic <topic-name>",
		Command: "pulsarctl topics delete-partitioned-topic <topic-name>",
	}
	desc.CommandExamples = append(examples, deleteTopic)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Delete topic [<topic-name>] successfully",
	}

	topicNotExistError := Output{
		Desc: "the topic is not exist",
		Out:  "[✖]  code: 404 reason: Partitioned topic does not exist",
	}
	out = append(out, successOut, ArgError, topicNotExistError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete-partitioned-topic",
		"Delete a partitioned topic",
		desc.ToString(),
		"dp")

	var force bool
	var deleteSchema bool

	vc.FlagSetGroup.InFlagSet("Delete Partitioned Topic", func(set *pflag.FlagSet) {
		set.BoolP("force", "f", false,
			"Close all producer/consumer/replicator and delete topic forcefully")
		set.BoolP("delete-schema", "d", false, "Delete schema while deleting topic")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteTopic(vc, force, deleteSchema)
	})
}

// TODO add delete schema
func doDeleteTopic(vc *cmdutils.VerbCmd, force, deleteSchema bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().DeletePartitionedTopic(*topic, force)
	if err == nil {
		vc.Command.Printf("Delete topic [%s] successfully\n", topic.String())
	}

	return err
}
