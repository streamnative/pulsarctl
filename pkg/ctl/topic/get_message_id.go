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
	"time"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetMessageIDCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get message id by datetime for a topic."
	desc.CommandPermission = "This command requires tenant admin permissions."

	desc.CommandExamples = []cmdutils.Example{
		{
			Desc:    "Get message id by datetime for a topic",
			Command: "pulsarctl topics get-message-id --datetime 2021-06-28T16:53:08Z persistent://public/default/topic",
		},
	}

	vc.SetDescription(
		"get-message-id",
		"Get message id by datetime for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"get-message-id",
	)

	var datetime string
	vc.FlagSetGroup.InFlagSet("GetMessageID", func(set *pflag.FlagSet) {
		set.StringVarP(&datetime, "datetime", "d", "", "datetime in RFC3339 or RFC3339Nano format")
		_ = cobra.MarkFlagRequired(set, "datetime")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doGetMessageID(vc, datetime)
	}, "the topic name is not specified or the topic name is specified more than one")
	vc.EnableOutputFlagSet()
}

func doGetMessageID(vc *cmdutils.VerbCmd, datetime string) error {
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	timestamp, err := time.Parse(time.RFC3339Nano, datetime)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	messageID, err := admin.Topics().GetMessageID(*topic, timestamp.UnixMilli())
	if err == nil {
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), cmdutils.NewOutputContent().WithObject(messageID))
	}
	return err
}
