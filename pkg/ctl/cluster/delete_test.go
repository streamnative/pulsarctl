package cluster

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDeleteClusterCmd(t *testing.T) {
	args := []string{"add", "test"}
	_, _, _, err := TestClusterCommands(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, _, _, err := TestClusterCommands(listClustersCmd, args)
	assert.Nil(t, err)
	clusters := out.String()
	assert.True(t, strings.Contains(clusters, "test"))

	args = []string{"delete", "test"}
	_, _, _, err = TestClusterCommands(deleteClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, _, _, err = TestClusterCommands(listClustersCmd, args)
	assert.Nil(t, err)
	clusters = out.String()
	assert.False(t, strings.Contains(clusters, "test"))
}
