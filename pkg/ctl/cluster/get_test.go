package cluster

import (
	"encoding/json"
	"fmt"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)


func TestGetClusterConfiguration(t *testing.T) {
	args := []string{"get","--cluster-name", "standalone"}
	out, err := TestCommands(getClusterConfiguration, args)
	if err != nil {
		t.Error(err)
	}
	c := pulsar.ClusterData{}
	err = json.Unmarshal(out.Bytes(), &c)
	if err != nil {
		fmt.Println(err)
	}

	pulsarUrl, err := regexp.Compile("^pulsar://[a-z-A-Z]*:6650$")
	if err != nil {
		t.Error(err)
	}

	res := pulsarUrl.MatchString(c.BrokerServiceURL)
	assert.True(t, res)

	httpUrl, err := regexp.Compile("^http://[a-z-A-Z]*:8080$")
	if err != nil {
		t.Error(err)
	}

	res = httpUrl.MatchString(c.ServiceURL)
	assert.True(t, res)
}

