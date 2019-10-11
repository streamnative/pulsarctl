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
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func GrantPermissionsCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for granting permissions to a client role to access a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	grant := pulsar.Example{
		Desc:    "Grant permission (action) to the client role (role-name) to access the namespace (namespace-name)",
		Command: "pulsarctl namespaces grant-permission --role (role-name) --actions (action) (namespace-name)",
	}

	grantActions := pulsar.Example{
		Desc: "Grant permissions (actions) to the client role (role-name) to access the namespace (namespace-name)",
		Command: "pulsarctl namespaces grant-permission --role (role-name) --actions (action-1) --actions (action-2) " +
			"(namespace-name)",
	}
	examples = append(examples, grant, grantActions)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "Grant permissions (actions) to the client role (role-name) to access the namespace (namespace-name)" +
			" successfully",
	}
	out = append(out, successOut, ArgError, AuthNotEnable)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"grant-permission",
		"Grant permissions to a client role to access a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var role string
	var actions []string

	vc.SetRunFuncWithNameArg(func() error {
		return doGrantPermissions(vc, role, actions)
	})

	vc.FlagSetGroup.InFlagSet("Grant Permissions", func(set *pflag.FlagSet) {
		set.StringVar(&role, "role", "",
			"Client role to which grant permissions")
		set.StringSliceVar(&actions, "actions", []string{},
			"Actions to be granted (produce,consume,functions)")
		cobra.MarkFlagRequired(set, "role")
		cobra.MarkFlagRequired(set, "actions")
	})
}

func doGrantPermissions(vc *cmdutils.VerbCmd, role string, actions []string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := pulsar.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	a, err := parseActions(actions)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().GrantNamespacePermission(*ns, role, a)
	if err == nil {
		vc.Command.Printf("Grant permissions %+v to the client role %s to access the"+
			" namespace %s successfully\n", a, role, ns.String())
	}

	return err
}

func parseActions(actions []string) ([]pulsar.AuthAction, error) {
	r := make([]pulsar.AuthAction, 0)
	for _, v := range actions {
		a, err := pulsar.ParseAuthAction(v)
		if err != nil {
			return nil, err
		}
		r = append(r, a)
	}
	return r, nil
}
