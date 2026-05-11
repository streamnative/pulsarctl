// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package namespace

import (
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetMaxTopicsPerNamespaceCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the max topics per namespace of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the max topics per namespace of the namespace (namespace-name)",
		Command: "pulsarctl namespaces get-max-topics-per-namespace (namespace-name)",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The max topics per namespace of the namespace (namespace-name) is (size)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-max-topics-per-namespace",
		"Get the max topics per namespace of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetMaxTopicsPerNamespace(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetMaxTopicsPerNamespace(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	max, err := admin.Namespaces().GetMaxTopicsPerNamespace(*ns)
	if err == nil {
		if max == -1 {
			vc.Command.Printf("The max topics per namespace of the namespace %s is not set\n", ns.String())
		} else {
			vc.Command.Printf("The max topics per namespace of the namespace %s is %d\n", ns.String(), max)
		}
	}
	return err
}

func SetMaxTopicsPerNamespaceCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the max topics per namespace of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the max topics per namespace of the namespace (namespace-name) to (size)",
		Command: "pulsarctl namespaces set-max-topics-per-namespace --max-topics-per-namespace (size) (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the max topics per namespace of the namespace (namespace-name) to (size)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-max-topics-per-namespace",
		"Set the max topics per namespace of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	var max int
	vc.FlagSetGroup.InFlagSet("Max Topics Per Namespace", func(set *pflag.FlagSet) {
		set.IntVarP(&max, "max-topics-per-namespace", "t", -1, "max topics per namespace")
		_ = cobra.MarkFlagRequired(set, "max-topics-per-namespace")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doSetMaxTopicsPerNamespace(vc, max)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.EnableOutputFlagSet()
}

func doSetMaxTopicsPerNamespace(vc *cmdutils.VerbCmd, max int) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	if max < 0 {
		return errors.New("the specified max topics value must bigger than 0")
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetMaxTopicsPerNamespace(*ns, max)
	if err == nil {
		vc.Command.Printf("Successfully set the max topics per namespace of the namespace %s to %d\n",
			ns.String(), max)
	}
	return err
}

func RemoveMaxTopicsPerNamespaceCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for removing the max topics per namespace of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	remove := cmdutils.Example{
		Desc:    "Remove the max topics per namespace of the namespace (namespace-name)",
		Command: "pulsarctl namespaces remove-max-topics-per-namespace (namespace-name)",
	}
	examples = append(examples, remove)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully removed the max topics per namespace of the namespace (namespace-name)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"remove-max-topics-per-namespace",
		"Remove the max topics per namespace of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveMaxTopicsPerNamespace(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doRemoveMaxTopicsPerNamespace(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().RemoveMaxTopicsPerNamespace(*ns)
	if err == nil {
		vc.Command.Printf("Successfully removed the max topics per namespace of the namespace %s\n", ns.String())
	}
	return err
}
