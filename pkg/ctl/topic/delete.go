package topic

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func DeleteTopicCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for deleting an exist topic."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	deleteTopic := pulsar.Example{
		Desc:    "Delete a partitioned topic <topic-name>",
		Command: "pulsarctl topics delete <topic-name>",
	}

	deleteNonPartitionedTopic := pulsar.Example{
		Desc:    "Delete a non-partitioned topic <topic-name>",
		Command: "pulsarctl topics delete --non-partitioned <topic-name>",
	}

	desc.CommandExamples = append(examples, deleteTopic, deleteNonPartitionedTopic)
	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Delete topic <topic-name> successfully",
	}

	partitionedTopicNotExistError := pulsar.Output{
		Desc: "the partitioned topic does not exist",
		Out:  "[✖]  code: 404 reason: Partitioned topic does not exist",
	}

	nonPartitionedTopicNotExistError := pulsar.Output{
		Desc: "the non-partitioned topic does not exist",
		Out:  "[✖]  code: 404 reason: Topic not found",
	}
	out = append(out, successOut, ArgError,
		partitionedTopicNotExistError, nonPartitionedTopicNotExistError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete a topic",
		desc.ToString(),
		"d")

	var force bool
	var deleteSchema bool
	var nonPartitioned bool

	vc.FlagSetGroup.InFlagSet("Delete Topic", func(set *pflag.FlagSet) {
		set.BoolVarP(&nonPartitioned, "non-partitioned", "n", false,
			"Delete a non-partitioned topic")
		set.BoolVarP(&force, "force", "f", false,
			"Close all producer/consumer/replicator and delete topic forcefully")
		set.BoolVarP(&deleteSchema, "delete-schema", "d", false,
			"Delete schema while deleting topic")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteTopic(vc, force, deleteSchema, nonPartitioned)
	})
}

// TODO add delete schema
func doDeleteTopic(vc *cmdutils.VerbCmd, force, deleteSchema, nonPartitioned bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().Delete(*topic, force, nonPartitioned)
	if err == nil {
		vc.Command.Printf("Delete topic %s successfully\n", topic.String())
	}

	return err
}
