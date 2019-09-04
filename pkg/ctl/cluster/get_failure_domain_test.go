package cluster

import (
	"encoding/json"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFailureDomainSuccess(t *testing.T) {
	args := []string{"create", "failure-broker-A"}
	_, _, _, err := TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "failure-broker-B"}
	_, _, _, err = TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "--domain-name", "failure-domain", "--brokers", "failure-broker-A", "--brokers", "failure-broker-B", "standalone"}
	_, _, _, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"get-failure-domain", "--domain-name", "failure-domain", "standalone"}
	out, _, _, err := TestClusterCommands(getFailureDomainCmd, args)
	assert.Nil(t, err)

	var brokers pulsar.FailureDomainData
	err = json.Unmarshal(out.Bytes(), &brokers)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "failure-broker-A", brokers.BrokerList[0])
	assert.Equal(t, "failure-broker-B", brokers.BrokerList[1])
}

func TestGetFailureDomainArgsError(t *testing.T) {
	args := []string{"get-failure-domain", "standalone"}
	_, _, nameErr, _ := TestClusterCommands(getFailureDomainCmd, args)
	assert.Equal(t, "need to specified two names for cluster and failure domain", nameErr.Error())
}

func TestGetNonExistFailureDomain(t *testing.T) {
	args := []string{"get-failure-domain", "standalone", "non-exist"}
	_, execErr, _, _ := TestClusterCommands(getFailureDomainCmd, args)
	assert.NotNil(t, execErr)
}
