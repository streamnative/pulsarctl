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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func SetMaxProducersPerTopicCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the max producers per topic of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the max producers per topic of the namespace (namespace-name) to (size)",
		Command: "pulsarctl namespaces set-max-producers-per-topic --size (size) (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the max producers per topic of namespace (namespace-name) to (size)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-max-producers-per-topic",
		"Set max producers per topic of namespace",
		desc.ToString(),
		desc.ExampleToString())

	var num int

	vc.SetRunFuncWithNameArg(func() error {
		return doSetMaxProducersPerTopic(vc, num)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Max Producers Per Topic", func(set *pflag.FlagSet) {
		set.IntVar(&num, "size", -1, "max producers per topic")
		cobra.MarkFlagRequired(set, "size")
	})
}

func doSetMaxProducersPerTopic(vc *cmdutils.VerbCmd, max int) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	if max < 0 {
		return errors.New("the specified producers value must bigger than 0")
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetMaxProducersPerTopic(*ns, max)
	if err == nil {
		vc.Command.Printf("Successfully set the max producers per topic of the namespace %s to %d\n", ns.String(), max)
	}

	return err
}
