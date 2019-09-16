package pulsar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMessageId(t *testing.T) {
	id, err := ParseMessageId("1;1")
	assert.Nil(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid message id string. 1;1", err.Error())

	id, err = ParseMessageId("a:1")
	assert.Nil(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid ledger id string. a:1", err.Error())

	id, err = ParseMessageId("1:a")
	assert.Nil(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid entry id string. 1:a", err.Error())
}
