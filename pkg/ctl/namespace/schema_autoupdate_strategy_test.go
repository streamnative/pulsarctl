package namespace

import (
	"fmt"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestSchemaAutoUpdateStrategyCmd(t *testing.T) {
	ns := "public/test-schema-autoupdate-strategy"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-schema-autoupdate-strategy", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The schema auto-update strategy is %s for the namespace %s", pulsar.Full.String(), ns),
		out.String())

	args = []string{"set-schema-autoupdate-strategy", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Set the schema auto-update strategy %s for the namespace %s", pulsar.AutoUpdateDisabled.String(), ns),
		out.String())

	args = []string{"get-schema-autoupdate-strategy", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The schema auto-update strategy is %s for the namespace %s", pulsar.AutoUpdateDisabled.String(), ns),
		out.String())

	args = []string{"set-schema-autoupdate-strategy", "--compatibility", "BackwardTransitive", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Set the schema auto-update strategy %s for the namespace %s", pulsar.BackwardTransitive.String(), ns),
		out.String())

	args = []string{"get-schema-autoupdate-strategy", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The schema auto-update strategy is %s for the namespace %s", pulsar.BackwardTransitive.String(), ns),
		out.String())
}

func TestSetSchemaAutoUpdateStrategyOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-schema-autoupdate-strategy", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetSchemaAutoUpdateStrategyCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetSchemaAutoUpdateStrategyOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-schema-autoupdate-strategy", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetSchemaAutoUpdateStrategyCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
