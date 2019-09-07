package topic

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
		Command: "pulsarctl topics list <tenant/namespace>",
	}
	desc.CommandExamples = []Example{listTopics}

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:
`+--------------------------------+--------------------------------+
|   PUBLIC/DEFAULT PARTITIONED   | PUBLIC/DEFAULT NON-PARTITIONED |
|             TOPICS             |             TOPICS             |
+--------------------------------+--------------------------------+
|                                |                                |
+--------------------------------+--------------------------------+`,
	}
	out =append(out, successOut, ArgError, TenantNotExistError, NamespaceNotExistError)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
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
	partitionedTopics, nonPartitionedTopics, err := admin.Topics().List(*namespace)
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{
			fmt.Sprintf("%s partitioned topics", namespace),
			fmt.Sprintf("%s non-partitioned topics", namespace),
		})

		var row int
		if len(partitionedTopics) >= len(nonPartitionedTopics) {
			row = len(partitionedTopics)
		} else {
			row = len(nonPartitionedTopics)
		}

		for i := 0; i < row; i++ {
			table.Append([]string{getValue(partitionedTopics, i), getValue(nonPartitionedTopics, i)})
		}

		table.Render()
	}

	return err
}

func getValue(array []string, index int) string {
	if index >= len(array) {
		return ""
	}
	return array[index]
}
