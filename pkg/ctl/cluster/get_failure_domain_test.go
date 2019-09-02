package cluster

import (
	"encoding/json"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFailureDomainCmd(t *testing.T)  {
	args := []string{"create", "failure-broker-A"}
	_, err := TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "failure-broker-B"}
	_, err = TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "--domain-name", "failure-domain", "--brokers", "failure-broker-A", "--brokers", "failure-broker-B", "standalone"}
	_, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"get-failure-domain", "--domain-name", "failure-domain", "standalone",}
	out, err := TestClusterCommands(getFailureDomainCmd, args)
	assert.Nil(t, err)

	var brokers pulsar.FailureDomainData
	err = json.Unmarshal(out.Bytes(), &brokers)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "failure-broker-A", brokers.BrokerList[0])
	assert.Equal(t, "failure-broker-B", brokers.BrokerList[1])
}