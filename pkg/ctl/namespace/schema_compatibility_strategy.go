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
	"strings"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetSchemaCompatibilityStrategyCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the schema compatibility strategy of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the schema compatibility strategy of the namespace (namespace-name)",
		Command: "pulsarctl namespaces get-schema-compatibility-strategy (namespace-name)",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The schema compatibility strategy of the namespace (namespace-name) is (strategy)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-schema-compatibility-strategy",
		"Get the schema compatibility strategy of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetSchemaCompatibilityStrategy(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetSchemaCompatibilityStrategy(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	s, err := admin.Namespaces().GetSchemaCompatibilityStrategy(*ns)
	if err == nil {
		vc.Command.Printf("The schema compatibility strategy of the namespace %s is %s\n", ns.String(), s.String())
	}
	return err
}

func SetSchemaCompatibilityStrategyCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the schema compatibility strategy of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the schema compatibility strategy to (strategy)",
		Command: "pulsarctl namespaces set-schema-compatibility-strategy --compatibility (strategy) (namespace-name)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the schema compatibility strategy of the namespace (namespace-name) to (strategy)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-schema-compatibility-strategy",
		"Set the schema compatibility strategy of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	var strategy string
	vc.FlagSetGroup.InFlagSet("Schema Compatibility Strategy", func(set *pflag.FlagSet) {
		set.StringVarP(&strategy, "compatibility", "c", "",
			"Compatibility level required for new schemas created via a Producer. Possible values "+
				"(UNDEFINED, ALWAYS_INCOMPATIBLE, ALWAYS_COMPATIBLE, BACKWARD, FORWARD, FULL, "+
				"BACKWARD_TRANSITIVE, FORWARD_TRANSITIVE, FULL_TRANSITIVE)")
		_ = cobra.MarkFlagRequired(set, "compatibility")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSchemaCompatibilityStrategy(vc, strategy)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doSetSchemaCompatibilityStrategy(vc *cmdutils.VerbCmd, strategy string) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	s, err := utils.ParseSchemaCompatibilityStrategy(strings.ToUpper(strategy))
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSchemaCompatibilityStrategy(*ns, s)
	if err == nil {
		vc.Command.Printf("Successfully set the schema compatibility strategy of the namespace %s to %s\n",
			ns.String(), s.String())
	}
	return err
}
