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

func GetMaxConsumersPerTopicCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the max consumers per topic of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Get the max consumers per topic of the namespace (namespace-name)",
		Command: "pulsarctl namespaces get-max-consumers-per-topic (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The max consumers per topic of the namespace (namespace-name) is (size)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-max-consumers-per-topic",
		"Get the max consumers per topic of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetMaxConsumerPerTopic(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetMaxConsumerPerTopic(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	max, err := admin.Namespaces().GetMaxConsumersPerTopic(*ns)
	if err == nil {
		vc.Command.Printf("The max consumers per topic of the namespace %s is %d\n", ns.String(), max)
	}

	return err
}
