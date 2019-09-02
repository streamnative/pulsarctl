package cluster

import (
	"encoding/json"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListFailureDomainsCmd(t *testing.T)  {
	args := []string{"create", "list-failure-test"}
	_, err := TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "list-failure-broker-A"}
	_, err = TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "list-failure-broker-B"}
	_, err = TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "--domain-name", "list-failure-A", "--brokers", "list-failure-broker-A", "list-failure-test"}
	_, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "--domain-name", "list-failure-B", "--brokers", "list-failure-broker-B", "list-failure-test"}
	_, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"list-failure-domains", "list-failure-test"}
	out, err := TestClusterCommands(listFailureDomainCmd, args)
	assert.Nil(t, err)

	var brokerMap pulsar.FailureDomainMap
	err = json.Unmarshal(out.Bytes(), &brokerMap)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, brokerMap["list-failure-A"])
	assert.Equal(t, "list-failure-broker-A", brokerMap["list-failure-A"].BrokerList[0])
	assert.NotNil(t, brokerMap["list-failure-B"])
	assert.Equal(t, "list-failure-broker-B", brokerMap["list-failure-B"].BrokerList[0])
}
