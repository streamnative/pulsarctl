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

package functions

import (
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

func createFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for creating a new Pulsar Function in cluster mode."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with jar file",
		Command: "pulsarctl functions create \n" +
			"\t--tenant public \n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Functions>) \n" +
			"\t--inputs test-input-topic  \n" +
			"\t--output persistent://public/default/test-output-topic \n" +
			"\t--classname org.apache.pulsar.functions.api.examples.ExclamationFunction \n" +
			"\t--jar /examples/api-examples.jar",
	}
	examples = append(examples, create)

	createWithConf := cmdutils.Example{
		Desc: "Create a Pulsar Function use function config yaml file",
		Command: "pulsarctl functions create \n" +
			"\t--function-config-file (the path of function config yaml file) \n" +
			"\t--jar (the path of user code jar)",
	}
	examples = append(examples, createWithConf)

	createWithPkgURL := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with pkg URL",
		Command: "pulsarctl functions create \n" +
			"\t--tenant public \n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Functions) \n" +
			"\t--inputs test-input-topic  \n" +
			"\t--output persistent://public/default/test-output-topic \n" +
			"\t--classname org.apache.pulsar.functions.api.examples.ExclamationFunction \n" +
			"\t--jar file:/http: + /examples/api-examples.jar",
	}
	examples = append(examples, createWithPkgURL)

	createWithLogTopic := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with log topic",
		Command: "pulsarctl functions create \n" +
			"\t--log-topic persistent://public/default/test-log-topic\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithLogTopic)

	createWithDeadLetterTopic := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with dead letter topic",
		Command: "pulsarctl functions create \n" +
			"\t--dead-letter-topic persistent://public/default/test-dead-letter-topic\n" +
			"\t--max-message-retries 10\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithDeadLetterTopic)

	createWithAutoAck := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with auto ack",
		Command: "pulsarctl functions create \n" +
			"\t--auto-ack \n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithAutoAck)

	createWithFQFN := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with FQFN",
		Command: "pulsarctl functions create \n" +
			"\t--fqfn tenant/namespace/name eg:public/default/test-fqfn-function\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithFQFN)

	createWithTopicsPattern := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with topics pattern",
		Command: "pulsarctl functions create \n" +
			"\t--topics-pattern persistent://tenant/ns/topicPattern*\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithTopicsPattern)

	createWithUserConfig := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with user config",
		Command: "pulsarctl functions create \n" +
			"\t--user-config \"{\"publishTopic\":\"publishTopic\", \"key\":\"pulsar\"}\"\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithUserConfig)

	createWithRetainOrdering := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with retain ordering",
		Command: "pulsarctl functions create \n" +
			"\t--retain-ordering \n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithRetainOrdering)

	createWithCustomSchemasInputs := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with custom schema for inputs topic",
		Command: "pulsarctl functions create \n" +
			"\t--custom-schema-inputs \"{\"topic-1\":\"schema.STRING\", \"topic-2\":\"schema.JSON\"}\"\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithCustomSchemasInputs)

	createWithSchema := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with schema type for output topic",
		Command: "pulsarctl functions create \n" +
			"\t--schema-type schema.STRING\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithSchema)

	createWithParallelism := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with parallelism",
		Command: "pulsarctl functions create \n" +
			"\t--parallelism 1\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithParallelism)

	createWithResource := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with resource",
		Command: "pulsarctl functions create \n" +
			"\t--ram 5656565656\n" +
			"\t--disk 8080808080808080\n" +
			"\t--cpu 5.0\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithResource)

	createWithWindowFunctions := cmdutils.Example{
		Desc: "Create a Pulsar Function in cluster mode with window functions",
		Command: "pulsarctl functions create \n" +
			"\t--window-length-count 10\n" +
			"\t--window-length-duration-ms 1000\n" +
			"\t--sliding-interval-count 3\n" +
			"\t--sliding-interval-duration-ms 1000\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithWindowFunctions)

	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Created (the name of a Pulsar Function) successfully",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Create a Pulsar Function to run on a Pulsar cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"create",
	)

	functionData := &util.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doCreateFunctions(vc, functionData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("FunctionsConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&functionData.FQFN,
			"fqfn",
			"",
			"The Fully Qualified Function Name (FQFN) for the function")

		flagSet.StringVar(
			&functionData.Tenant,
			"tenant",
			"",
			"The tenant of a Pulsar Function")

		flagSet.StringVar(
			&functionData.Namespace,
			"namespace",
			"",
			"The namespace of a Pulsar Function")

		flagSet.StringVar(
			&functionData.FuncName,
			"name",
			"",
			"The name of a Pulsar Function")

		flagSet.StringVar(
			&functionData.ClassName,
			"classname",
			"",
			"The class name of a Pulsar Function")

		flagSet.StringVarP(
			&functionData.FunctionType,
			"function-type",
			"t",
			"",
			"The built-in Pulsar Function type")

		flagSet.BoolVar(
			&functionData.CleanupSubscription,
			"cleanup-subscription",
			true,
			"Whether delete the subscription when function is deleted")

		flagSet.StringVar(
			&functionData.Jar,
			"jar",
			"",
			"Path to the JAR file for the function (if the function is written in Java) "+
				"It also supports URL path [http/https/file (file protocol assumes that file "+
				"already exists on worker host)] from which worker can download the package.")

		flagSet.StringVar(
			&functionData.Py,
			"py",
			"",
			"Path to the main Python file/Python Wheel file for the function (if the function is written in Python) "+
				"It also supports URL path [http/https/file (file protocol assumes that file "+
				"already exists on worker host)] from which worker can download the package.")

		flagSet.StringVar(
			&functionData.Go,
			"go",
			"",
			"Path to the main Go executable binary for the function (if the function is written in Go) "+
				"It also supports URL path [http/https/file (file protocol assumes that file "+
				"already exists on worker host)] from which worker can download the package.")

		flagSet.StringVarP(
			&functionData.Inputs,
			"inputs",
			"i",
			"",
			"The input topic or topics (multiple topics can be specified as a comma-separated list) of a Pulsar Function")

		flagSet.StringVar(
			&functionData.TopicsPattern,
			"topics-pattern",
			"",
			"The topic pattern to consume from list of topics under a namespace "+
				"that match the pattern. [--input] and [--topic-pattern] are mutually exclusive. "+
				"Add SerDe class name for a pattern in --custom-serde-inputs (supported for java fun only)")

		flagSet.StringVarP(
			&functionData.Output,
			"output",
			"o",
			"",
			"The output topic of a Pulsar Function (If none is specified, no output is written)")

		flagSet.StringVar(
			&functionData.ProducerConfig,
			"producer-config",
			"",
			"The custom producer configuration (as a JSON string)")

		flagSet.StringVar(
			&functionData.LogTopic,
			"log-topic",
			"",
			"The topic to which the logs of a Pulsar Function are produced")

		flagSet.StringVar(
			&functionData.SchemaType,
			"schema-type",
			"",
			"The builtin schema type or custom schema class name to be used for messages output by the function")

		flagSet.StringVar(
			&functionData.CustomSerDeInputs,
			"custom-serde-inputs",
			"",
			"The map of input topics to SerDe class names (as a JSON string)")

		flagSet.StringVar(
			&functionData.CustomSchemaInput,
			"custom-schema-inputs",
			"",
			"The map of input topics to Schema class names (as a JSON string)")

		flagSet.StringVar(
			&functionData.CustomSchemaOutput,
			"custom-schema-outputs",
			"",
			"The map of input topics to Schema properties (as a JSON string)")

		flagSet.StringVar(
			&functionData.InputSpecs,
			"input-specs",
			"",
			"The map of inputs to custom configuration (as a JSON string)")

		flagSet.StringVar(
			&functionData.InputTypeClassName,
			"input-type-class-name",
			"",
			"The class name of input type class")

		flagSet.StringVar(
			&functionData.OutputSerDeClassName,
			"output-serde-classname",
			"",
			"The SerDe class to be used for messages output by the function")

		flagSet.StringVar(
			&functionData.OutputTypeClassName,
			"output-type-class-name",
			"",
			"The class name of output type class")

		flagSet.StringVar(
			&functionData.FunctionConfigFile,
			"function-config-file",
			"",
			"The path to a YAML config file that specifies the configuration of a Pulsar Function")

		flagSet.StringVar(
			&functionData.ProcessingGuarantees,
			"processing-guarantees",
			"",
			"The processing guarantees (aka delivery semantics) applied to the function")

		flagSet.StringVar(
			&functionData.UserConfig,
			"user-config",
			"",
			"User-defined config key/values")

		flagSet.BoolVar(
			&functionData.RetainOrdering,
			"retain-ordering",
			false,
			"Function consumes and processes messages in order")

		flagSet.BoolVar(
			&functionData.RetainKeyOrdering,
			"retain-key-ordering",
			false,
			"Function consumes and processes messages in key order")

		flagSet.StringVar(
			&functionData.BatchBuilder,
			"batch-builder",
			"",
			"BatcherBuilder provides two types of batch construction methods, DEFAULT and KEY_BASED."+
				"The default value is: DEFAULT")

		flagSet.BoolVar(
			&functionData.ForwardSourceMessageProperty,
			"forward-source-message-property",
			true,
			"Forwarding input message's properties to output topic when processing (use false to disable it)")

		flagSet.StringVar(
			&functionData.SubsName,
			"subs-name",
			"",
			"Pulsar source subscription name if user wants a specific subscription-name for input-topic consumer")

		flagSet.StringVar(
			&functionData.SubsPosition,
			"subs-position",
			"",
			"Pulsar source subscription position if user wants to consume messages from the specified location")

		flagSet.BoolVar(
			&functionData.SkipToLatest,
			"skip-to-latest",
			false,
			"Whether or not the consumer skip to latest message upon function instance restart")

		flagSet.IntVar(
			&functionData.Parallelism,
			"parallelism",
			0,
			"The parallelism factor of a Pulsar Function (i.e. the number of function instances to run)")

		flagSet.Float64Var(
			&functionData.CPU,
			"cpu",
			0,
			"The cpu in cores that need to be allocated per function instance(applicable only to docker runtime)")

		flagSet.Int64Var(
			&functionData.RAM,
			"ram",
			0,
			"The ram in bytes that need to be allocated per function instance"+
				"(applicable only to process/docker runtime)")

		flagSet.Int64Var(
			&functionData.Disk,
			"disk",
			0,
			"The disk in bytes that need to be allocated per function instance"+
				"(applicable only to docker runtime)")

		flagSet.IntVar(
			&functionData.WindowLengthCount,
			"window-length-count",
			0,
			"The number of messages per window")

		flagSet.Int64Var(
			&functionData.WindowLengthDurationMs,
			"window-length-duration-ms",
			0,
			"The time duration of the window in milliseconds")

		flagSet.IntVar(
			&functionData.SlidingIntervalCount,
			"sliding-interval-count",
			0,
			"The number of messages after which the window slides")

		flagSet.Int64Var(
			&functionData.SlidingIntervalDurationMs,
			"sliding-interval-duration-ms",
			0,
			"The time duration after which the window slides")

		flagSet.BoolVar(
			&functionData.AutoAck,
			"auto-ack",
			true,
			"Whether or not the framework acknowledges messages automatically")

		flagSet.Int64Var(
			&functionData.TimeoutMs,
			"timeout-ms",
			0,
			"The message timeout in milliseconds")

		flagSet.IntVar(
			&functionData.MaxMessageRetries,
			"max-message-retries",
			0,
			"How many times should we try to process a message before giving up")

		flagSet.StringVar(
			&functionData.CustomRuntimeOptions,
			"custom-runtime-options",
			"",
			"A string that encodes options to customize the runtime, see docs for configured runtime for details #Java")

		flagSet.StringVar(
			&functionData.Secrets,
			"secrets",
			"",
			"The map of secretName to an object that encapsulates how the secret is fetched by the underlying secrets provider")

		flagSet.StringVar(
			&functionData.DeadLetterTopic,
			"dead-letter-topic",
			"",
			"The topic where messages that are not processed successfully are sent to")
	})
	vc.EnableOutputFlagSet()
}

func doCreateFunctions(vc *cmdutils.VerbCmd, funcData *util.FunctionData) error {
	err := processArgs(funcData)
	if err != nil {
		_ = vc.Command.Help()
		return err
	}

	err = validateFunctionConfigs(funcData.FuncConf)
	if err != nil {
		_ = vc.Command.Help()
		return err
	}

	formatFuncConf(funcData.FuncConf)

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)

	if utils.IsPackageURLSupported(funcData.UserCodeFile) {
		err = admin.Functions().CreateFuncWithURL(funcData.FuncConf, funcData.UserCodeFile)
	} else {
		err = admin.Functions().CreateFunc(funcData.FuncConf, funcData.UserCodeFile)
	}

	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Created %s successfully\n", funcData.FuncName)
	}
	return err
}
