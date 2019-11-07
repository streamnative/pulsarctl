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

func RevokeSubPermissionsCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for revoking a client role permissions of " +
		"accessing a subscription of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions and " +
		"broker has read-writer permissions on the zookeeper."

	var examples []cmdutils.Example
	revoke := cmdutils.Example{
		Desc: "Revoke a client role (role-name) permissions of accessing the subscription " +
			"(subscription-name) of the (namespace-name)",
		Command: "pulsarctl namespaces revoke-subscription-permission --role (role-name) " +
			"(namespace-name) (subscription-name)",
	}
	examples = append(examples, revoke)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "Revoke the client role (role-name) permissions of accessing the subscription " +
			"(subscription-name) of the namespace (namespace-name) successfully",
	}

	argsError := cmdutils.Output{
		Desc: "the namespace name is not specified or the subscription name is not specified",
		Out:  "[✖]  need to specified namespace name and subscription name",
	}
	out = append(out, successOut, argsError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"revoke-subscription-permission",
		"Revoke a client role permissions of accessing a subscription of a namespace",
		desc.ToString(),
		desc.ToString())

	var role string

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doRevokeSubPermissions(vc, role)
	}, func(args []string) error {
		if len(args) != 2 {
			return errors.New("need to specified namespace name and subscription name")
		}
		return nil
	})

	vc.FlagSetGroup.InFlagSet("Revoke Subscription Permissions", func(set *pflag.FlagSet) {
		set.StringVar(&role, "role", "",
			"Client role to which revoke permissions")
		cobra.MarkFlagRequired(set, "role")
	})
}

func doRevokeSubPermissions(vc *cmdutils.VerbCmd, role string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := utils.GetNamespaceName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().RevokeSubPermission(*ns, vc.NameArgs[1], role)
	if err == nil {
		vc.Command.Printf("Revoke the client role %s permissions of accessing the "+
			"subscription %s of the namespace %s successfully\n",
			role, vc.NameArgs[1], ns.String())
	}

	return err
}
