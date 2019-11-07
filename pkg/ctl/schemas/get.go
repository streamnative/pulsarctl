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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func getSchema(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the schema for a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	vc.SetDescription(
		"get",
		"Get the schema for a topic",
		desc.ToString(),
		"get",
	)

	var examples []cmdutils.Example
	del := cmdutils.Example{
		Desc:    "Get the schema for a topic",
		Command: "pulsarctl schemas get (topic name)",
	}

	delWithVersion := cmdutils.Example{
		Desc: "Get the schema for a topic with version",
		Command: "pulsarctl schemas get (topic name) \n" +
			"\t--version 2",
	}

	examples = append(examples, del, delWithVersion)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"name\": \"test-schema\",\n" +
			"  \"schema\": {\n" +
			"    \"type\": \"record\",\n" +
			"    \"name\": \"Test\",\n" +
			"    \"fields\": [\n" +
			"      {\n" +
			"        \"name\": \"id\",\n" +
			"        \"type\": [\n" +
			"          \"null\",\n" +
			"          \"int\"\n" +
			"        ]\n" +
			"      },\n" +
			"      {\n" +
			"        \"name\": \"name\",\n" +
			"        \"type\": [\n" +
			"          \"null\",\n" +
			"          \"string\"\n" +
			"        ]\n" +
			"      }\n" +
			"    ]\n" +
			"  },\n" +
			"  \"type\": \"AVRO\",\n" +
			"  \"properties\": {}\n" +
			"}",
	}

	failOut := cmdutils.Output{
		Desc: "HTTP 404 Not Found, please check if the topic name you entered is correct",
		Out:  "[✖]  code: 404 reason: Not Found",
	}

	notTopicName := cmdutils.Output{
		Desc: "you must specify a topic name, please check if the topic name is provided",
		Out:  "[✖]  the topic name is not specified or the topic name is specified more than one",
	}

	out = append(out, successOut, failOut, notTopicName)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"Get the schema for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"get",
	)

	schemaData := &utils.SchemaData{}

	vc.SetRunFuncWithNameArg(func() error {
		return doGetSchema(vc, schemaData)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("SchemaConfig", func(flagSet *pflag.FlagSet) {
		flagSet.Int64Var(
			&schemaData.Version,
			"version",
			0,
			"the schema version info")
	})
}

func doGetSchema(vc *cmdutils.VerbCmd, schemaData *utils.SchemaData) error {
	topic := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	if schemaData.Version == 0 {
		schemaInfoWithVersion, err := admin.Schemas().GetSchemaInfoWithVersion(topic)
		if err == nil {
			PrintSchema(vc.Command.OutOrStdout(), schemaInfoWithVersion)
		}
		return err
	}
	info, err := admin.Schemas().GetSchemaInfoByVersion(topic, schemaData.Version)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), info)
	}

	return err
}
