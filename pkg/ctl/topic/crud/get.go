package crud

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetTopicCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting the metadata of an exist topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example
	getTopic := pulsar.Example{
		Desc:    "Get hte metadata of an exist topic <topic-name> metadata",
		Command: "pulsarctl topics get <topic-name>",
	}
	examples = append(examples, getTopic)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"partitions\": \"<partitions>\"\n" +
			"}",
	}
	out = append(out, successOut, e.ArgError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
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

	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	meta, err := admin.Topics().GetMetadata(*topic)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), meta)
	}

	return err
}
