package args

import (
	"github.com/pkg/errors"
)

func CheckTopicNameTwoArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("only two argument is allowed to be used as names")
	}

	return nil
}
