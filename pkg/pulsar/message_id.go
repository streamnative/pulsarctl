package pulsar

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type MessageId struct {
	LedgerId       int64 `json:"ledgerId"`
	EntryId        int64 `json:"entryId"`
	PartitionIndex int   `json:"partitionIndex"`
}

var Latest = MessageId{0x7fffffffffffffff, 0x7fffffffffffffff, -1}
var Earliest = MessageId{-1, -1, -1}

func ParseMessageId(str string) (*MessageId, error) {
	s := strings.Split(str, ":")
	if len(s) != 2 {
		return nil, errors.Errorf("Invalid message id string. %s", str)
	}

	ledgerId, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return nil, errors.Errorf("Invalid ledger id string. %s", str)
	}

	entryId, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return nil, errors.Errorf("Invalid entry id string. %s", str)
	}

	return &MessageId{LedgerId: ledgerId, EntryId: entryId, PartitionIndex: -1}, nil
}
