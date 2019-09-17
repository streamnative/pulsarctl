package namespace

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func SetMaxConsumersPerTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for setting the max consumers per topic of namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []Example
	set := Example{
		Desc: "Set the max consumers per topic of namespace <namespace-name> to <size>",
		Command: "pulsarctl namespaces set-max-consumers-per-topic --size <size> <namespace-name>",
	}
	desc.CommandExamples = append(examples, set)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: "Successfully set the max consumers per topic of namespace <namespace-name> to <size>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-max-consumers-per-topic",
		"Set the max consumers per topic of namespace",
		desc.ToString())

	var num int

	vc.SetRunFuncWithNameArg(func() error {
		return doSetMaxConsumersPerTopic(vc, num)
	})

	vc.FlagSetGroup.InFlagSet("Max Consumers Per Topic", func(set *pflag.FlagSet) {
		set.IntVar(&num, "size", -1, "max consumers per topic")
		cobra.MarkFlagRequired(set, "size")
	})
}

func doSetMaxConsumersPerTopic(vc *cmdutils.VerbCmd, max int) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	if max < 0 {
		return errors.New("the specified consumers value must bigger than 0")
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetMaxConsumersPerTopic(*ns, max)
	if err == nil {
		vc.Command.Printf("Successfully set the max consumers per topic of namespace %s to %d", ns.String(), max)
	}

	return err
}
