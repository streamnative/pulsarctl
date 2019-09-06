package partitioned

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting an exist topic."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	getTopic := Example{
		Desc:    "Get an exist topic <topic-name> metadata",
		Command: "pulsarctl topics get-partitioned-topic-metadata <topic-name>",
	}
	desc.CommandExamples = append(examples, getTopic)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "{\n" +
			"  \"partitions\": \"<partitions>\"\n" +
			"}",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-partitioned-topic-metadata",
		"Get partitioned topic metadata",
		desc.ToString(),
		"gp")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetTopic(vc)
	})
}

func doGetTopic(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	meta, err := admin.Topics().GetPartitionedTopicMeta(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), meta)
	}

	return err
}
