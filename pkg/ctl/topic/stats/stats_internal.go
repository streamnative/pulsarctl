package stats

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetInternalStatsCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the internal stats for an existing non-partitioned topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get internal stats for an existing non-partitioned-topic <topic-name>",
		Command: "pulsarctl topic internal-stats <topic-name>",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: `{
  "entriesAddedCounter": 0,
  "numberOfEntries": 0,
  "totalSize": 0,
  "currentLedgerEntries": 0,
  "currentLedgerSize": 0,
  "lastLedgerCreatedTimestamp": "",
  "lastLedgerCreationFailureTimestamp": "",
  "waitingCursorsCount": 0,
  "pendingAddEntriesCount": 0,
  "lastConfirmedEntry": "",
  "state": "",
  "ledgers": [
    {
      "ledgerId": 0,
      "entries": 0,
      "size": 0,
      "offloaded": false
    }
  ],
  "cursors": {}
}`,
	}
	out = append(out, successOut, ArgError)

	partitionedTopicInternalStatsError := Output{
		Desc: "the specified topic is not exist or the specified topic is a partitioned topic",
		Out:  "[✖]  code: 404 reason: Topic not found",
	}
	out = append(out, partitionedTopicInternalStatsError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"internal-stats",
		"Get the internal stats of the specified topic",
		desc.ToString(),
		"")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetInternalStats(vc)
	})
}

func doGetInternalStats(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	stats, err := admin.Topics().GetInternalStats(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), stats)
	}

	return err
}
