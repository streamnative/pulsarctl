package args

import (
	"github.com/pkg/errors"
)

func CheckSubscriptionNameTwoArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("need to specified the topic name and the subscription name")
	}

	return nil
}