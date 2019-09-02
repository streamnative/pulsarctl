package cluster

import (
	"encoding/json"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdatePeerClusters(t *testing.T) {
	args := []string{"add", "test_peer_cluster"}
	_, err := TestClusterCommands(createClusterCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	args = []string{"update-peer-clusters", "-p", "test_peer_cluster", "standalone"}
	_, err = TestClusterCommands(updatePeerClustersCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	args = []string{"get", "standalone"}
	out, err := TestClusterCommands(getClusterDataCmd, args)

	var clusterData pulsar.ClusterData
	err = json.Unmarshal(out.Bytes(), &clusterData)
	if err != nil {
		t.Fatal(err)
	}

	peer_cluster := clusterData.PeerClusterNames[0]
	assert.Equal(t, "test_peer_cluster", peer_cluster)
}
