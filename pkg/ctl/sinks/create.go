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
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	util "github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func createSinksCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Create a Pulsar IO sink connector to run in a Pulsar cluster."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc: "Create a Pulsar Sink in cluster mode",
		Command: "pulsarctl sink create \n" +
			"\t--tenant public \n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Sink) \n" +
			"\t--inputs test-jdbc  \n" +
			"\t--archive connectors/pulsar-io-jdbc-2.4.0.nar \n" +
			"\t--sink-config-file connectors/mysql-jdbc-sink.yaml \n" +
			"\t--parallelism 1",
	}

	createWithPkgURL := cmdutils.Example{
		Desc: "Create a Pulsar Sink in cluster mode with pkg URL",
		Command: "pulsarctl sink create \n" +
			"\t--tenant public \n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Sink) \n" +
			"\t--inputs test-jdbc  \n" +
			"\t--archive file:/http: + connectors/pulsar-io-jdbc-2.4.0.nar",
	}

	createWithSchema := cmdutils.Example{
		Desc: "Create a Pulsar Sink in cluster mode with schema type",
		Command: "pulsarctl sink create \n" +
			"\t--schema-type schema.STRING\n" +
			"\t// Other sink parameters ",
	}

	createWithParallelism := cmdutils.Example{
		Desc: "Create a Pulsar Sink in cluster mode with parallelism",
		Command: "pulsarctl sink create \n" +
			"\t--parallelism 1\n" +
			"\t// Other sink parameters ",
	}

	createWithResource := cmdutils.Example{
		Desc: "Create a Pulsar Sink in cluster mode with resource",
		Command: "pulsarctl sink create \n" +
			"\t--ram 5656565656\n" +
			"\t--disk 8080808080808080\n" +
			"\t--cpu 5.0\n" +
			"\t// Other sink parameters ",
	}

	createWithSinkConfig := cmdutils.Example{
		Desc: "Create a Pulsar Sink in cluster mode with sink config",
		Command: "pulsarctl sink create \n" +
			"\t--sink-config \"{\"publishTopic\":\"publishTopic\", \"key\":\"pulsar\"}\"\n" +
			"\t// Other sink parameters ",
	}

	createWithProcessingGuarantees := cmdutils.Example{
		Desc: "Create a Pulsar Sink in cluster mode with processing guarantees",
		Command: "pulsarctl sink create \n" +
			"\t--processing-guarantees EFFECTIVELY_ONCE\n" +
			"\t// Other sink parameters ",
	}

	examples = append(examples, create, createWithPkgURL, createWithSchema, createWithParallelism,
		createWithResource, createWithSinkConfig, createWithProcessingGuarantees)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Created (the name of a Pulsar Sinks) successfully",
	}

	failureOut := cmdutils.Output{
		Desc: "sink archive not specified, please check --archive arg",
		Out:  "[✖]  Sink archive not specified",
	}

	sinkTypeOut := cmdutils.Output{
		Desc: "Cannot specify both archive and sink-type, please check --archive and --sink-type args",
		Out:  "[✖]  Cannot specify both archive and sink-type",
	}

	out = append(out, successOut, failureOut, sinkTypeOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Create a Pulsar IO sink connector to run in a Pulsar cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"create",
	)

	sinkData := &util.SinkData{}

	// set the run sink
	vc.SetRunFunc(func() error {
		return doCreateSinks(vc, sinkData)
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
			"TopicsPattern to consume from list of topics under a namespace that match the pattern. "+
				"[--input] and [--topicsPattern] are mutually exclusive. Add SerDe class name for a pattern "+
				"in --customSerdeInputs  (supported for java fun only)")

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
			0.0,
			"The CPU (in cores) that needs to be allocated per sink instance (applicable only to Docker runtime)")

		flagSet.Int64Var(
			&sinkData.Disk,
			"disk",
			0,
			"The disk (in bytes) that need to be allocated per sink instance (applicable only to Docker runtime)")

		flagSet.Int64Var(
			&sinkData.RAM,
			"ram",
			0,
			"The RAM (in bytes) that need to be allocated per sink instance "+
				"(applicable only to the process and Docker runtimes)")

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

func doCreateSinks(vc *cmdutils.VerbCmd, sinkData *util.SinkData) error {
	err := processArguments(sinkData)
	if err != nil {
		vc.Command.Help()
		return err
	}

	err = validateSinkConfigs(sinkData.SinkConf)
	if err != nil {
		vc.Command.Help()
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	if utils.IsPackageURLSupported(sinkData.Archive) {
		err = admin.Sinks().CreateSinkWithURL(sinkData.SinkConf, sinkData.Archive)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Created instanceID[%s] of Pulsar Sink[%s] successfully\n", sinkData.InstanceID, sinkData.Name)
		}
	} else {
		err = admin.Sinks().CreateSink(sinkData.SinkConf, sinkData.Archive)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Created %s successfully\n", sinkData.Name)
		}
	}

	return err
}
