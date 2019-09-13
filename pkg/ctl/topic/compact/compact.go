package compact

import (
	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func CompactCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for compacting a persistent topic."
	desc.CommandPermission = "This command is requires tenant admin permissions."

	var examples []Example
	compact := Example{
		Desc: "Compact a persistent topic <topic-name>",
		Command: "pulsarctl topic compact <topic-name>",
	}
	desc.CommandExamples = append(examples, compact)

	var out []Output
	successOut := Output{
		Desc:  "normal output",
		Out: "Sending compact topic <topic-name> request successfully",
	}
	out = append(out, successOut, ArgError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput  = out

	vc.SetDescription(
		"compact",
		"Compact a topic",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doCompact(vc)
	})
}

func doCompact(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if !topic.IsPersistent() {
		return errors.New("Need to provide a persistent topic.")
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().Compact(*topic)
	if err == nil {
		vc.Command.Printf("Sending compact topic %s request successfully/n", topic.String())
	}

	return err
}
