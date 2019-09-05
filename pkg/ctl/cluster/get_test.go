package cluster

import (
	"encoding/json"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)


func TestGetClusterData(t *testing.T) {
	args := []string{"get", "standalone"}
	out, _, _, err := TestClusterCommands(getClusterDataCmd, args)
	if err != nil {
		t.Error(err)
	}
	c := ClusterData{}
	err = json.Unmarshal(out.Bytes(), &c)
	if err != nil {
		t.Error(err)
	}

	pulsarUrl, err := regexp.Compile("^pulsar://[a-z-A-Z0-9]*:6650$")
	if err != nil {
		t.Error(err)
	}

	res := pulsarUrl.MatchString(c.BrokerServiceURL)
	assert.True(t, res)

	httpUrl, err := regexp.Compile("^http://[a-z-A-Z0-9]*:8080$")
	if err != nil {
		t.Error(err)
	}

	res = httpUrl.MatchString(c.ServiceURL)
	assert.True(t, res)
}

