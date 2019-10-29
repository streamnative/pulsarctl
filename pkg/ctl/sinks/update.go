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

package sinks

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func updateSinksCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Update a Pulsar IO sink connector."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []pulsar.Example

	update := pulsar.Example{
		Desc: "Update a Pulsar IO sink connector",
		Command: "pulsarctl sink update \n" +
			"\t--tenant public \n" +
			"\t--namespace default \n" +
			"\t--name update-source \n" +
			"\t--archive pulsar-io-kafka-2.4.0.nar \n" +
			"\t--classname org.apache.pulsar.io.kafka.KafkaBytesSource \n" +
			"\t--destination-topic-name my-topic \n" +
			"\t--cpu 2",
	}

	updateWithSchema := pulsar.Example{
		Desc: "Update a Pulsar IO sink connector with schema type",
		Command: "pulsarctl sink create \n" +
			"\t--schema-type schema.STRING\n" +
			"\t// Other sink parameters ",
	}

	updateWithParallelism := pulsar.Example{
		Desc: "Update a Pulsar IO sink connector with parallelism",
		Command: "pulsarctl sink create \n" +
			"\t--parallelism 1\n" +
			"\t// Other sink parameters ",
	}

	updateWithResource := pulsar.Example{
		Desc: "Update a Pulsar IO sink connector with resource",
		Command: "pulsarctl sink create \n" +
			"\t--ram 5656565656\n" +
			"\t--disk 8080808080808080\n" +
			"\t--cpu 5.0\n" +
			"\t// Other sink parameters ",
	}

	updateWithSinkConfig := pulsar.Example{
		Desc: "Update a Pulsar IO sink connector with sink config",
		Command: "pulsarctl sink create \n" +
			"\t--sink-config \"{\"publishTopic\":\"publishTopic\", \"key\":\"pulsar\"}\"\n" +
			"\t// Other sink parameters ",
	}

	examples = append(examples, update, updateWithSinkConfig, updateWithResource, updateWithParallelism, updateWithSchema)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Updated (the name of a Pulsar Sink) successfully",
	}

	nameNotExistOut := pulsar.Output{
		Desc: "sink doesn't exist",
		Out:  "code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist",
	}

	out = append(out, successOut, nameNotExistOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"update",
		"Update a Pulsar IO sink connector",
		desc.ToString(),
		desc.ExampleToString(),
		"update",
	)

	sinkData := &pulsar.SinkData{}
	// set the run sink
	vc.SetRunFunc(func() error {
		return doUpdateSink(vc, sinkData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("SinksConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&sinkData.Tenant,
			"tenant",
			"",
			"The sink's tenant")

		flagSet.StringVar(
			&sinkData.Namespace,
			"namespace",
			"",
			"The sink's namespace")

		flagSet.StringVar(
			&sinkData.Name,
			"name",
			"",
			"The sink's name")

		flagSet.StringVarP(
			&sinkData.SinkType,
			"sink-type",
			"t",
			"",
			"The sink's connector provider")

		flagSet.StringVarP(
			&sinkData.Inputs,
			"inputs",
			"i",
			"",
			"The sink's input topic or topics (multiple topics can be specified as a comma-separated list)")

		flagSet.StringVar(
			&sinkData.TopicsPattern,
			"topics-pattern",
			"",
			"TopicsPattern to consume from list of topics under a namespace that match the pattern."+
				" [--input] and [--topicsPattern] are mutually exclusive. Add SerDe class name for a pattern"+
				" in --customSerdeInputs  (supported for java fun only)")

		flagSet.StringVar(
			&sinkData.SubsName,
			"subs-name",
			"",
			"Pulsar source subscription name if user wants a specific subscription-name for input-topic consumer")

		flagSet.StringVar(
			&sinkData.CustomSerdeInputString,
			"custom-serde-inputs",
			"",
			"The map of input topics to SerDe class names (as a JSON string)")

		flagSet.StringVar(
			&sinkData.CustomSchemaInputString,
			"custom-schema-inputs",
			"",
			"The map of input topics to Schema types or class names (as a JSON string)")

		flagSet.StringVar(
			&sinkData.ProcessingGuarantees,
			"processing-guarantees",
			"",
			"The processing guarantees (aka delivery semantics) applied to the sink")

		flagSet.BoolVar(
			&sinkData.RetainOrdering,
			"retain-ordering",
			false,
			"Sink consumes and sinks messages in order")

		flagSet.IntVar(
			&sinkData.Parallelism,
			"parallelism",
			0,
			"The sink's parallelism factor (i.e. the number of sink instances to run)")

		flagSet.StringVar(
			&sinkData.Archive,
			"archive",
			"",
			"Path to the archive file for the sink. It also supports url-path "+
				"[http/https/file (file protocol assumes that file already exists on worker host)] "+
				"from which worker can download the package.")

		flagSet.StringVar(
			&sinkData.ClassName,
			"classname",
			"",
			"The sink's class name if archive is file-url-path (file://)")

		flagSet.StringVar(
			&sinkData.SinkConfigFile,
			"sink-config-file",
			"",
			"The path to a YAML config file specifying the sink's configuration")

		flagSet.Float64Var(
			&sinkData.CPU,
			"cpu",
			0,
			"The CPU (in cores) that needs to be allocated per sink instance"+
				" (applicable only to Docker runtime)")

		flagSet.Int64Var(
			&sinkData.Disk,
			"disk",
			0,
			"The disk (in bytes) that need to be allocated per sink instance"+
				" (applicable only to Docker runtime)")

		flagSet.Int64Var(
			&sinkData.RAM,
			"ram",
			0,
			"The RAM (in bytes) that need to be allocated per sink instance"+
				" (applicable only to the process and Docker runtimes)")

		flagSet.StringVar(
			&sinkData.SinkConfigString,
			"sink-config",
			"",
			"User defined configs key/values")

		flagSet.BoolVar(
			&sinkData.AutoAck,
			"auto-ack",
			false,
			"Whether or not the framework will automatically acknowledge messages")

		flagSet.Int64Var(
			&sinkData.TimeoutMs,
			"timeout-ms",
			0,
			"The message timeout in milliseconds")
	})
}

func doUpdateSink(vc *cmdutils.VerbCmd, sinkData *pulsar.SinkData) error {
	err := processArguments(sinkData)
	if err != nil {
		vc.Command.Help()
		return err
	}

	checkArgsForUpdate(sinkData.SinkConf)

	admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V3)

	updateOptions := pulsar.NewUpdateOptions()
	updateOptions.UpdateAuthData = sinkData.UpdateAuthData

	if utils.IsPackageURLSupported(sinkData.Archive) {
		err = admin.Sinks().UpdateSinkWithURL(sinkData.SinkConf, sinkData.Archive, updateOptions)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Updated instanceID[%s] of Pulsar Sinks[%s] successfully\n", sinkData.InstanceID, sinkData.Name)
		}
	} else {
		err = admin.Sinks().UpdateSink(sinkData.SinkConf, sinkData.Archive, updateOptions)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Updated %s successfully\n", sinkData.Name)
		}
	}

	return err
}
