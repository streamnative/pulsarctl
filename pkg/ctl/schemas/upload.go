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
	"io/ioutil"

	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func uploadSchema(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Update the schema for a topic"
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []cmdutils.Example

	upload := cmdutils.Example{
		Desc: "Update the schema for a topic",
		Command: "pulsarctl schemas upload \n" +
			"(topic name) \n " +
			"--filename (the file path of schema)",
	}

	examples = append(examples, upload)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Upload (topic name) successfully",
	}

	notTopicName := cmdutils.Output{
		Desc: "you must specify a topic name, please check if the topic name is provided",
		Out:  "[✖]  the topic name is not specified or the topic name is specified more than one",
	}

	filePathNotExist := cmdutils.Output{
		Desc: "no such file or directory",
		Out:  "[✖]  open (file path): no such file or directory",
	}

	out = append(out, successOut, notTopicName, filePathNotExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"upload",
		"Update the schema for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"upload",
	)
	schemaData := &utils.SchemaData{}

	vc.SetRunFuncWithNameArg(func() error {
		return doUploadSchema(vc, schemaData)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("SchemaConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&schemaData.Filename,
			"filename",
			"f",
			"",
			"filename")
	})
}

func doUploadSchema(vc *cmdutils.VerbCmd, schemaData *utils.SchemaData) error {
	var payload utils.PostSchemaPayload
	topic := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	file, err := ioutil.ReadFile(schemaData.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &payload)
	if err != nil {
		return err
	}

	err = admin.Schemas().CreateSchemaByPayload(topic, payload)
	if err == nil {
		vc.Command.Printf("Upload %s successfully\n", topic)
	}

	return err
}
