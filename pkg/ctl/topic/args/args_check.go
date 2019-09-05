package args

import (
	"github.com/pkg/errors"
)

<<<<<<< HEAD
func CheckTopicNameTwoArgs(args []string) error {
=======
func CheckArgs(args []string) error {
>>>>>>> Add partitioned topic command CURD
	if len(args) != 2 {
		return errors.New("need to specified the topic name and the partitions")
	}

	return nil
}
