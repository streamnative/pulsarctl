package cluster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteFailureDomainCmd(t *testing.T) {
	args := []string{"create", "delete-failure-test"}
	_, _, _, err := TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "delete-failure-domain-A"}
	_, _, _, err = TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain","-b", "delete-failure-domain-A", "delete-failure-test", "delete-failure-domain"}
	_, _, _, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"delete-failure-domain", "delete-failure-test", "delete-failure-domain"}
	_, _, _, err = TestClusterCommands(deleteFailureDomainCmd, args)
	assert.Nil(t, err)
}

func TestDeleteFilureDomainArgsError(t *testing.T)  {
	args := []string{"delete-failure-domain", "standalone"}
	_, _, nameErr, _ := TestClusterCommands(deleteFailureDomainCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the cluster name and the failure domain name", nameErr.Error())
}
