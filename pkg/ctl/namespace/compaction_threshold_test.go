package namespace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompactionThresholdCmd(t *testing.T) {
	ns := "public/test-compaction-threshold"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-compaction-threshold", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetCompactionThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The compaction threshold is %d byte(s) for the namespace %s", 0, ns),
		out.String())

	args = []string{"set-compaction-threshold", "--size", "10m", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetCompactionThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Set the compaction threshold %d for the namespace %s", 10 * 1024 *1024, ns),
		out.String())

	args = []string{"get-compaction-threshold", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetCompactionThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The compaction threshold is %d byte(s) for the namespace %s", 10 * 1024 *1024, ns),
		out.String())
}

func TestSetCompactionThresholdOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-compaction-threshold", "--size", "10m", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetCompactionThresholdCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetCompactionThresholdOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-compaction-threshold", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetCompactionThresholdCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
