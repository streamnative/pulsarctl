package partitioned

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"strconv"
)

func UpdateTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for updating an exist topic with new partition number."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	updateTopic := Example{
		Desc:    "",
		Command: "pulsarctl topics update-partitioned-topic <topic-name> <partition-num>",
	}
	desc.CommandExamples = append(examples, updateTopic)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Update topic [<topic-name>] with [<partition-num>] partitions successfully",
	}

	topicNotExist := Output{
		Desc: "the topic is not exist",
		Out:  "[✖]  code: 409 reason: Topic is not partitioned topic",
	}
	out = append(out, successOut, ArgsError, topicNotExist)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"update-partitioned-topic",
		"Update partitioned topic",
		desc.ToString(),
		"up")

	vc.SetRunFuncWithNameArgs(func() error {
		return doUpdateTopic(vc)
	}, CheckArgs)
}

func doUpdateTopic(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	partitions, err := strconv.Atoi(vc.NameArgs[1])

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().UpdatePartitionedTopic(*topic, partitions)
	if err == nil {
		vc.Command.Printf("Update topic [%s] with [%d] partitions successfully\n", topic.String(), partitions)
	}

	return err
}
