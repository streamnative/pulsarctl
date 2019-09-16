package crud

import (
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

var o = `+----------------------+
|    SUBSCRIPTIONS     |
+----------------------+
+----------------------+
`

func ListCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for listing all existing subscriptions under a topic."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []Example
	list := Example{
		Desc:    "List all existing subscriptions under a topic <topic-name>",
		Command: "pulsarctl subscriptions list <topic-name>",
	}
	desc.CommandExamples = append(examples, list)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  o,
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"list all existing subscriptions under a topic",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doList(vc)
	})
}

func doList(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	r, err := admin.Subscriptions().List(*topic)
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Subscriptions"})
		for _, v := range r {
			table.Append([]string{v})
		}
		table.Render()
	}

	return err
}
