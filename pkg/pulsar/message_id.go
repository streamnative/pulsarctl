package pulsar

type MessageID struct {
	LedgerID         int64 `json:"ledgerId"`
	EntryID          int64 `json:"entryId"`
	PartitionedIndex int   `json:"partitionedIndex"`
}
