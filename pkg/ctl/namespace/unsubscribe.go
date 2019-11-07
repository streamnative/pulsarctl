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
	"github.com/spf13/pflag"
)

func UnsubscribeCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for unsubscribing the specified " +
		"subscription for all topics of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	unsub := cmdutils.Example{
		Desc:    "Unsubscribe the specified subscription <subscription-name> for all topic of the namespace (namespace-name)",
		Command: "pulsarctl namespaces unsubscribe (namespace-name) (subscription-name)",
	}

	unsubWithBundle := cmdutils.Example{
		Desc: "Unsubscribe the specified subscription (subscription-name) for all topic of the namespace (namespace-name) " +
			"with bundle range <bundle>",
		Command: "pulsarctl namespaces unsubscribe --bundle (bundle) (namespace-name) (subscription-name)",
	}
	examples = append(examples, unsub, unsubWithBundle)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "Successfully unsubscribe the subscription (subscription-name) " +
			"for all topics of the namespace (namespace-name)",
	}

	argsError := cmdutils.Output{
		Desc: "the namespace name is not specified or the subscription name is not specified",
		Out:  "[✖]  need two arguments apply to the command",
	}
	out = append(out, successOut, argsError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"unsubscribe",
		"Unsubscribe the specified subscription for all topic of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var bundle string

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doUnsubscribe(vc, bundle)
	}, func(args []string) error {
		if len(args) != 2 {
			return errors.New("need two arguments apply to the command")
		}
		return nil
	})

	vc.FlagSetGroup.InFlagSet("Unsubscribe", func(set *pflag.FlagSet) {
		set.StringVarP(&bundle, "bundle", "b", "",
			"{start_boundary}_{end_boundary}")
	})
}

func doUnsubscribe(vc *cmdutils.VerbCmd, bundle string) (err error) {
	ns, err := utils.GetNamespaceName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	sName := vc.NameArgs[1]

	admin := cmdutils.NewPulsarClient()
	if bundle == "" {
		err = admin.Namespaces().UnsubscribeNamespace(*ns, sName)
	} else {
		err = admin.Namespaces().UnsubscribeNamespaceBundle(*ns, bundle, sName)
	}

	if err == nil {
		if bundle == "" {
			vc.Command.Printf("Successfully unsubscribe the subscription %s for all topics of "+
				"the namespace %s", sName, ns.String())
		} else {
			vc.Command.Printf("Successfully unsubscribe the subscription %s for all topics of "+
				"the namespace %s with bundle range %s", sName, ns.String(), bundle)
		}
	}

	return
}
