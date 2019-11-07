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
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func ClearBacklogCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for clearing backlog for all topics of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	clear := cmdutils.Example{
		Desc:    "Clear backlog for all topics of the namespace (namespace-name)",
		Command: "pulsarctl namespaces clear-backlog (namespace-name)",
	}

	clearWithBundle := cmdutils.Example{
		Desc:    "Clear backlog for all topic of the namespace (namespace-name) with a bundle range <bundle>",
		Command: "pulsarctl namespaces clear-backlog --bundle (bundle) (namespace-name)",
	}

	clearWithSubName := cmdutils.Example{
		Desc: "Clear the specified subscription (subscription-name) backlog for all topics of the " +
			"namespace (namespace-name)",
		Command: "pulsarctl namespaces clear-backlog --subscription (subscription-name) (namespace-name)",
	}
	examples = append(examples, clear, clearWithBundle, clearWithSubName)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully clear backlog for all topics of the namespace (namespace-name)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"clear-backlog",
		"Clear backlog for all topics of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var sName, bundle string
	var force bool

	vc.SetRunFuncWithNameArg(func() error {
		return doClearBacklog(vc, sName, bundle, force)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Clear Backlog", func(set *pflag.FlagSet) {
		set.StringVar(&sName, "sub", "", "subscription name")
		set.StringVarP(&bundle, "bundle", "b", "", "{start-boundary}_{end-boundary}")
		set.BoolVarP(&force, "force", "f", false,
			"Whether to force clear backlog without prompt")
	})
}

func doClearBacklog(vc *cmdutils.VerbCmd, sName, bundle string, force bool) (err error) {
	if !force {
		if !prompt("Are you sure you want to clear the backlog?") {
			vc.Command.Println("Cancel clear backlog")
			return nil
		}
	}
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()

	switch {
	case sName != "":
		if bundle != "" {
			err = admin.Namespaces().ClearNamespaceBundleBacklogForSubscription(*ns, bundle, sName)
		} else {
			err = admin.Namespaces().ClearNamespaceBacklogForSubscription(*ns, sName)
		}
	case bundle != "":
		err = admin.Namespaces().ClearNamespaceBundleBacklog(*ns, bundle)
	default:
		err = admin.Namespaces().ClearNamespaceBacklog(*ns)
	}

	if err == nil {
		vc.Command.Printf("Successfully clear backlog for all topics of the namespace %s\n", ns.String())
	}

	return err
}

func prompt(prompt string) bool {
	for {
		fmt.Println(prompt + " (Y or N)")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response := scanner.Text()
		switch response {
		case "y":
			fallthrough
		case "yes":
			return true
		case "n":
			fallthrough
		case "no":
			return false
		}
	}
}
