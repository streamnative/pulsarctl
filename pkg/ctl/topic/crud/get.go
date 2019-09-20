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

package crud

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetTopicCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used for getting the metadata of an exist topic."
	desc.CommandPermission = "This command requires namespace admin permissions."

	var examples []Example
	getTopic := Example{
		Desc:    "Get hte metadata of an exist topic <topic-name> metadata",
		Command: "pulsarctl topics get <topic-name>",
	}
	desc.CommandExamples = append(examples, getTopic)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "{\n" +
			"  \"partitions\": \"<partitions>\"\n" +
			"}",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"Get the specified topic metadata",
		desc.ToString(),
		"get")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetTopic(vc)
	})
}

func doGetTopic(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	meta, err := admin.Topics().GetMetadata(*topic)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), meta)
	}

	return err
}
