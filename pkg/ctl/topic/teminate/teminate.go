package teminate

import (
	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func TerminateCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor  = "This command is used for terminating a non-partitioned topic and don't allow any more " +
		"messages to  be published."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	terminate := Example{
		Desc: "Terminate a non-partitioned topic <topic-name> and don't allow any messages to be published",
		Command: "pulsarctl topic terminate <topic-name>",
	}
	desc.CommandExamples = append(examples, terminate)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: "Topic <topic-name> successfully terminated at <message-id>",
	}

	partitionError := Output{
		Desc: "the specified is a partitioned topic",
		Out: "[✖]  code: 405 reason: Termination of a partitioned topic is not allowed",
	}
	out = append(out, successOut, ArgError, TopicNotFoundError, partitionError)
	out = append(out, TopicNameErrors...)
	out  = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"terminate",
		"Terminate a non-partitioned topic",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doTerminate(vc)
	})
}

func doTerminate(vc *cmdutils.VerbCmd) error {
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
	messageId, err :=admin.Topics().Terminate(*topic)
	if err == nil {
		vc.Command.Printf("Topic %s successfully terminated at %+v /n", topic.String(), messageId)
	}

	return err
}
