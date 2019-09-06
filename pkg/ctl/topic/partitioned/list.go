package partitioned

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func ListTopicsCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for listing all exist topics under the specified namespace."
	desc.CommandPermission = "This command requires super-user permissions."

	listTopics := Example{
		Desc: "List all exist topics under the namespace <tenant/namespace>",
		Command: "pulsarctl topics list-partitioned-topics <tenant/namespace>",
	}
	desc.CommandExamples = []Example{listTopics}

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: `+----------------------------------------+
|                 TOPICS                 |
+----------------------------------------+
| <domain>://<tenant>/<namespace>/<topic>|
+----------------------------------------+`,
	}
	out =append(out, successOut, ArgError, TenantNotExistError)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"list-partitioned-topics",
		"List all exist partitioned topics under the specified namespace",
		desc.ToString(),
		"lp")

	vc.SetRunFuncWithNameArg(func() error {
		return doListTopics(vc)
	})
}

func doListTopics(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	namespace, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	topics, err := admin.Topics().ListPartitionedTopic(*namespace)
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{fmt.Sprintf("%s topics", namespace)})
		for _, t := range topics {
			table.Append([]string{t})
		}
		table.Render()
	}

	return err
}
