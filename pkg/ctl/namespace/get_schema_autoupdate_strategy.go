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

func GetSchemaAutoUpdateStrategyCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the schema auto-update strategy of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the schema auto-update strategy of the namespace (namespace-name)",
		Command: "pulsarctl namespaces get-schema-autoupdate-strategy (namespace-name)",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The schema auto-update strategy of the namespace (namespace-name) is (strategy)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-schema-autoupdate-strategy",
		"Get the schema auto-update strategy of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetSchemaAutoUpdateStrategy(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetSchemaAutoUpdateStrategy(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	s, err := admin.Namespaces().GetSchemaAutoUpdateCompatibilityStrategy(*ns)
	if err == nil {
		vc.Command.Printf("The schema auto-update strategy of the namespace %s is %s\n", ns.String(), s.String())
	}

	return err
}
