package lookup

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func LookupTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for looking up a topic."
	desc.CommandPermission = "This command does not require permissions. "

	var examples []Example
	lookup := Example{
		Desc:    "Look up a topic <topic-name>",
		Command: "pulsarctl topic lookup <topic-name>",
	}
	desc.CommandExamples = append(examples, lookup)

	var out []Output
	successOut := Output{
		Desc: "",
		Out: "{\n" +
			"  \"brokerUlr\": \"\",\n" +
			"  \"brokerUrlTls\": \"\",\n" +
			"  \"httpUrl\": \"\",\n" +
			"  \"httpUrlTls\": \"\",\n" +
			"}",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out  = append(out, NamespaceErrors...)

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

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	lookup, err := admin.Topics().Lookup(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), lookup)
	}
	return err
}
