package cluster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFailureDomainCmdSuccess(t *testing.T) {
	args := []string{"create-failure-domain", "-b", "cluster-A", "standalone", "standalone-failure-domain"}
	_, execErr, NameErr, err := TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, NameErr)
	assert.Nil(t, err)
}

func TestCreateFailureDomainCmdBrokerListError(t *testing.T) {
	args := []string{"create-failure-domain", "standalone", "standalone-failure-domain"}
	_, execErr, _, _ := TestClusterCommands(createFailureDomainCmd, args)
	assert.Equal(t, "broker list must be specified", execErr.Error())
}
