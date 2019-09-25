package cluster

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestUpdatePeerClusters(t *testing.T) {
	args := []string{"add", "test_peer_cluster"}
	_, _, _, err := TestClusterCommands(CreateClusterCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	args = []string{"update-peer-clusters", "-p", "test_peer_cluster", "standalone"}
	_, _, _, err = TestClusterCommands(updatePeerClustersCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	args = []string{"get", "standalone"}
	out, _, _, _ := TestClusterCommands(getClusterDataCmd, args)

	var clusterData pulsar.ClusterData
	err = json.Unmarshal(out.Bytes(), &clusterData)
	if err != nil {
		t.Fatal(err)
	}

	peerCluster := clusterData.PeerClusterNames[0]
	assert.Equal(t, "test_peer_cluster", peerCluster)
}
