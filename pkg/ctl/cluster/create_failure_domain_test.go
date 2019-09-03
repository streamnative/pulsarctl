package cluster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFailureDomainCmdSuccess(t *testing.T) {
	args := []string{"create-failure-domain", "standalone", "standalone-failure-domain"}
	_, err := TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)
}
