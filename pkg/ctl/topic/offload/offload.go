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

package offload

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/args"
	e "github.com/streamnative/pulsarctl/pkg/ctl/topic/errors"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
)

func OffloadCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for triggering offloading the data from a topic " +
		"to long-term storage (e.g. Amazon S3)"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	offload := pulsar.Example{
		Desc: "Trigger offloading the data from a topic <topic-name> to a long-term storage and " +
			"keep the configured amount of data in BookKeeper only (e.g. 10M, 5G, default is byte)",
		Command: "pulsarctl topic offload <topic-name> <threshold>",
	}
	examples = append(examples, offload)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Offload trigger for <topic-name> for messages before <message-id>",
	}

	nothingOut := pulsar.Output{
		Desc: "noting to offload",
		Out:  "Nothing to offload",
	}

	argsError := pulsar.Output{
		Desc: "the topic name is not specified or the offload threshold is not specified",
		Out:  "[✖]  only two argument is allowed to be used as names",
	}
	out = append(out, successOut, nothingOut, argsError, e.TopicNotFoundError)
	out = append(out, e.TopicNameErrors...)
	out = append(out, e.NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"offload",
		"Offload the data form a topic to a long-term storage",
		desc.ToString())

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doOffload(vc)
	}, args.CheckTopicNameTwoArgs)

}

func doOffload(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := pulsar.GetTopicName(vc.NameArgs[0])
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
		return errors.New("Topic doesn't have any data.")
	}

	ledgers := stats.Ledgers
	ledgers[len(ledgers)-1].Size = stats.CurrentLedgerSize
	messageId := findFirstLedgerWithinThreshold(ledgers, size)
	if err == nil {
		vc.Command.Printf("Nothing to offload")
		return nil
	}

	err = admin.Topics().Offload(*topic, *messageId)
	if err == nil {
		vc.Command.Printf("Trigger offloading the data before the message %+v of the topic %s successfully",
			messageId, topic.String())
	}

	return err
}

func findFirstLedgerWithinThreshold(ledgers []pulsar.LedgerInfo, sizeThreshold int64) *pulsar.MessageId {
	var suffixSize int64
	previousLedger := ledgers[len(ledgers)-1].LedgerId
	for i := len(ledgers) - 1; i >= 0; i-- {
		suffixSize += ledgers[i].Size
		if suffixSize > sizeThreshold {
			return &pulsar.MessageId{
				LedgerId:         previousLedger,
				EntryId:          0,
				PartitionedIndex: -1,
			}
		}
		previousLedger = ledgers[i].LedgerId
	}
	return nil
}
