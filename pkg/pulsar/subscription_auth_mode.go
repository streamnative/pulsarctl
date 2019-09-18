package pulsar

import "github.com/pkg/errors"

type SubscriptionAuthMode string

const (
	None   SubscriptionAuthMode = "None"
	Prefix SubscriptionAuthMode = "Prefix"
)

func ParseSubscriptionAuthMode(s string) (SubscriptionAuthMode, error) {
	switch s {
	case "None":
		return None, nil
	case "Prefix":
		return Prefix, nil
	default:
		return "", errors.New("Invalid subscription auth mode")
	}
}

func (s SubscriptionAuthMode) String() string {
	return string(s)
}
