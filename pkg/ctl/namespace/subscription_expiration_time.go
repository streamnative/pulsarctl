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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetSubscriptionExpirationTimeCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the subscription expiration time of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the subscription expiration time of the namespace (namespace-name)",
		Command: "pulsarctl namespaces get-subscription-expiration-time (namespace-name)",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The subscription expiration time of the namespace (namespace-name) is (time)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-subscription-expiration-time",
		"Get the subscription expiration time of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetSubscriptionExpirationTime(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetSubscriptionExpirationTime(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	expirationTime, err := admin.Namespaces().GetSubscriptionExpirationTime(*ns)
	if err == nil {
		if expirationTime == -1 {
			vc.Command.Printf("The subscription expiration time of the namespace %s is not set\n", ns.String())
		} else {
			vc.Command.Printf("The subscription expiration time of the namespace %s is %d\n", ns.String(), expirationTime)
		}
	}
	return err
}

func SetSubscriptionExpirationTimeCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the subscription expiration time of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the subscription expiration time of the namespace (namespace-name) to (time)",
		Command: "pulsarctl namespaces set-subscription-expiration-time --time (time) (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the subscription expiration time of the namespace (namespace-name) to (time)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-subscription-expiration-time",
		"Set the subscription expiration time of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	var expirationTime int
	vc.FlagSetGroup.InFlagSet("Subscription Expiration Time", func(set *pflag.FlagSet) {
		set.IntVarP(&expirationTime, "time", "t", -1, "subscription expiration time in minutes")
		_ = cobra.MarkFlagRequired(set, "time")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSubscriptionExpirationTime(vc, expirationTime)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.EnableOutputFlagSet()
}

func doSetSubscriptionExpirationTime(vc *cmdutils.VerbCmd, expirationTime int) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	if expirationTime < 0 {
		return errors.New("the specified subscription expiration time must bigger than or equal to 0")
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSubscriptionExpirationTime(*ns, expirationTime)
	if err == nil {
		vc.Command.Printf("Successfully set the subscription expiration time of the namespace %s to %d\n",
			ns.String(), expirationTime)
	}
	return err
}

func RemoveSubscriptionExpirationTimeCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for removing the subscription expiration time of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	remove := cmdutils.Example{
		Desc:    "Remove the subscription expiration time of the namespace (namespace-name)",
		Command: "pulsarctl namespaces remove-subscription-expiration-time (namespace-name)",
	}
	examples = append(examples, remove)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully removed the subscription expiration time of the namespace (namespace-name)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"remove-subscription-expiration-time",
		"Remove the subscription expiration time of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveSubscriptionExpirationTime(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doRemoveSubscriptionExpirationTime(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().RemoveSubscriptionExpirationTime(*ns)
	if err == nil {
		vc.Command.Printf("Successfully removed the subscription expiration time of the namespace %s\n", ns.String())
	}
	return err
}
