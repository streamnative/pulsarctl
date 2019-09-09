package lookup

import (
	"encoding/json"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestLookupTopicCmd(t *testing.T) {
	args := []string{"create", "test-lookup-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"lookup", "test-lookup-topic"}
	out, execErr, _, _  := TestTopicCommands(LookupTopicCmd, args)
	assert.Nil(t, execErr)

	var data pulsar.LookupData
	err := json.Unmarshal(out.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}

	brokerUrl, err := regexp.Compile("^pulsar://[a-z-A-Z0-9]*:6650$")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, brokerUrl.MatchString(data.BrokerUrl))

	httpUrl, err :=  regexp.Compile("^http://[a-z-A-Z0-9]*:8080$")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, httpUrl.MatchString(data.HttpUrl))
}

func TestLookupTopicArgError(t *testing.T)  {
	args  := []string{"lookup"}
	_, _, nameErr, _ := TestTopicCommands(LookupTopicCmd,  args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}
