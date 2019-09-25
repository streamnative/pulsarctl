package cluster

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestGetClusterData(t *testing.T) {
	args := []string{"get", "standalone"}
	out, _, _, err := TestClusterCommands(getClusterDataCmd, args)
	if err != nil {
		t.Error(err)
	}
	c := pulsar.ClusterData{}
	err = json.Unmarshal(out.Bytes(), &c)
	if err != nil {
		t.Error(err)
	}

	pulsarURL := regexp.MustCompile("^pulsar://[a-z-A-Z0-9]*:6650$")
	res := pulsarURL.MatchString(c.BrokerServiceURL)
	assert.True(t, res)

	httpURL := regexp.MustCompile("^http://[a-z-A-Z0-9]*:8080$")
	res = httpURL.MatchString(c.ServiceURL)
	assert.True(t, res)
}
