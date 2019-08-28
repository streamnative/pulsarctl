package pulsar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNil(t *testing.T) {
	var a interface{}
	var b interface{} = (*int)(nil)

	assert.True(t, a == nil)
	assert.False(t, b == nil)
}
