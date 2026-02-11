package topic

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestSetReplicationClustersCmd(t *testing.T) {
	g := gomega.NewWithT(t)

	topic := fmt.Sprintf("test-replication-clusters-topic-%s", test.RandomSuffix())

	args := []string{"create", topic, "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	args = []string{"set-replication-clusters", topic, "--clusters", "standalone"}
	out, execErr, nameErr, cmdErr := TestTopicCommands(SetReplicationClustersCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(nameErr).Should(gomega.BeNil())
	g.Expect(cmdErr).Should(gomega.BeNil())
	g.Expect(out).ShouldNot(gomega.BeNil())
	g.Expect(out.String()).ShouldNot(gomega.BeEmpty())

	// Since there is no get-replication-clusters command in this PR, we only test the set command success.
	// In a real scenario, we might want to verify using the client or adding a get command.
	// The set command output verification implies the call was successful.
}
