// +build tls

package cluster

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTLS(t *testing.T) {
	args := []string{"clusters", "add", "tls"}
	_, err := TestTLSHelp(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"clusters", "list"}
	out, err := TestTLSHelp(listClustersCmd, args)
	assert.Nil(t, err)
	clusters := out.String()
	assert.True(t, strings.Contains(clusters, "tls"))

	args = []string{"clusters", "delete", "tls"}
	_, err = TestTLSHelp(deleteClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"clusters", "list"}
	out, err = TestTLSHelp(listClustersCmd, args)
	assert.Nil(t, err)
	clusters = out.String()
	assert.False(t, strings.Contains(clusters, "tls"))
}
