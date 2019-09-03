package cluster

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetPeerClustersCmd(t *testing.T) {
	args := []string{"add", "test_get_peer", "--peer-cluster", "standalone"}
	_, err := TestClusterCommands(createClusterCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	args = []string{"gpc", "test_get_peer"}
	out, err := TestClusterCommands(getPeerClustersCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	res := out.String()
	assert.True(t, strings.Contains(res, "standalone"))
}

