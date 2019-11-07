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

package subscription

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func ExpireCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for expiring messages that older than given expiry time (in seconds)" +
		" for a subscription."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []cmdutils.Example
	expire := cmdutils.Example{
		Desc: "Expire messages that older than given expire time (in seconds) for a subscription " +
			"<subscription-name> under a topic",
		Command: "pulsarctl subscription expire --expire-time (expire-time) (topic-name) (subscription-name)",
	}

	expireAllSub := cmdutils.Example{
		Desc: "Expire message that older than given expire time (in second) for all subscriptions " +
			"under a topic",
		Command: "pulsarctl subscriptions expire --all --expire-time (expire-time) (topic-name)",
	}
	examples = append(examples, expire, expireAllSub)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "Expire messages after (time)(s) for the subscription (subscription-name) of the topic (topic-name) " +
			"successfully",
	}
	out = append(out, successOut, ArgsError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"expire",
		"Expiring messages that older than given expire time (in seconds)",
		desc.ToString(),
		"expire")

	var time int64
	var all bool

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doExpire(vc, time, all)
	}, func(args []string) error {
		if len(args) > 2 || len(args) < 1 {
			return errors.New("need to specified the topic name and the subscription name")
		}
		return nil
	})

	vc.FlagSetGroup.InFlagSet("ExpireMessages", func(set *pflag.FlagSet) {
		set.Int64VarP(&time, "expire-time", "t", 0,
			"Expire messages older than time in seconds")
		cobra.MarkFlagRequired(set, "expire-time")
		set.BoolVarP(&all, "all", "a", false, "Expire all messages")
	})
}

func doExpire(vc *cmdutils.VerbCmd, time int64, all bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	var sName string
	if !all {
		if len(vc.NameArgs) != 2 {
			return errors.New("the subscription name should be specified")
		}
		sName = vc.NameArgs[1]
	}

	admin := cmdutils.NewPulsarClient()
	if all {
		err = admin.Subscriptions().ExpireAllMessages(*topic, time)
	} else {
		err = admin.Subscriptions().ExpireMessages(*topic, sName, time)
	}

	if err == nil {
		out := fmt.Sprintf("Expire messages after %d(s) for the subscription %s of the topic %s successfully",
			time, sName, topic.String())
		if all {
			out = fmt.Sprintf("Expire messages after %d(s) for all the subscriptions of the topic %s "+
				"successfully", time, topic.String())
		}
		vc.Command.Println(out)
	}

	return err
}
