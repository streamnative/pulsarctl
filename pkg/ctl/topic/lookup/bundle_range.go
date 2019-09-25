package lookup

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetBundleRangeCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting namespace bundle range of a topic (partition)."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Get namespace bundle range of a topic <topic-name>",
		Command: "pulsarctl topic bundle-range <topic-name>",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "The bundle range of the topic <topic-name> is: <bundle-range>",
	}
	out = append(out, successOut, e.ArgError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"bundle-range",
		"Get the namespace bundle range of a topic",
		desc.ToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetBundleRange(vc)
	})
}

func doGetBundleRange(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	bundleRange, err := admin.Topics().GetBundleRange(*topic)
	if err == nil {
		vc.Command.Printf("The bundle range of the topic %s is: %s", topic.String(), bundleRange)
	}

	return err
}
