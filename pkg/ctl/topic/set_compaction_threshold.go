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
	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

func SetCompactionThresholdCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Set the compaction threshold for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the compaction threshold for a topic",
		Command: "pulsarctl topics set-compaction-threshold topic --threshold 3T",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set compaction threshold to (size) for topic (topic-name)",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-compaction-threshold",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString())

	var threshold string
	vc.FlagSetGroup.InFlagSet("Set Compaction Threshold", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&threshold,
			"threshold",
			"t",
			"0",
			"Maximum number of bytes in a topic backlog before compaction is triggered (eg: 10M, 16G, 3T). "+
				"0 disables automatic compaction")
		_ = cobra.MarkFlagRequired(flagSet, "threshold")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doSetCompactionThreshold(vc, threshold)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doSetCompactionThreshold(vc *cmdutils.VerbCmd, threshold string) error {
	topic, err := util.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	size, err := utils.ValidateSizeString(threshold)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetCompactionThreshold(*topic, size)
	if err == nil {
		vc.Command.Printf("Successfully set compaction threshold to %d for topic %s", size, topic)
	}

	return err
}
