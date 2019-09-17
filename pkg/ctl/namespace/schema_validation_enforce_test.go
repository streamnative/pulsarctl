package namespace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaValidationEnforcedCmd(t *testing.T) {
	ns := "public/test-schema-validation-enforced"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-schema-validation-enforced", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetSchemaValidationEnforcedCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Schema validation enforced is disabled\n", out.String())

	args = []string{"set-schema-validation-enforced", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSchemaValidationEnforcedCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Enable schema validation enforced\n", out.String())

	args = []string{"get-schema-validation-enforced", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSchemaValidationEnforcedCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Schema validation enforced is enabled\n", out.String())
}

func TestSetSchemaValidationEnforcedOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-schema-validation-enforced", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetSchemaValidationEnforcedCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetSchemaValidationEnforcedOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-schema-validation-enforced", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetSchemaValidationEnforcedCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
