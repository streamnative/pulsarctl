// +build tls

package cluster

import (
    `github.com/stretchr/testify/assert`
    `strings`
    `testing`
)

func TestTLS(t *testing.T) {
    args := []string{"clusters","add", "tls"}
    _, err := TestTlsHelp(createClusterCmd, args)
    assert.Nil(t, err)

    args = []string{"clusters","list"}
    out, err := TestTlsHelp(listClustersCmd, args)
    assert.Nil(t, err)
    clusters := out.String()
    assert.True(t, strings.Contains(clusters, "tls"))

    args = []string{"clusters","delete", "tls"}
    _, err = TestTlsHelp(deleteClusterCmd, args)
    assert.Nil(t, err)

    args = []string{"clusters","list"}
    out, err = TestTlsHelp(listClustersCmd, args)
    assert.Nil(t, err)
    clusters = out.String()
    assert.False(t, strings.Contains(clusters, "tls"))
}

