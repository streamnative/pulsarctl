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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"

	util "github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/pkg/errors"
)

func OffloadCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for triggering offloading the messages of a topic " +
		"to a long-term storage (e.g. Amazon S3)"
	desc.CommandPermission = "This command requires tenant admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic"

	var examples []cmdutils.Example
	offload := cmdutils.Example{
		Desc: "Trigger offloading the messages of a topic (topic-name) to a long-term storage and " +
			"keep the configured amount of data in BookKeeper only (e.g. 10M, 5G, " +
			"the unit is byte if the specified value without the unit.)",
		Command: "pulsarctl topic offload (topic-name) (threshold)",
	}
	examples = append(examples, offload)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Trigger offloading the data before the message (messageId) of the topic (topic-name) successfully",
	}

	nothingOut := cmdutils.Output{
		Desc: "noting to offload",
		Out:  "Nothing to offload",
	}

	argsError := cmdutils.Output{
		Desc: "the topic name is not specified or the offload threshold is not specified",
		Out:  "[✖]  only two arguments are allowed to be used as names",
	}
	out = append(out, successOut, nothingOut, argsError, TopicNotFoundError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"offload",
		"Offload the messages of a topic to a long-term storage",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doOffload(vc)
	}, func(args []string) error {
		if len(args) != 2 {
			return errors.New("only two arguments are allowed to be used as names")
		}
		return nil
	})

}

func doOffload(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := util.GetTopicName(vc.NameArgs[0])
	if err != nil {
		return err
	}

	size, err := utils.ValidateSizeString(vc.NameArgs[1])
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()

	stats, err := admin.Topics().GetInternalStats(*topic)
	if err != nil {
		return err
	}

	if len(stats.Ledgers) < 1 {
		return errors.New("topic doesn't have any data")
	}

	ledgers := stats.Ledgers
	ledgers[len(ledgers)-1].Size = stats.CurrentLedgerSize
	messageID := findFirstLedgerWithinThreshold(ledgers, size)
	if messageID == nil {
		vc.Command.Printf("Nothing to offload\n")
		return nil
	}

	err = admin.Topics().Offload(*topic, *messageID)
	if err == nil {
		vc.Command.Printf("Trigger offloading the data before the message %+v of the topic %s successfully\n",
			messageID, topic.String())
	}

	return err
}

func findFirstLedgerWithinThreshold(ledgers []util.LedgerInfo, sizeThreshold int64) *util.MessageID {
	var suffixSize int64
	previousLedger := ledgers[len(ledgers)-1].LedgerID
	for i := len(ledgers) - 1; i >= 0; i-- {
		suffixSize += ledgers[i].Size
		if suffixSize > sizeThreshold {
			return &util.MessageID{
				LedgerID:         previousLedger,
				EntryID:          0,
				PartitionedIndex: -1,
			}
		}
		previousLedger = ledgers[i].LedgerID
	}
	return nil
}
