package cluster

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDeleteClusterCmd(t *testing.T)  {
	args := []string{"add", "test"}
	_, err := TestClusterCommand(createClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, err := TestClusterCommand(listClustersCmd, args)
	assert.Nil(t, err)
	clusters := out.String()
	assert.True(t, strings.Contains(clusters, "test"))

	args = []string{"delete", "test"}
	_, err = TestClusterCommand(deleteClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, err = TestClusterCommand(listClustersCmd, args)
	assert.Nil(t, err)
	clusters = out.String()
	assert.False(t, strings.Contains(clusters, "test"))
}
