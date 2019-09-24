package lookup

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetBundleRangeCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting namespace bundle range of a topic (partition)."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []Example
	get := Example{
		Desc:    "Get namespace bundle range of a topic (topic-name)",
		Command: "pulsarctl topic bundle-range (topic-name)",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "The bundle range of the topic (topic-name) is: (bundle-range)",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"bundle-range",
		"Get the namespace bundle range of a topic",
		desc.ToString(),
		desc.ExampleToString(),
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

	topic, err := GetTopicName(vc.NameArg)
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
