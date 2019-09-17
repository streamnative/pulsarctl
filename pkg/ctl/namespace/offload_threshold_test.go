package namespace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOffloadThresholdCmd(t *testing.T) {
	ns := "public/test-offload-threshold"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-offload-threshold", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetOffloadThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload threshold is %d byte(s) for the namespace %s", -1, ns),
		out.String())

	args = []string{"set-offload-threshold", "--size", "10m", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetOffloadThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Set the offload threshold %d for the namespace %s", 10 * 1024 *1024, ns),
		out.String())

	args = []string{"get-offload-threshold", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetOffloadThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload threshold is %d byte(s) for the namespace %s", 10 * 1024 *1024, ns),
		out.String())
}

func TestSetOffloadThresholdOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-offload-threshold", "--size", "10m", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetOffloadThresholdCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetOffloadThresholdOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-offload-threshold", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetOffloadThresholdCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
