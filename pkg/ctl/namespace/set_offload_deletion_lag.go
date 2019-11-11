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
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func SetOffloadDeletionLagCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the offload deletion of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the offload deletion (duration) of the namespace (namespace-name)",
		Command: "pulsarctl namespaces set-offload-deletion-lag --lag (duration) (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the offload deletion lag of the namespace (namespace-name) to (duration)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-offload-deletion-lag",
		"Set the offload deletion lag of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var d string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetOffloadDeletionLag(vc, d)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Offload Deletion Lag", func(set *pflag.FlagSet) {
		set.StringVarP(&d, "lag", "l", "",
			"Duration to wait after offloading a ledger segment, before deleting the copy of that segment "+
				"from cluster local storage. (e.g. 1s, 1m, 1h)")
		cobra.MarkFlagRequired(set, "lag")
	})
}

func doSetOffloadDeletionLag(vc *cmdutils.VerbCmd, d string) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	t, err := time.ParseDuration(d)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetOffloadDeleteLag(*ns, t.Nanoseconds()/1e6)
	if err == nil {
		vc.Command.Printf("Successfully set the offload deletion lag of the namespace %s to %s\n", ns.String(), d)
	}

	return err
}
