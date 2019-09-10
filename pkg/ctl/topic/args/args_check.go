package args

import (
	"github.com/pkg/errors"
)

func CheckTopicNameTwoArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("need to specified the topic name and the partitions")
	}

	return nil
}
