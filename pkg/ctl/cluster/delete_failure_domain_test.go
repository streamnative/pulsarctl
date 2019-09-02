package cluster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteFailureDomainCmd(t *testing.T) {
	args := []string{"create", "delete-failure-test"}
	_, err := TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "delete-failure-domain-A"}
	_, err = TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "--domain-name", "delete-failure-domain-A", "delete-failure-test"}
	_, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"delete-failure-domain", "--domain-name", "delete-failure-domain-A", "delete-failure-test"}
	_, err = TestClusterCommands(deleteFailureDomainCmd, args)
	assert.Nil(t, err)
}
