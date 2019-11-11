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

func GrantSubPermissionsCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for granting client roles to access a subscription of a namespace."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	grant := cmdutils.Example{
		Desc: "Grant the client roles (roles-name) to access the subscription (subscription-name) of the " +
			"namespace (namespace-name)",
		Command: "pulsarctl namespaces grant-subscription-permission --role (role1-name) --role (role2-name) " +
			"(namespace-name) (subscription-name)",
	}
	examples = append(examples, grant)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "Grant the client role <role-name> to access the subscription <subscription-name> of the " +
			"namespace <namespace-name> successfully",
	}

	argsError := cmdutils.Output{
		Desc: "the namespace name is not specified or the subscription name is not specified",
		Out:  "[✖]  need to specified namespace name and subscription name",
	}
	out = append(out, successOut, argsError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"grant-subscription-permission",
		"Grant a client role to access a subscription of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var role []string

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doGrantSubscriptionPermissions(vc, role)
	}, func(args []string) error {
		if len(args) != 2 {
			return errors.New("need to specified namespace name and subscription name")
		}
		return nil
	})

	vc.FlagSetGroup.InFlagSet("Grant Subscription Permissions", func(set *pflag.FlagSet) {
		set.StringSliceVar(&role, "role", nil,
			"Client role to which grant permissions")
		cobra.MarkFlagRequired(set, "role")
	})
}

func doGrantSubscriptionPermissions(vc *cmdutils.VerbCmd, role []string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := utils.GetNamespaceName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().GrantSubPermission(*ns, vc.NameArgs[1], role)
	if err == nil {
		vc.Command.Printf("Grant the client role %+v to access the subscription %s of "+
			"the namespace %s successfully\n", role, vc.NameArgs[1], ns.String())
	}

	return err
}
