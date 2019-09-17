package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetMaxProducersPerTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the max producers per topic of namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	set := Example{
		Desc: "Get the max producers per topic of namespace <namespace-name>",
		Command: "pulsarctl namespaces get-max-producers-per-topic <namespace-name>",
	}
	desc.CommandExamples = append(examples, set)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out: "The max producers per topic of namespace <namespace-name> is <size>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-max-producers-per-topic",
		"Get the max producers per topic of namespace",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetMaxProducersPerTopic(vc)
	})
}

func doGetMaxProducersPerTopic(vc *cmdutils.VerbCmd) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	max, err := admin.Namespaces().GetMaxProducersPerTopic(*ns)
	if err == nil {
		vc.Command.Printf("The max producers per topic of namespace %s is %d", ns.String(), max)
	}

	return err
}
