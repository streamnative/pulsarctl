package cluster

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDeleteClusterCmd(t *testing.T) {
	args := []string{"add", "delete-test"}
	_, _, _, err := TestClusterCommands(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, _, _, err := TestClusterCommands(listClustersCmd, args)
	assert.Nil(t, err)
	clusters := out.String()
	assert.True(t, strings.Contains(clusters, "delete-test"))

	args = []string{"delete", "delete-test"}
	_, _, _, err = TestClusterCommands(deleteClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, _, _, err = TestClusterCommands(listClustersCmd, args)
	assert.Nil(t, err)
	clusters = out.String()
	assert.False(t, strings.Contains(clusters, "delete-test"))
}
