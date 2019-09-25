package pulsar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiVersion_String(t *testing.T) {
	assert.Equal(t, "", V1.String())
	assert.Equal(t, "v2", V2.String())
	assert.Equal(t, "v3", V3.String())
}
