package cluster

import (
	"encoding/json"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
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

	_, _, _, err := TestClusterCommands(updateClusterCmd, args)
	if err != nil {
		t.Error(err)
	}

	args = []string{"get", "standalone"}
	out, _, _, err := TestClusterCommands(getClusterDataCmd, args)

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
