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
	"strings"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/rest"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetBookieAffinityGroupCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Get bookie affinity group configured for a namespace"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get bookie affinity group configured for a namespace",
		Command: "pulsarctl namespaces get-bookie-affinity-group tenant/namespace",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "{\n  \"bookkeeperAffinityGroupPrimary\": \"primary\",\n  \"bookkeeperAffinityGroupSecondary\": \"secondary\"\n}",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-bookie-affinity-group",
		"Get bookie affinity group configured for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		return doGetBookieAffinityGroup(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetBookieAffinityGroup(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	group, err := admin.Namespaces().GetBookieAffinityGroup(vc.NameArg)
	if err != nil {
		if restErr, ok := err.(rest.Error); ok && restErr.Code == 404 &&
			strings.Contains(strings.ToLower(restErr.Reason), "local-policies") {
			return vc.OutputConfig.WriteOutput(
				vc.Command.OutOrStdout(),
				cmdutils.NewOutputContent().WithObject(nil),
			)
		}
		return err
	}

	oc := cmdutils.NewOutputContent().WithObject(group)
	return vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
}

func SetBookieAffinityGroupCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Set bookie affinity group configured for a namespace"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc: "Set bookie affinity group configured for a namespace",
		Command: "pulsarctl namespaces set-bookie-affinity-group tenant/namespace \n" +
			"\t--primary-group primary-group \n" +
			"\t--secondary-group secondary-group",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set bookie affinity group successfully for [tenant/namespace]",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-bookie-affinity-group",
		"Set bookie affinity group configured for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	data := utils.BookieAffinityGroupData{}
	vc.FlagSetGroup.InFlagSet("Bookie Affinity Group", func(set *pflag.FlagSet) {
		set.StringVar(&data.BookkeeperAffinityGroupPrimary, "primary-group", "", "primary affinity group")
		set.StringVar(&data.BookkeeperAffinityGroupSecondary, "secondary-group", "", "secondary affinity group")
		_ = cobra.MarkFlagRequired(set, "primary-group")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doSetBookieAffinityGroup(vc, data)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doSetBookieAffinityGroup(vc *cmdutils.VerbCmd, data utils.BookieAffinityGroupData) error {
	admin := cmdutils.NewPulsarClient()
	err := admin.Namespaces().SetBookieAffinityGroup(vc.NameArg, data)
	if err == nil {
		vc.Command.Printf("Set bookie affinity group successfully for [%s]\n", vc.NameArg)
	}
	return err
}

func DeleteBookieAffinityGroupCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Delete bookie affinity group configured for a namespace"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	del := cmdutils.Example{
		Desc:    "Delete bookie affinity group configured for a namespace",
		Command: "pulsarctl namespaces delete-bookie-affinity-group tenant/namespace",
	}
	examples = append(examples, del)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Deleted bookie affinity group successfully for [tenant/namespace]",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete-bookie-affinity-group",
		"Delete bookie affinity group configured for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteBookieAffinityGroup(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doDeleteBookieAffinityGroup(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	err := admin.Namespaces().DeleteBookieAffinityGroup(vc.NameArg)
	if err == nil {
		vc.Command.Printf("Deleted bookie affinity group successfully for [%s]\n", vc.NameArg)
	}
	return err
}
