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

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetInactiveTopicCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Get the inactive topic policies on a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    desc.CommandUsedFor,
		Command: "pulsarctl namespaces get-inactive-topic-policies tenant/namespace",
	})
	desc.CommandExamples = examples

	var out []cmdutils.Output
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-inactive-topic-policies",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString())

	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doGetInactiveTopicPolicies(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetInactiveTopicPolicies(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	response, err := admin.Namespaces().GetInactiveTopicPolicies(*ns)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), &response)
	}

	return err
}
