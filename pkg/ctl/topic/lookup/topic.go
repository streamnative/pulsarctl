package lookup

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func TopicCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for looking up the owner broker of a topic."
	desc.CommandPermission = "This command does not require permissions. "

	var examples []pulsar.Example
	lookup := pulsar.Example{
		Desc:    "Lookup the owner broker of the topic <topic-name>",
		Command: "pulsarctl topic lookup <topic-name>",
	}
	examples = append(examples, lookup)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "",
		Out: "{\n" +
			"  \"brokerUlr\": \"\",\n" +
			"  \"brokerUrlTls\": \"\",\n" +
			"  \"httpUrl\": \"\",\n" +
			"  \"httpUrlTls\": \"\",\n" +
			"}",
	}
	out = append(out, successOut, e.ArgError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"lookup",
		"Look up a topic",
		desc.ToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doLookupTopic(vc)
	})
}

func doLookupTopic(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	lookup, err := admin.Topics().Lookup(*topic)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), lookup)
	}
	return err
}
