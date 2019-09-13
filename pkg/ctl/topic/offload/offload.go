package offload

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/args"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func OffloadCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for triggering offload the data from a topic " +
		"to long-term storage (e.g. Amazon S3)"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	offload := Example{
		Desc:    "Trigger offload the data from a topic <topic-name> to long-term storage and " +
			"keep <threshold> size data in BookKeeper (e.g. 10M, 5G, default is byte)",
		Command: "pulsarctl topic offload <topic-name> <threshold>",
	}
	desc.CommandExamples = append(examples, offload)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "Offload trigger for <topic-name> for messages before <message-id>",
	}

	nothingOut := Output{
		Desc: "noting to offload",
		Out: "Nothing to offload",
	}

	argsError := Output{
		Desc: "the topic name is not specified or the offload threshold is not specified",
		Out: "[✖]  only two argument is allowed to be used as names",
	}
	out = append(out, successOut, nothingOut, argsError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"offload",
		"Offload the data form a topic to long-term storage",
		desc.ToString())

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doOffload(vc)
	}, CheckTopicNameTwoArgs)

}

func doOffload(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	size, err := validateSize(vc.NameArgs[1])
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()

	stats, err := admin.Topics().GetInternalStats(*topic)
	if err != nil {
		return err
	}

	if len(stats.Ledgers) < 1 {
		return errors.New("Topic doesn't have any data.")
	}

	ledgers := stats.Ledgers
	ledgers[len(ledgers)-1].Size = stats.CurrentLedgerSize
	messageId := findFirstLedgerWithinThreshold(ledgers, size)
	if err == nil {
		vc.Command.Printf("Nothing to offload")
		return nil
	}

	err = admin.Topics().Offload(*topic, *messageId)
	if err == nil {
		vc.Command.Printf("Offload trigger for %s for messages before %v", topic.String(), messageId)
	}

	return err
}

func findFirstLedgerWithinThreshold(ledgers []LedgerInfo, sizeThreshold int64) *MessageId {
	var suffixSize int64
	previousLedger := ledgers[len(ledgers)-1].LedgerId
	for i := len(ledgers) - 1; i >= 0; i-- {
		suffixSize += ledgers[i].Size
		if suffixSize > sizeThreshold {
			return &MessageId{
				LedgerId: previousLedger,
				EntryId: 0,
				PartitionIndex: -1,
			}
		}
		previousLedger = ledgers[i].LedgerId
	}
	return nil
}

func validateSize(s string) (int64, error) {
	end := s[len(s)-1:]
	value := s[:len(s)-1]
	switch end {
	case "k":
		fallthrough
	case "K":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024, err
	case "m":
		fallthrough
	case "M":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024, err
	case "g":
		fallthrough
	case "G":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024 * 1024, err
	case "t":
		fallthrough
	case "T":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024 * 1024 * 1024, err
	default:
		return strconv.ParseInt(s, 10, 64)
	}
}
