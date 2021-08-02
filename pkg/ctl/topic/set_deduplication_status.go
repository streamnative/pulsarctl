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
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func SetDeduplicationStatusCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set the deduplication policy for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Set the deduplication policy for a topic",
		Command: "pulsarctl topics set-deduplication topic ",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set the deduplication policy successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-deduplication",
		"Set the deduplication policy for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-deduplication",
	)
	var enable, disable bool
	vc.SetRunFuncWithNameArg(func() error {
		return doSetDeduplicationStatus(vc, enable, disable)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Deduplication", func(set *pflag.FlagSet) {
		set.BoolVarP(
			&enable,
			"enable",
			"e",
			false,
			"Enable deduplication")
		set.BoolVarP(
			&disable,
			"disable",
			"d",
			false,
			"Disable deduplication")
	})
	vc.EnableOutputFlagSet()
}

func doSetDeduplicationStatus(vc *cmdutils.VerbCmd, enable bool, disable bool) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}
	if enable == disable {
		msg := "Need to specify either --enable or --disable"
		vc.Command.Printf(msg)
		return errors.Errorf(msg)
	}
	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetDeduplicationStatus(*topic, enable)
	if err == nil {
		vc.Command.Printf("Set the deduplication policy successfully for [%s]\n", topic.String())
	}
	return err
}
