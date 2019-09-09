package crud

import (
	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"strconv"
)

func CreateTopicCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for creating topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example
	createNonPartitions := pulsar.Example{
		Desc:    "Create a non-partitioned topic <topic-name>",
		Command: "pulsarctl topics create <topic-name> 0",
	}
	examples = append(examples, createNonPartitions)

	create := pulsar.Example{
		Desc:    "Create a partitioned topic <topic-name> with <partitions-num> partitions",
		Command: "pulsarctl topics create <topic-name> <partition-num>",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Create topic <topic-name> with <partition-num> partitions successfully",
	}
	out = append(out, successOut, ArgsError, TopicAlreadyExistError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Create a topic with n partitions",
		desc.ToString(),
		"c")

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doCreateTopic(vc)
	}, CheckTopicNameArgs)
}

func doCreateTopic(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	partitions, err := strconv.Atoi(vc.NameArgs[1])
	if err != nil || partitions < 0 {
		return errors.Errorf("invalid partition number '%s'", vc.NameArgs[1])
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().Create(*topic, partitions)
	if err == nil {
		vc.Command.Printf("Create topic %s with %d partitions successfully\n", topic.String(), partitions)
	}

	return err
}
