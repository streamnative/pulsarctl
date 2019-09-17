package namespace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProducersPerTopicCmd(t *testing.T) {
	ns := "public/test-max-producers-per-topic"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-max-producers-per-topic", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetMaxProducersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The max producers per topic of namespace %s is %d", ns, 0),
		out.String())

	args = []string{"set-max-producers-per-topic", "--size", "10", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetMaxProducersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the max producers per topic of namespace %s to %d", ns, 10),
		out.String())

	args = []string{"get-max-producers-per-topic", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetMaxProducersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The max producers per topic of namespace %s is %d", ns, 10),
		out.String())
}

func TestSetMaxProducersPerTopicOnNonExistingTopic(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-max-producers-per-topic", "--size", "10", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetMaxProducersPerTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestSetMaxProducersPerTopicWithInvalidSize(t *testing.T)  {
	args := []string{"set-max-producers-per-topic", "--size", "-1", "public/invalid-size"}
	_, execErr, _, _ := TestNamespaceCommands(SetMaxProducersPerTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "the specified producers value must bigger than 0", execErr.Error())
}

func TestGetMaxProducersPerTopicOnNonExistingTopic(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-max-producers-per-topic", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetMaxProducersPerTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
