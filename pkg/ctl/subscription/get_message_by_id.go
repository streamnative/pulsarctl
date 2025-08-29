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

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

type readMessage struct {
	Properties      map[string]string `json:"properties"`
	MessageID       utils.MessageID   `json:"messageId"`
	Payload         []byte            `json:"payload"`
	PayloadAsString string            `json:"PayloadString"`
}

func GetMessageByIDCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting messages by the given ledgerID and entryID" +
		" for a subscription."
	desc.CommandPermission = "This command requires tenant admin and namespace produce or consume permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Get message by the given ledgerID an entryID",
		Command: "pulsarctl subscription get-message-by-id --ledger-id (ledger-id) --entry-id (entry-id) (topic-name)",
	})
	desc.CommandExamples = examples

	vc.SetDescription(
		"get-message-by-id",
		"Getting messages by the given ledgerID and entryID",
		desc.ToString(),
		desc.ExampleToString(),
	)

	var ledgerID int64
	var entryID int64

	vc.SetRunFuncWithNameArg(func() error {
		return doGetMessageByID(vc, ledgerID, entryID)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("GetMessageByID", func(set *pflag.FlagSet) {
		set.Int64VarP(&ledgerID, "ledger-id", "l", 0, "ledger id pointing to the desired ledger")
		_ = cobra.MarkFlagRequired(set, "ledger-id")
		set.Int64VarP(&entryID, "entry-id", "e", 0, "entry id pointing to the desired entry")
		_ = cobra.MarkFlagRequired(set, "entry-id")
	})
	vc.EnableOutputFlagSet()
}

func doGetMessageByID(vc *cmdutils.VerbCmd, ledgerID int64, entryID int64) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	client := cmdutils.NewPulsarClient()
	messages, err := client.Subscriptions().GetMessagesByID(*topic, ledgerID, entryID)
	if err != nil {
		return err
	}
	if len(messages) == 0 {
		return fmt.Errorf("no message found with the given ledgerID and entryID")
	}
	message := messages[0]

	propertiesJSON, err := json.Marshal(message.GetProperties())
	if err != nil {
		return err
	}

	//nolint:govet
	oc := cmdutils.NewOutputContent().
		WithObject(&readMessage{
			MessageID:       message.GetMessageID(),
			Properties:      message.GetProperties(),
			Payload:         message.Payload,
			PayloadAsString: string(message.Payload),
		}).
		WithText(`Message ID: %s
		Properties: %s
		Message: 
				%s`, message.GetMessageID(), propertiesJSON, hex.Dump(message.Payload))

	err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)

	return err
}
