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

func RevokePermissionsCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for revoking a client role permissions of accessing a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions and " +
		"broker has read-writer permissions on the zookeeper."

	var examples []cmdutils.Example
	revoke := cmdutils.Example{
		Desc:    "Revoke the client role (role-name) of accessing the namespace (namespace-name)",
		Command: "pulsarctl namespaces revoke-permission --role (role-name) (namespace-name)",
	}
	examples = append(examples, revoke)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Revoke the client role (role-name) permissions of accessing the namespace (namespace-name) successfully",
	}
	out = append(out, successOut, ArgError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"revoke-permission",
		"Revoke a client role permissions of accessing a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var role string

	vc.SetRunFuncWithNameArg(func() error {
		return doRevokePermissions(vc, role)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Revoke Permissions", func(set *pflag.FlagSet) {
		set.StringVar(&role, "role", "",
			"Client role to which revoke permissions")
		cobra.MarkFlagRequired(set, "role")
	})
}

func doRevokePermissions(vc *cmdutils.VerbCmd, role string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().RevokeNamespacePermission(*ns, role)
	if err == nil {
		vc.Command.Printf("Revoke the client role %s permissions of accessing the namespace %s successfully\n",
			role, ns.String())
	}

	return err
}
