package crud

import (
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func ListTopicsCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for listing all exist topics under the specified namespace."
	desc.CommandPermission = "This command requires admin permissions."

	listTopics := Example{
		Desc:    "List all exist topics under the namespace(tenant/namespace)",
		Command: "pulsarctl topics list (tenant/namespace)",
	}
	desc.CommandExamples = []Example{listTopics}

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: `+----------------------------------------------------------+---------------+
|                        TOPIC NAME                        | PARTITIONED ? |
+----------------------------------------------------------+---------------+
+----------------------------------------------------------+---------------+`,
	}

	argError := Output{
		Desc: "the namespace is not specified",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}
	out = append(out, successOut, argError, TenantNotExistError, NamespaceNotExistError)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all exist topics under the specified namespace",
		desc.ToString(),
		desc.ExampleToString(),
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
		table.SetHeader([]string{"topic name", "partitioned ?"})

		for _, v := range partitionedTopics {
			table.Append([]string{v, "Y"})
		}

		for _, v := range nonPartitionedTopics {
			table.Append([]string{v, "N"})
		}
		table.Render()
	}

	return err
}
