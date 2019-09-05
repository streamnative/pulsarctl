package partitioned

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"strconv"
)

func CreateTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for creating partitioned topic."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	create := Example{
		Desc:    "Create topic <topic-name> with <partitions-num> partitions",
		Command: "pulsarctl topics create-partitioned-topic <topic-name> <partition-num>",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Create topic <topic-name> with <partition-num> partitions successfully",
	}
	out = append(out, successOut, ArgsError, TopicAlreadyExist)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"create-partitioned-topic",
		"Create partitioned topic",
		desc.ToString(),
		"cp")

	vc.SetRunFuncWithNameArgs(func() error {
		return doCreateTopic(vc)
	}, CheckArgs)
}

func doCreateTopic(vc *cmdutils.VerbCmd) error {
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
	err = admin.Topics().CreatePartitionedTopic(*topic, partitions)
	if err == nil {
		vc.Command.Printf("Create topic [%s] with [%d] partitions successfully\n", topic.String(), partitions)
	}

	return err
}
