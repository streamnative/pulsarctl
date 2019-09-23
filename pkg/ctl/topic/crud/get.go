package crud

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the metadata of an exist topic."
	desc.CommandPermission = "This command requires namespace admin permissions."
	desc.CommandScope = "non-partitioned topic, partitioned topic, a partition of a partitioned topic"

	var examples []Example
	getTopic := Example{
		Desc:    "Get hte metadata of an exist topic <topic-name> metadata",
		Command: "pulsarctl topics get <topic-name>",
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
		"get",
		"Get the specified topic metadata",
		desc.ToString(),
		"get")

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
	meta, err := admin.Topics().GetMetadata(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), meta)
	}

	return err
}
