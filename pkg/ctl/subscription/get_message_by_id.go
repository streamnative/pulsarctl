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
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func GetMessageByIDCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting messages by the given legerID and entryID" +
		" for a subscription."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Get message by the given legerID an entryID",
		Command: "pulsarctl subscription get-message-by-id --leger-id (leger-id) --entry-id (entry-id) (topic-name)",
	})
	desc.CommandExamples = examples

	vc.SetDescription(
		"get-message-by-id",
		"Getting messages by the given legerID and entryID",
		desc.ToString(),
		desc.ExampleToString(),
	)

	var legerID int64
	var entryID int64

	vc.SetRunFuncWithNameArg(func() error {
		return doGetMessageByID(vc, legerID, entryID)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("GetMessageByID", func(set *pflag.FlagSet) {
		set.Int64VarP(&legerID, "leger-id", "l", 0, "ledger id pointing to the desired ledger")
		cobra.MarkFlagRequired(set, "leger-id")
		set.Int64VarP(&entryID, "entry-id", "e", 0, "entry id pointing to the desired entry")
		cobra.MarkFlagRequired(set, "entry-id")
	})
	vc.EnableOutputFlagSet()
}

func doGetMessageByID(vc *cmdutils.VerbCmd, legerID int64, entryID int64) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	client := cmdutils.NewPulsarClient()
	message, err := client.Subscriptions().GetMessageByID(*topic, legerID, entryID)
	if err == nil {
		propertiesJSON, err := json.Marshal(message.GetProperties())
		if err != nil {
			return err
		}

		vc.Command.Println(fmt.Sprintf(`Message ID: %s
Properties: %s
Message: %s`, message.GetMessageID(), propertiesJSON, hex.Dump(message.Payload)))
	}

	return err
}
