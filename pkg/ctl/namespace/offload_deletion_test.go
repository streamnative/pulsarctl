package namespace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOffloadDeletionLagCmd(t *testing.T) {
	ns := "public/test-offload-deletion-lag"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-offload-deletion-lag", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload deletion lag is %f minute(s) for the namespace %s", 0.000000, ns),
		out.String())

	args = []string{"set-offload-deletion-lag", "--lag", "10m", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Set offload deletion lag %s for the namespace %s", "10m", ns),
		out.String())

	args = []string{"get-offload-deletion-lag", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload deletion lag is %f minute(s) for the namespace %s", 10.000000, ns),
		out.String())

	args = []string{"clear-offload-deletion-lag", ns}
	out, execErr, _, _ = TestNamespaceCommands(ClearOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Clear offload deletion lag for the namespace %s successfully", ns),
		out.String())

	args = []string{"get-offload-deletion-lag", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload deletion lag is not set for the namespace %s", ns),
		out.String())

}

func TestOffloadDeletionLagOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-offload-deletion-lag", "--lag", "10m", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetOffloadDeletionLagCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetOfloadThresholdOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-offload-deletion-lag", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetOffloadDeletionLagCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestClearOffloadDeletionLagOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"clear-offload-deletion-lag", ns}
	_, execErr, _, _ := TestNamespaceCommands(ClearOffloadDeletionLagCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
