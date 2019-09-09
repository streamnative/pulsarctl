package stats

import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetPartitionedStatsCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the stats of the partitioned topic " +
		"and its producers and consumers. (All the rate are computed over a 1 minute window " +
		"and are relative the last completed 1 minute period)"
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get the partitioned topic <topic-name> stats",
		Command: "pulsarctl topic partitioned-stats <topic-name>",
	}

	getPerPartition := Example{
		Desc:    "Get the partitioned topic <topic-name> stats and per partition stats",
		Command: "pulsarctl topic partitioned-stats --per-partition <topic-name>",
	}
	desc.CommandExamples = append(examples, get, getPerPartition)

	var out []Output
	successOut := Output{
		Desc: "Get the partitioned topic stats",
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
  "deduplicationStatus": "",
  "metadata": {
    "partitions": 1
  },
  "partitions": {}
}`,
	}
	out = append(out, successOut)

	perPartitionOut := Output{
		Desc: "Get the partitioned topic stats and per partition topic stats",
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
  "deduplicationStatus": "",
  "metadata": {
    "partitions": 1
  },
  "partitions": {
    "<topic-name>": {
      "msgRateIn": 0,
      "msgRateOut": 0,
      "msgThroughputIn": 0,
      "msgThroughputOut": 0,
      "averageMsgSize": 0,
      "storageSize": 0,
      "publishers": [],
      "subscriptions": {},
      "replication": {},
      "deduplicationStatus": ""
    }
  }
}`,
	}

	out = append(out, perPartitionOut, ArgError)

	topicNotExist := Output{
		Desc: "the specified topic is not exist or the specified topic is not a partitioned topic",
		Out:  "[✖]  code: 404 reason: Partitioned Topic not found",
	}
	out = append(out, topicNotExist)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"partitioned-stats",
		"Get the partitioned topic stats",
		desc.ToString(),
		"ps")

	var perPartition bool

	vc.FlagSetGroup.InFlagSet("PartitionedStats", func(set *pflag.FlagSet) {
		set.BoolVarP(&perPartition, "per-partition", "p", false,
			"Get the per partition topic stats")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doGetPartitionedStats(vc, perPartition)
	})
}

func doGetPartitionedStats(vc *cmdutils.VerbCmd, perPartition bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	stats, err := admin.Topics().GetPartitionedStats(*topic, perPartition)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), stats)
	}

	return err
}
