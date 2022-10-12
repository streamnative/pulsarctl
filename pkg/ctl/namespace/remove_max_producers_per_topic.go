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
)

func RemoveMaxProducersPerTopicCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used to remove the max producers per topic for a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Removes the max producers per topic setting for namespace (namespace-name)",
		Command: "pulsarctl namespaces remove-max-producers-per-topic (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully removed the max producers per topic for namespace (namespace-name)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"remove-max-producers-per-topic",
		"Remove the max producers per topic setting for a namespace",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveMaxProducersPerTopic(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doRemoveMaxProducersPerTopic(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().RemoveMaxProducersPerTopic(*ns)
	if err == nil {
		vc.Command.Printf("Successfully removed the max producers per topic for namespace %s\n", ns.String())
	}

	return err
}
