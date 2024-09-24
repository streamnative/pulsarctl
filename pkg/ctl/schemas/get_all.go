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
	"io"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getAllSchemas(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get all schemas for a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []cmdutils.Example
	del := cmdutils.Example{
		Desc:    "Get all schemas for a topic",
		Command: "pulsarctl schemas get-all (topic name)",
	}

	examples = append(examples, del)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "[\n" +
			"  {\n" +
			"    \"name\": \"test-schema\",\n" +
			"    \"schema\": {\n" +
			"      \"type\": \"record\",\n" +
			"      \"name\": \"Test\",\n" +
			"      \"fields\": [\n" +
			"        {\n" +
			"          \"name\": \"id\",\n" +
			"          \"type\": [\n" +
			"            \"null\",\n" +
			"            \"int\"\n" +
			"          ]\n" +
			"        },\n" +
			"        {\n" +
			"          \"name\": \"name\",\n" +
			"          \"type\": [\n" +
			"            \"null\",\n" +
			"            \"string\"\n" +
			"          ]\n" +
			"        }\n" +
			"      ]\n" +
			"    },\n" +
			"    \"type\": \"AVRO\",\n" +
			"    \"properties\": {}\n" +
			"  }\n" +
			"]",
	}

	failOut := cmdutils.Output{
		Desc: "HTTP 404 Not Found, please check if the topic name you entered is correct",
		Out:  "[✖]  code: 404 reason: Not Found",
	}

	notTopicName := cmdutils.Output{
		Desc: "you must specify a topic name, please check if the topic name is provided",
		Out:  "[✖]  the topic name is not specified or the topic name is specified more than once",
	}

	out = append(out, successOut, failOut, notTopicName)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-all",
		"Get the schema for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"get-all",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetAllSchemas(vc)
	}, "the topic name is not specified or the topic name is specified more than once")

	vc.EnableOutputFlagSet()
}

func doGetAllSchemas(vc *cmdutils.VerbCmd) error {
	topic := vc.NameArg

	admin := cmdutils.NewPulsarClient()
	infos, err := admin.Schemas().GetAllSchemas(topic)
	if err == nil {
		oc := cmdutils.NewOutputContent().
			WithObject(infos).
			WithTextFunc(func(w io.Writer) error {
				PrintSchemas(w, infos)
				return nil
			})
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
	}
	return err
}
