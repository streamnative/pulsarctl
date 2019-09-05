package cluster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteFailureDomainCmd(t *testing.T) {
	args := []string{"create", "delete-failure-test"}
	_, _, _, err := TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "cluster-A"}
	_, _, _, err = TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "-b", "cluster-A", "delete-failure-test", "delete-failure-domain"}
	_, _, _, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"delete-failure-domain", "delete-failure-test", "delete-failure-domain"}
	_, _, _, err = TestClusterCommands(deleteFailureDomainCmd, args)
	assert.Nil(t, err)
}

func TestDeleteFailureDomainArgsError(t *testing.T) {
	args := []string{"delete-failure-domain", "standalone"}
	_, _, nameErr, _ := TestClusterCommands(deleteFailureDomainCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the cluster name and the failure domain name", nameErr.Error())
}

// delete a non-existent failure domain in an existing cluster
func TestDeleteNonExistentFailureDomain(t *testing.T) {
	args := []string{"delete-failure-domain", "standalone", "non-existent-failure-domain"}
	_, execErr, _, _ := TestClusterCommands(deleteFailureDomainCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Domain-name non-existent-failure-domain"+
		" or cluster standalone does not exist", execErr.Error())
}

// delete a non-existent failure domain in a non-existent cluster
func TestDeleteNonExistentFailureDomainInNonExistentCluster(t *testing.T) {
	args := []string{"delete-failure-domain", "non-existent-cluster", "non-existent-failure-domain"}
	_, execErr, _, _ := TestClusterCommands(deleteFailureDomainCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 412 reason: Cluster non-existent-cluster does not exist.", execErr.Error())
}
