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

package schemas

import (
	"encoding/json"
	"os"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func testCompatibility(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Test schema compatibility"
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    desc.CommandUsedFor,
		Command: "pulsarctl schemas compatibility (topic name) --filename (schema file path)",
	})
	desc.CommandExamples = examples

	vc.SetDescription(
		"compatibility",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString(),
		"compatibility",
	)

	schemaData := &utils.SchemaData{}
	vc.FlagSetGroup.InFlagSet("SchemaConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(&schemaData.Filename, "filename", "f", "", "filename")
		_ = cobra.MarkFlagRequired(flagSet, "filename")
	})
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		return doTestCompatibility(vc, schemaData)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doTestCompatibility(vc *cmdutils.VerbCmd, schemaData *utils.SchemaData) error {
	var payload utils.PostSchemaPayload

	file, err := os.ReadFile(schemaData.Filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(file, &payload); err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	compatibility, err := admin.Schemas().TestCompatibilityWithPostSchemaPayload(vc.NameArg, payload)
	if err != nil {
		return err
	}

	return vc.OutputConfig.WriteOutput(
		vc.Command.OutOrStdout(),
		cmdutils.NewOutputContent().WithObject(compatibility),
	)
}
