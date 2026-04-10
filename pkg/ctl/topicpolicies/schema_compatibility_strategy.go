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

package topicpolicies

import (
	"fmt"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetSchemaCompatibilityStrategyCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-schema-compatibility-strategy", "Get schema compatibility strategy", "Get schema compatibility strategy", "")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		value, err := policies.GetSchemaCompatibilityStrategy(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		if value == nil {
			return writePolicyOutput(vc, "", "\n")
		}
		return writePolicyOutput(vc, value.String(), "%s\n", value.String())
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetSchemaCompatibilityStrategyCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var strategy string
	vc.SetDescription("set-schema-compatibility-strategy", "Set schema compatibility strategy", "Set schema compatibility strategy", "")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("SchemaCompatibilityStrategy", func(set *pflag.FlagSet) {
		set.StringVarP(&strategy, "compatibility", "c", "", "schema compatibility strategy")
	})
	vc.SetRunFuncWithNameArg(func() error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		parsed, err := utils.ParseSchemaCompatibilityStrategy(strategy)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		err = policies.SetSchemaCompatibilityStrategy(vc.Command.Context(), *topic, parsed)
		if err == nil {
			vc.Command.Printf("Set schema compatibility strategy successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveSchemaCompatibilityStrategyCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-schema-compatibility-strategy", "Removed schema compatibility strategy", func(global bool) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.RemoveSchemaCompatibilityStrategy(vc.Command.Context(), *topic)
	})
}

func writePolicyOutputString(vc *cmdutils.VerbCmd, value string) error {
	return writePolicyOutput(vc, value, fmt.Sprintf("%%s\n"), value)
}
