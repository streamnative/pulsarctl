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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

func updateSinksCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Update a Pulsar IO sink connector."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example

	update := cmdutils.Example{
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

	updateWithSchema := cmdutils.Example{
		Desc: "Update a Pulsar IO sink connector with schema type",
		Command: "pulsarctl sink create \n" +
			"\t--schema-type schema.STRING\n" +
			"\t// Other sink parameters ",
	}

	updateWithParallelism := cmdutils.Example{
		Desc: "Update a Pulsar IO sink connector with parallelism",
		Command: "pulsarctl sink create \n" +
			"\t--parallelism 1\n" +
			"\t// Other sink parameters ",
	}

	updateWithResource := cmdutils.Example{
		Desc: "Update a Pulsar IO sink connector with resource",
		Command: "pulsarctl sink create \n" +
			"\t--ram 5656565656\n" +
			"\t--disk 8080808080808080\n" +
			"\t--cpu 5.0\n" +
			"\t// Other sink parameters ",
	}

	updateWithSinkConfig := cmdutils.Example{
		Desc: "Update a Pulsar IO sink connector with sink config",
		Command: "pulsarctl sink create \n" +
			"\t--sink-config \"{\"publishTopic\":\"publishTopic\", \"key\":\"pulsar\"}\"\n" +
			"\t// Other sink parameters ",
	}

	examples = append(examples, update, updateWithSinkConfig, updateWithResource, updateWithParallelism, updateWithSchema)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Updated (the name of a Pulsar Sink) successfully",
	}

	nameNotExistOut := cmdutils.Output{
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

	sinkData := &util.SinkData{}
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

		flagSet.BoolVar(
			&sinkData.CleanupSubscription,
			"cleanup-subscription",
			true,
			"Whether delete the subscription when sink is deleted")

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
			&sinkData.SubsPosition,
			"subs-position",
			"",
			"Pulsar source subscription position if user wants to consume messages from the specified location. "+
				"Possible Values: [Latest, Earliest]")

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
			&sinkData.InputSpecs,
			"input-specs",
			"",
			"The map of inputs to custom configuration (as a JSON string)")

		flagSet.IntVar(
			&sinkData.MaxMessageRetries,
			"max-redeliver-count",
			0,
			"Maximum number of times that a message will be redelivered before being sent to the dead letter queue")

		flagSet.StringVar(
			&sinkData.DeadLetterTopic,
			"dead-letter-topic",
			"",
			"Name of the dead topic where the failing messages will be sent.")

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

		flagSet.BoolVar(
			&sinkData.RetainKeyOrdering,
			"retain-key-ordering",
			false,
			"Sink consumes and processes messages in key order")

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

		flagSet.Int64Var(
			&sinkData.NegativeAckRedeliveryDelayMs,
			"negative-ack-redelivery-delay-ms",
			0,
			"The negative ack message redelivery delay in milliseconds")

		flagSet.StringVar(
			&sinkData.CustomRuntimeOptions,
			"custom-runtime-options",
			"",
			"A string that encodes options to customize the runtime, see docs for configured runtime for details")

		flagSet.StringVar(
			&sinkData.Secrets,
			"secrets",
			"",
			"The map of secretName to an object that encapsulates how the secret is fetched by the underlying secrets"+
				"provider")

		flagSet.StringVar(
			&sinkData.TransformFunction,
			"transform-function",
			"",
			"Transform function applied before the Sink")

		flagSet.StringVar(
			&sinkData.TransformFunctionClassName,
			"transform-function-classname",
			"",
			"The transform function class name")

		flagSet.StringVar(
			&sinkData.TransformFunctionConfig,
			"transform-function-config",
			"",
			"Configuration of the transform function applied before the Sink")

		flagSet.BoolVar(
			&sinkData.UpdateAuthData,
			"update-auth-data",
			false,
			"Whether or not to update the auth data")

		_ = flagSet.MarkDeprecated("auto-ack", "this value is immutable")
		_ = flagSet.MarkDeprecated("processing-guarantees", "this value is immutable")
		_ = flagSet.MarkDeprecated("retain-ordering", "this value is immutable")
	})
	vc.EnableOutputFlagSet()
}

func doUpdateSink(vc *cmdutils.VerbCmd, sinkData *util.SinkData) error {
	err := processArguments(sinkData)
	if err != nil {
		_ = vc.Command.Help()
		return err
	}

	checkArgsForUpdate(sinkData.SinkConf)

	// convert the map[interface{}]interface{} to a map[string]interface{} for unmarshal
	for k, v := range sinkData.SinkConf.Secrets {
		sinkData.SinkConf.Secrets[k] = utils.ConvertMap(v)
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)

	latestConfig, err := admin.Sinks().GetSink(sinkData.Tenant, sinkData.Namespace, sinkData.Name)
	if err != nil {
		return err
	}

	sinkData.SinkConf.AutoAck = latestConfig.AutoAck
	sinkData.SinkConf.RetainOrdering = latestConfig.RetainOrdering
	sinkData.SinkConf.ProcessingGuarantees = latestConfig.ProcessingGuarantees

	updateOptions := util.NewUpdateOptions()
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
