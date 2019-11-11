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

	"github.com/spf13/pflag"
)

func SetSchemaValidationEnforcedCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the schema whether open schema validation enforced."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	enable := cmdutils.Example{
		Desc:    "Enable schema validation enforced",
		Command: "pulsarctl namespaces set-schema-validation-enforced <namespace-name>",
	}

	disable := cmdutils.Example{
		Desc:    "Disable schema validation enforced",
		Command: "pulsarctl namespaces set-schema-validation-enforced --disable <namespace-name>",
	}
	examples = append(examples, enable, disable)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Enable/Disable schema validation enforced",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-schema-validation-enforced",
		"Enable/Disable schema validation enforced",
		desc.ToString(),
		desc.ExampleToString())

	var d bool

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSchemaValidationEnforced(vc, d)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Schema Validation Enforced", func(set *pflag.FlagSet) {
		set.BoolVarP(&d, "disable", "d", false,
			"Disable schema validation enforced")
	})
}

func doSetSchemaValidationEnforced(vc *cmdutils.VerbCmd, disable bool) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSchemaValidationEnforced(*ns, !disable)
	if err == nil {
		var out string
		if disable {
			out += "Disable "
		} else {
			out += "Enable "
		}
		vc.Command.Printf(out+"the namespace %s schema validation enforced\n", ns.String())
	}

	return err
}
