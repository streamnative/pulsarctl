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

package cat

import (
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cat"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/thediveo/enumflag"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func CatReaderCommand(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for reading a topic and writing the contents " +
		"to stdout or a file. It includes options for formatting as well as how many messages to read. "
	desc.CommandPermission = "This command requires permissions to read the topic."

	var examples []cmdutils.Example
	basic := cmdutils.Example{
		Desc:    "Reads a topic to stdout with metadata",
		Command: "pulsarctl cat reader (topic)",
	}
	asString := cmdutils.Example{
		Desc:    "Reads a topic to stdout with string only",
		Command: "pulsarctl cat reader (topic) --format payload-string",
	}
	examples = append(examples, basic, asString)
	desc.CommandExamples = examples

	// update the descriptiogo build reader.gon
	vc.SetDescription(
		"reader",
		"Read a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"read")

	readerOpts := &cat.ReaderOpts{}

	var ID string
	var output string
	var outputStats bool
	// set the run function with name argument
	vc.SetRunFuncWithNameArg(func() error {
		return doCatReader(vc, readerOpts, ID, output, outputStats)
	}, "the topic name is not specified or the topic name is specified more than once")

	// register the params
	vc.FlagSetGroup.InFlagSet("ReaderOpts", func(flagSet *pflag.FlagSet) {
		flagSet.IntVar(
			&readerOpts.MsgCount,
			"message-count",
			-1,
			"The number of messages to read, defaults to -1, which will read all messages. Tailing overrides this property")
		flagSet.BoolVarP(
			&readerOpts.Tailing,
			"tail",
			"t",
			false,
			"Indicates to tail the topic, by default, stop reading once end of topic is reached")
		flagSet.StringVarP(&ID, "messageId", "m", "latest",
			"Message id where the reader starts from. It can be either 'latest', "+
				"'earliest' or (ledgerId:entryId)")
		flagSet.VarP(enumflag.New(
			&readerOpts.Format,
			"format",
			cat.ReaderFormatIds,true),
			"format",
			"f",
			"The mode to output, can be payload-string, payload-hex, payload-base64 and full-message")
		flagSet.BoolVar(
			&outputStats,
			"stats",
			false,
			"Periodically log stats")
		flagSet.BoolVar(
			&readerOpts.Compacted,
			"compacted",
			false,
			"Indicates to read the compacted topic")
		flagSet.StringVarP(
			&output,
			"output",
		"o",
		"-",
			"The file (or - for stdout) to output data")
	})
}

func doCatReader(vc *cmdutils.VerbCmd, readerOpts *cat.ReaderOpts, msgId, output string, stats bool) error {
	readerOpts.Topic = vc.NameArg

	var messageID pulsar.MessageID
	var err error
	switch msgId {
	case "latest":
		messageID = pulsar.LatestMessageID()
	case "earliest":
		messageID = pulsar.EarliestMessageID()
	default:
		messageID, err = pulsar.MessageIDFromString(msgId)
		if err != nil {
			return err
		}
	}
	readerOpts.StartAt = messageID
	if readerOpts.Tailing {
		readerOpts.MsgCount = -1
	}

	pulsarClient, err := cmdutils.NewBrokerClient()
	if err != nil {
		return err
	}
	reader, err := cat.NewReader(pulsarClient, *readerOpts)
	if err != nil {
		return err
	}

	var outputWriter io.Writer
	if output == "-" {
		outputWriter = os.Stdout
	} else {
		outFile, err := os.Create(output)
		if err != nil {
			return err
		}
		defer outFile.Close()
		outputWriter = outFile
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		reader.Close()
	}()

	if stats {
		go func() {
			for _ = range time.Tick(time.Second * 10) {
				fmt.Println(reader.ReadStats())
			}
		}()
	}

	bytes, err := io.Copy(outputWriter, reader)
	fmt.Printf("Read %v bytes total\n", bytes)
	return err
}

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"cat",
		"Read a topic",
		"",
		"cat")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, CatReaderCommand)

	return resourceCmd
}
