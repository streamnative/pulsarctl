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
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func deleteSchema(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Delete the latest schema for a topic"
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []pulsar.Example
	del := pulsar.Example{
		Desc:    "Delete the latest schema for a topic",
		Command: "pulsarctl schemas delete (topic name)",
	}
	examples = append(examples, del)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Deleted (topic name) successfully",
	}

	notTopicName := pulsar.Output{
		Desc: "you must specify a topic name, please check if the topic name is provided",
		Out:  "[✖]  the topic name is not specified or the topic name is specified more than one",
	}
	out = append(out, successOut, notTopicName)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete the latest schema for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"delete",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteSchema(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doDeleteSchema(vc *cmdutils.VerbCmd) error {
	topic := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Schemas().DeleteSchema(topic)
	if err == nil {
		vc.Command.Printf("Deleted %s successfully\n", topic)
	}
	return err
}
