package stats

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetStatsCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the stats for an existing non-partitioned topic and its " +
		"connected producers and consumers. (All the rates are computed over a 1 minute window " +
		"and are relative the last completed 1 minute period)"
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get the stats of the specified topic <topic-name>",
		Command: "pulsarctl topic stats <topic-name>",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: `{
  "msgRateIn": 0,
  "msgRateOut": 0,
  "msgThroughputIn": 0,
  "msgThroughputOut": 0,
  "averageMsgSize": 0,
  "storageSize": 0,
  "publishers": [],
  "subscriptions": {},
  "replication": {},
  "deduplicationStatus": "Disabled"
}`,
	}
	out = append(out, successOut, ArgError)

	topicNotFoundError := Output{
		Desc: "the specified topic is not exist or the specified topic is a partitioned-topic",
		Out:  "code: 404 reason: Topic not found",
	}
	out = append(out, topicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"stats",
		"Get the stats of an existing non-partitioned topic",
		desc.ToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetStats(vc)
	})
}

func doGetStats(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	topicStats, err := admin.Topics().GetStats(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), topicStats)
	}

	return err
}
