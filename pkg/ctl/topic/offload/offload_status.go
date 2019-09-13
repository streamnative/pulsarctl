package offload

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func OffloadStatusCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for checking the status of data offloading" +
		" from a persistent topic to long-term storage."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	offloadStatus := Example{
		Desc:    "Check the status of data offloading from a topic <persistent-topic-name> to long-term storage",
		Command: "pulsarctl topic offload-status <persistent-topic-name>",
	}

	waiting := Example{
		Desc:    "Wait for offloading to complete",
		Command: "pulsarctl topic offload-status --wait <persistent-topic-name>",
	}
	desc.CommandExamples = append(examples, offloadStatus, waiting)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Offload was a success",
	}

	notRun := Output{
		Desc: "Offloading is not running",
		Out: "Offload has not been run for <topic-name> since broker startup",
	}

	running := Output{
		Desc: "Offloading is running",
		Out: "Offload is currently running",
	}

	errorOut := Output{
		Desc: "Offload is error",
		Out:  "Error in offload",
	}
	out = append(out, successOut, notRun, running, errorOut, ArgError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"offload-status",
		"Check the status of data offloading",
		desc.ToString())

	var wait bool

	vc.SetRunFuncWithNameArg(func() error {
		return doOffloadStatus(vc, wait)
	})

	vc.FlagSetGroup.InFlagSet("OffloadStatus", func(set *pflag.FlagSet) {
		set.BoolVarP(&wait, "wait", "w", false, "Wait for offloading to complete")
	})
}

func doOffloadStatus(vc *cmdutils.VerbCmd, wait bool) error {
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
	status, err := admin.Topics().OffloadStatus(*topic)
	if err != nil {
		return err
	}

	for wait && status.Status == RUNNING {
		time.Sleep(1 * time.Second)
		status, err = admin.Topics().OffloadStatus(*topic)
		if err != nil {
			return err
		}
	}

	switch status.Status {
	case NOT_RUN:
		vc.Command.Printf("Offload has not been run for %s since broker startup/n", topic.String())
	case RUNNING:
		vc.Command.Printf("Offload is currently running/n")
	case SUCCESS:
		vc.Command.Printf("Offload was a success/n")
	case ERROR:
		vc.Command.Printf("Error in offload")
		err = errors.New(status.LastError)
	}

	return err
}
