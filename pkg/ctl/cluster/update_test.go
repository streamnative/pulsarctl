package cluster

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCluster(t *testing.T) {
	args := []string{
		"update",
		"--url", "http://example:8080",
		"--url-tls", "https://example:8080",
		"--broker-url", "pulsar://example:6650",
		"--broker-url-tls", "pulsar+ssl://example:6650",
		"-p", "cluster-a",
		"-p", "cluster-b",
		"standalone",
	}

	_, _, _, err := TestClusterCommands(UpdateClusterCmd, args)
	if err != nil {
		t.Error(err)
	}

	args = []string{"get", "standalone"}
	out, execErr, _, _ := TestClusterCommands(getClusterDataCmd, args)
	assert.Nil(t, execErr)

	var data pulsar.ClusterData
	err = json.Unmarshal(out.Bytes(), &data)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "http://example:8080", data.ServiceURL)
	assert.Equal(t, "https://example:8080", data.ServiceURLTls)
	assert.Equal(t, "pulsar://example:6650", data.BrokerServiceURL)
	assert.Equal(t, "pulsar+ssl://example:6650", data.BrokerServiceURLTls)
}
