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
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	ctlUtil "github.com/streamnative/pulsarctl/pkg/ctl/utils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func SetBacklogQuotaCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set backlog quota policy for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Set backlog quota policy for a topic",
		Command: "pulsarctl topics set-backlog-quota topic -l 1k -p producer_exception",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set backlog quota policy successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-backlog-quota",
		"Set backlog quota policy for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-backlog-quota",
	)
	backlogQuotaCmdData := &utils.BacklogQuotaCmdData{}
	vc.SetRunFuncWithNameArg(func() error {
		return doSetBacklogQuota(vc, backlogQuotaCmdData)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("BacklogQuota", func(set *pflag.FlagSet) {
		set.StringVarP(
			&backlogQuotaCmdData.LimitStr,
			"limit",
			"l",
			"",
			"Size limit (eg: 10M, 16G)")
		set.StringVarP(
			&backlogQuotaCmdData.PolicyStr,
			"policy",
			"p",
			"",
			"Retention policy to enforce when the limit is reached. Valid options are: "+policyTypeStr)
	})
	vc.EnableOutputFlagSet()
}

func doSetBacklogQuota(vc *cmdutils.VerbCmd, backlogQuotaCmdData *utils.BacklogQuotaCmdData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	backlogQuotaData := &utils.BacklogQuotaData{}
	_, validPolicy := policyTypeMap[backlogQuotaCmdData.PolicyStr]
	if !validPolicy {
		msg := "Invalid retention policy type '%s'. Valid options are: %s"
		return errors.Errorf(msg, backlogQuotaCmdData.PolicyStr, policyTypeStr)
	}
	limit, err := ctlUtil.ValidateSizeString(backlogQuotaCmdData.LimitStr)
	if err != nil {
		return err
	}
	backlogQuotaData.Policy = backlogQuotaCmdData.PolicyStr
	backlogQuotaData.Limit = limit
	err = admin.Topics().SetBacklogQuota(*topic, *backlogQuotaData)
	if err == nil {
		vc.Command.Printf("Set backlog quota policy successfully for [%s]\n", topic.String())
	}
	return err
}

var policyTypeMap = map[string]int{
	"producer_request_hold":     0,
	"producer_exception":        1,
	"consumer_backlog_eviction": 2,
}

var policyTypeStr = "[producer_request_hold, producer_exception, consumer_backlog_eviction]"
