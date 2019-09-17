package namespace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxConsumersPerSubscriptionCmd(t *testing.T) {
	ns := "public/test-max-consumers-per-subscription"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-max-consumers-per-subscription", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetMaxConsumersPerSubscriptionCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The max consumers per subscription of namespace %s is %d", ns, 0),
		out.String())

	args = []string{"set-max-consumers-per-subscription", "--size", "10", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetMaxConsumersPerSubscriptionCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the max consumers per subscription of namespace %s to %d", ns, 10),
		out.String())

	args = []string{"get-max-consumers-per-subscription", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetMaxConsumersPerSubscriptionCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The max consumers per subscription of namespace %s is %d", ns, 10),
		out.String())
}

func TestSetMaxConsumersPerSubscriptionOnNonExistingTopic(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-max-consumers-per-subscription", "--size", "10", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetMaxConsumersPerSubscriptionCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestSetMaxConsumersPerSubscriptionWithInvalidSize(t *testing.T)  {
	args := []string{"set-max-consumers-per-subscription", "--size", "-1", "public/invalid-size"}
	_, execErr, _, _ := TestNamespaceCommands(SetMaxConsumersPerSubscriptionCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "the specified consumers value must bigger than 0", execErr.Error())
}

func TestGetMaxConsumersPerSubscriptionOnNonExistingTopic(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-max-consumers-per-subscription", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetMaxConsumersPerSubscriptionCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
