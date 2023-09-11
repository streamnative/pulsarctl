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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetCompactionThresholdCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Get the compaction threshold for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Get the compaction threshold for a topic",
		Command: "pulsarctl topics get-compaction-threshold [topic]",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The compaction threshold of the topic (topic-name) is (size) byte(s)",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-compaction-threshold",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString())

	var applied bool
	vc.FlagSetGroup.InFlagSet("Get Compaction Threshold", func(flagSet *pflag.FlagSet) {
		flagSet.BoolVarP(&applied, "applied", "", false, "Get the applied policy for the topic")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doGetCompactionThreshold(vc, applied)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doGetCompactionThreshold(vc *cmdutils.VerbCmd, applied bool) error {
	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	threshold, err := admin.Topics().GetCompactionThreshold(*topic, applied)
	if err == nil {
		vc.Command.Printf("The compaction threshold of the topic %s is %d byte(s)", topic, threshold)
	}

	return err
}
