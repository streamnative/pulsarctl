package cluster

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetPeerClustersCmd(t *testing.T) {
	args := []string{"add", "test_get_peer", "--peer-cluster", "standalone"}
	_, _, _, err := TestClusterCommands(CreateClusterCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	args = []string{"gpc", "test_get_peer"}
	out, _, _, err := TestClusterCommands(getPeerClustersCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	res := out.String()
	assert.True(t, strings.Contains(res, "standalone"))
}
