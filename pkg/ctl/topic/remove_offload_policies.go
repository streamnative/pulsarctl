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

func RemoveOffloadPoliciesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Remove the offload policies for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    desc.CommandUsedFor,
		Command: "pulsarctl topics remove-offload-policies tenant/namespace/topic",
	})
	desc.CommandExamples = examples

	vc.SetDescription(
		"remove-offload-policies",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString(),
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveOffloadPolicies(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doRemoveOffloadPolicies(vc *cmdutils.VerbCmd) error {
	topic := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	topicName, err := util.GetTopicName(topic)
	if err != nil {
		return err
	}
	err = admin.Topics().RemoveOffloadPolicies(*topicName)
	if err == nil {
		vc.Command.Printf("Remove the offload policies successfully for [%s]", topic)
	}
	return err
}
