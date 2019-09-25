package lookup

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestLookupTopicCmd(t *testing.T) {
	args := []string{"create", "test-lookup-topic", "0"}
	_, execErr, _, _ := test.TestTopicCommands(crud.CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"lookup", "test-lookup-topic"}
	out, execErr, _, _ := test.TestTopicCommands(TopicCmd, args)
	assert.Nil(t, execErr)

	var data pulsar.LookupData
	err := json.Unmarshal(out.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}

	brokerURL := regexp.MustCompile("^pulsar://[a-z-A-Z0-9]*:6650$")
	assert.True(t, brokerURL.MatchString(data.BrokerURL))

	httpURL := regexp.MustCompile("^http://[a-z-A-Z0-9]*:8080$")
	assert.True(t, httpURL.MatchString(data.HTTPURL))
}

func TestLookupTopicArgError(t *testing.T) {
	args := []string{"lookup"}
	_, _, nameErr, _ := test.TestTopicCommands(TopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}
