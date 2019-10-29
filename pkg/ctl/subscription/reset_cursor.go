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
	"strings"
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func ResetCursorCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for resetting the position of a " +
		"subscription to a position that is closest to the provided timestamp or messageId."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []pulsar.Example
	resetCursorTime := pulsar.Example{
		Desc: "Reset the position of the subscription (subscription-name) to a " +
			"position that is closest to the provided timestamp (time)",
		Command: "pulsarctl seek --time (time) (topic-name) (subscription-name)",
	}

	resetCursorMessageID := pulsar.Example{
		Desc: "Reset the position of the subscription <subscription-name> to a " +
			"position that is closest to the provided message id (message-id)",
		Command: "pulsarctl seek --message-id (message-id) (topic-name) (subscription-name)",
	}
	examples = append(examples, resetCursorTime, resetCursorMessageID)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Reset the cursor of the subscription (subscription-name) to (time)/(message-id) successfully",
	}

	resetFlagError := pulsar.Output{
		Desc: "the time is not specified or the message id is not specified",
		Out:  "[✖]  The reset position must be specified",
	}

	out = append(out, successOut, ArgsError, resetFlagError, TopicNotFoundError, SubNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"seek",
		"Reset the cursor to a position that is closest to the provided timestamp or messageId",
		desc.ToString(),
		desc.ExampleToString(),
		"seek")

	var t string
	var mID string

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doResetCursor(vc, t, mID)
	}, CheckSubscriptionNameTwoArgs)

	vc.FlagSetGroup.InFlagSet("ResetCursor", func(set *pflag.FlagSet) {
		set.StringVarP(&t, "time", "t", "",
			"time to reset back to (e.g. 1s, 1m, 1h)")
		set.StringVarP(&mID, "message-id", "m", "",
			"message id to reset back to (e.g. ledgerId:entryId)")
	})
}

func doResetCursor(vc *cmdutils.VerbCmd, t, mID string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	if t != "" && mID != "" {
		return errors.New("the time and message-id can not specified at the same time")
	}

	topic, err := pulsar.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	sName := vc.NameArgs[1]

	if topic.GetDomain().String() != "persistent" {
		return errors.New("the specified topic name is not a persistent topic")
	}

	admin := cmdutils.NewPulsarClient()
	switch {
	case t != "":
		d, err := time.ParseDuration(t)
		if err != nil {
			return err
		}
		resetTime := time.Now().Add(-d).UnixNano() / 1e6
		err = admin.Subscriptions().ResetCursorToTimestamp(*topic, sName, resetTime)
		if err == nil {
			vc.Command.Printf("Reset the cursor of the subscription %s to %s successfully\n", sName, t)
		}

		return err
	case mID != "":
		if len(strings.Split(mID, ":")) != 2 {
			return errors.Errorf("invalid position value : %s", mID)
		}
		id, err := pulsar.ParseMessageID(mID)
		if err != nil {
			return err
		}
		err = admin.Subscriptions().ResetCursorToMessageID(*topic, sName, *id)
		if err == nil {
			vc.Command.Printf("Reset the cursor of the subscription %s to %s successfully\n", sName, mID)
		}
		return err
	default:
		return errors.New("either timestamp or message-id should be specified")
	}
}
