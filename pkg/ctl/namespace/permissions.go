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

func GetPermissionsCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting permissions configure data of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	getNs := pulsar.Example{
		Desc:    "Get permissions configure data of a namespace (tenant)/(namespace)",
		Command: "pulsarctl namespaces permissions (tenant)/(namespace)",
	}
	examples = append(examples, getNs)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"<role>\": [\n" +
			"    \"<action>\"\n" +
			"  ]" +
			"\n}",
	}
	out = append(out, successOut, ArgError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"permissions",
		"Get permissions configure data of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetPermissions(vc)
	})
}

func doGetPermissions(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	ns, err := pulsar.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	data, err := admin.Namespaces().GetNamespacePermissions(*ns)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), data)
	}

	return err
}
