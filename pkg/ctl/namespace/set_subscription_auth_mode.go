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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func SetSubscriptionAuthModeCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the default subscription auth mode of a namespace."
	desc.CommandPermission = "This command requires tenant admin and " +
		"a broker needs the read-write operations of the global zookeeper."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the default subscription auth mode (mode) of the namespace (namespace-name)",
		Command: "pulsarctl namespaces set-subscription-auth-mode --mode (mode) (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the default subscription auth mode of namespace <namespace-name> to <mode>",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-subscription-auth-mode",
		"Set the default subscription auth mode of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var mode string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSubscriptionAuthMode(vc, mode)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Subscription Auth Mode", func(set *pflag.FlagSet) {
		set.StringVarP(&mode, "mode", "m", "",
			"Subscription authorization mode of a namespace. (e.g. None, Prefix)")
		cobra.MarkFlagRequired(set, "mode")
	})
}

func doSetSubscriptionAuthMode(vc *cmdutils.VerbCmd, mode string) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	m, err := utils.ParseSubscriptionAuthMode(mode)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSubscriptionAuthMode(*ns, m)
	if err == nil {
		vc.Command.Printf("Successfully set the default subscription auth mode of namespace %s to %s",
			ns.String(), m.String())
	}

	return err
}
