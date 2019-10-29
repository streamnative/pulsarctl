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
)

func GetMaxConsumersPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting the max consumers per subscription of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	set := pulsar.Example{
		Desc:    "Get the max consumers per subscription of the namespace (namespace-name)",
		Command: "pulsarctl namespaces get-max-consumers-per-subscription (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "The max consumers per subscription of the namespace (namespace-name) is (size)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-max-consumers-per-subscription",
		"Get the max consumers per subscription of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetMaxConsumerPerSubscription(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetMaxConsumerPerSubscription(vc *cmdutils.VerbCmd) error {
	ns, err := pulsar.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	max, err := admin.Namespaces().GetMaxConsumersPerSubscription(*ns)
	if err == nil {
		vc.Command.Printf("The max consumers per subscription of the namespace %s is %d\n", ns.String(), max)
	}

	return err
}
