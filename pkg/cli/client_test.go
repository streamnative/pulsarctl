package cli

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeJSONBody(t *testing.T) {
	testcases := []struct {
		obj      interface{}
		expected int
	}{
		{obj: "1", expected: 3},
		{obj: "12", expected: 4},
		{obj: 1, expected: 1},
		{obj: 12, expected: 2},
	}

	for _, testcase := range testcases {
		r, err := encodeJSONBody(testcase.obj)
		require.NoError(t, err)

		b, err := ioutil.ReadAll(r)
		require.NoError(t, err)

		require.Equal(t, testcase.expected, len(b))
	}
}
