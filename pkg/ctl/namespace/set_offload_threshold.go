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
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"

	util "github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func SetOffloadThresholdCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the offload threshold of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the offload threshold of the namespace (namespace-name) to (size)",
		Command: "pulsarctl namespaces set-offload-threshold --size (size) (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the offload threshold of the namespace (namespace-name) to (size)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-offload-threshold",
		"Set the offload threshold of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var threshold string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetOffloadThreshold(vc, threshold)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Offload Threshold", func(set *pflag.FlagSet) {
		set.StringVar(&threshold, "size", "-1",
			"Maximum number of bytes stored in the pulsar cluster for a topic before data will  "+
				"start being automatically offloaded to longterm  storage (e.g. 10m, 16g, 3t, 100)\n"+
				"Negative values disable automatic offload.\n"+
				"0 triggers offloading as soon as possible.")
		cobra.MarkFlagRequired(set, "size")
	})
}

func doSetOffloadThreshold(vc *cmdutils.VerbCmd, threshold string) error {
	ns, err := util.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	size, err := utils.ValidateSizeString(threshold)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetOffloadThreshold(*ns, size)
	if err == nil {
		vc.Command.Printf("Successfully set the offload threshold of the namespace %s to %s\n",
			ns.String(), threshold)
	}

	return err
}
