package compact

import (
	"time"

	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func CompactStatusCmd(vc *cmdutils.VerbCmd)  {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting status of compaction on a topic."
	desc.CommandPermission  = "This command requires tenant admin permissions."

	var examples []Example
	compactStatus := Example{
		Desc: "Get status of compaction of a persistent topic <topic-name>",
		Command: "pulsarctl topic compact-status <topic-name>",
	}
	desc.CommandExamples =  append(examples, compactStatus)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Compaction was a success",
	}

	notRun := Output{
		Desc: "Compaction is not running",
		Out: "Compaction has not been run for <topic-name> since broker startup",
	}

	running := Output{
		Desc: "Compaction is running",
		Out: "Compaction is currently running",
	}

	errorOut := Output{
		Desc: "Compaction is error",
		Out:  "Error in compaction",
	}
	out = append(out, successOut, notRun, running, errorOut, ArgError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"compact-status",
		"Get status of compaction on a topic",
		desc.ToString())

	var wait bool

	vc.SetRunFuncWithNameArg(func() error {
		return doCompactStatus(vc, wait)
	})
}

func doCompactStatus(vc *cmdutils.VerbCmd, wait bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if topic.GetDomain().String() != "persistent" {
		return errors.New("Need to provide a persistent topic.")
	}

	admin := cmdutils.NewPulsarClient()
	status, err := admin.Topics().CompactStatus(*topic)
	if err != nil {
		return err
	}

	for wait && status.Status  == RUNNING {
		time.Sleep( 1 * time.Second)
		status,  err = admin.Topics().CompactStatus(*topic)
		if err != nil {
			return err
		}
	}

	switch status.Status {
	case NOT_RUN:
		vc.Command.Printf("Compaction has not been run for %s since broker startup/n", topic.String())
	case RUNNING:
		vc.Command.Printf("Compaction is currently running/n")
	case SUCCESS:
		vc.Command.Printf("Compaction was a success/n")
	case ERROR:
		vc.Command.Printf("Error in Compaction/n")
		err = errors.New(status.LastError)
	}

	return err
}
