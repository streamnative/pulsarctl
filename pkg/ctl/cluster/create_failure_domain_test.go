package cluster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFailureDomainCmdSuccess(t *testing.T) {
	args := []string{"cfd", "--domain-name", "test-domain", "standalone"}
	_, err := TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)
}
