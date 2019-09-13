package pulsar

type Status string

const (
	NOT_RUN Status = "NOT_RUN"
	RUNNING Status = "RUNNING"
	SUCCESS Status = "SUCCESS"
	ERROR Status = "ERROR"
)

type LongRunningProcessStatus struct {
	Status Status `json:"status"`
	LastError string `json:"lastError"`
}

type OffloadProcessStatus struct {
	Status Status `json:"status"`
	LastError string `json:"lastError"`
	FirstUnoffloadedMessage MessageId `json:"firstUnoffloadedMessage"`
}

func (s Status) String() string {
	return string(s)
}
