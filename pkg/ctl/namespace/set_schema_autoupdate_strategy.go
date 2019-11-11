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

func SetSchemaAutoUpdateStrategyCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the schema auto-update strategy of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the schema auto-update strategy to (strategy)",
		Command: "pulsarctl namespaces set-schema-autoupdate-strategy --compatibility (strategy) (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the schema auto-update strategy of the namespace (namespace-name) to (strategy)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-schema-autoupdate-strategy",
		"Set the schema auto-update strategy of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var s string

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSchemaAutoUpdateStrategy(vc, s)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Schema Auto Update Strategy", func(set *pflag.FlagSet) {
		set.StringVarP(&s, "compatibility", "c", "",
			"Compatibility level required for new schemas created via a Producer. Possible values "+
				"(AutoUpdateDisabled, Backward, Forward, Full, AlwaysCompatible, BackwardTransitive, "+
				"ForwardTransitive, FullTransitive)")
	})
}

func doSetSchemaAutoUpdateStrategy(vc *cmdutils.VerbCmd, strategy string) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	s := utils.AutoUpdateDisabled
	if strategy != "" {
		s, err = utils.ParseSchemaAutoUpdateCompatibilityStrategy(strategy)
		if err != nil {
			return err
		}
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSchemaAutoUpdateCompatibilityStrategy(*ns, s)
	if err == nil {
		vc.Command.Printf("Successfully set the schema auto-update strategy of the namespace %s to %s\n",
			ns.String(), s.String())
	}

	return err
}
