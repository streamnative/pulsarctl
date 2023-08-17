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
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func SetPersistenceCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set persistence for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Set persistence for a topic",
		Command: "pulsarctl topics set-persistence topic -e 4 -w 3 -a 2 -r 1.0",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set persistence successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-persistence",
		"Set persistence for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-persistence",
	)
	persistenceData := &utils.PersistenceData{}
	vc.SetRunFuncWithNameArg(func() error {
		return doSetPersistence(vc, persistenceData)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Persistence", func(set *pflag.FlagSet) {
		set.Int64VarP(
			&persistenceData.BookkeeperEnsemble,
			"bookkeeper-ensemble",
			"e",
			0,
			"Number of bookies to use for a topic")
		set.Int64VarP(
			&persistenceData.BookkeeperWriteQuorum,
			"bookkeeper-write-quorum",
			"w",
			0,
			"How many writes to make of each entry")
		set.Int64VarP(
			&persistenceData.BookkeeperAckQuorum,
			"bookkeeper-ack-quorum",
			"a",
			0,
			"Number of acks (guaranteed copies) to wait for each entry")
		set.Float64VarP(
			&persistenceData.ManagedLedgerMaxMarkDeleteRate,
			"ml-mark-delete-max-rate",
			"r",
			0.0,
			"Throttling rate of mark-delete operation (0 means no throttle)")
	})
	vc.EnableOutputFlagSet()
}

func doSetPersistence(vc *cmdutils.VerbCmd, persistenceData *utils.PersistenceData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetPersistence(*topic, *persistenceData)
	if err == nil {
		vc.Command.Printf("Set persistence successfully for [%s]\n", topic.String())
	}
	return err
}
