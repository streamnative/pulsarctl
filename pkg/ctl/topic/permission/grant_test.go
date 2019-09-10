package permission

import (
	"encoding/json"
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestGrantPermissionToNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-grant-permission-non-partitioned-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", "test-grant-permission-non-partitioned-topic"}
	out, execErr, _, _ := TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	var permissions map[string][]pulsar.AuthAction
	err := json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, map[string][]pulsar.AuthAction{}, permissions)

	args = []string{"grant-permissions",
		"--role", "grant-non-partitioned-role",
		"--actions", "produce",
		"test-grant-permission-non-partitioned-topic",
	}
	_, execErr, _, _ = TestTopicCommands(GrantPermissionCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", "test-grant-permission-non-partitioned-topic"}
	out, execErr, _, _ = TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	err = json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []pulsar.AuthAction{"produce"}, permissions["grant-non-partitioned-role"])
}

func TestGrantPermissionToPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-grant-permission-partitioned-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", "test-grant-permission-partitioned-topic"}
	out, execErr, _, _ := TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	var permissions map[string][]pulsar.AuthAction
	err := json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, map[string][]pulsar.AuthAction{}, permissions)

	args = []string{"grant-permissions",
		"--role", "grant-partitioned-role",
		"--actions", "consume",
		"test-grant-permission-partitioned-topic",
	}
	_, execErr, _, _ = TestTopicCommands(GrantPermissionCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", "test-grant-permission-partitioned-topic"}
	out, execErr, _, _ = TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	err = json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []pulsar.AuthAction{"consume"}, permissions["grant-partitioned-role"])
}

func TestGrantPermissionArgError(t *testing.T) {
	args := []string{"grant-permissions", "--role", "test-arg-error-role", "--actions", "produce"}
	_, _, nameErr, _ := TestTopicCommands(GrantPermissionCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())

	args = []string{"grant-permissions", "args-error-topic"}
	_, _, _, err := TestTopicCommands(GrantPermissionCmd, args)
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"actions\", \"role\" not set", err.Error())

	args = []string{"grant-permissions", "--role", "", "--actions", "produce", "role-empty-topic"}
	_, execErr, _, _ := TestTopicCommands(GrantPermissionCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Invalid role name", execErr.Error())

	args = []string{"grant-permissions",
		"--role", "args-error-role",
		"--actions", "args-error-action",
		"invalid-actions-topic",
	}
	_, execErr, _, _ = TestTopicCommands(GrantPermissionCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "The auth action  only can be specified as 'produce', "+
		"'consume', or 'functions'. Invalid auth action 'args-error-action'", execErr.Error())
}
