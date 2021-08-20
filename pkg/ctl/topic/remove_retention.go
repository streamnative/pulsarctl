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

package topic

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	util "github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func RemoveRetentionCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Remove the retention policy for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	removeRetention := cmdutils.Example{
		Desc:    "Remove the retention policy for a topic",
		Command: "pulsarctl topics remove-retention tenant/namespace/topic",
	}
	examples = append(examples, removeRetention)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Remove the retention policy for a [topic] successfully",
	}

	noTopicName := cmdutils.Output{
		Desc: "you must specify a tenant/namespace/topic name, please check if the tenant/namespace/topic name is provided",
		Out:  "[✖]  the topic name is not specified or the topic name is specified more than one",
	}

	tenantNotExistError := cmdutils.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := cmdutils.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	out = append(out, successOut, noTopicName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"remove-retention",
		"Remove the retention policy for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"remove-retention",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveRetention(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doRemoveRetention(vc *cmdutils.VerbCmd) error {
	topic := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	topicName, err := util.GetTopicName(topic)
	if err != nil {
		return err
	}
	err = admin.Topics().RemoveRetention(*topicName)
	if err == nil {
		vc.Command.Printf("Remove the retention policy for a [%s] successfully\n", topic)
	}
	return err
}
