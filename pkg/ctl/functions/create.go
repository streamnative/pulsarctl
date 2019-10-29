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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func createFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "This command is used for creating a new Pulsar Function in cluster mode."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	create := pulsar.Example{
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

	createWithConf := pulsar.Example{
		Desc: "Create a Pulsar Function use function config yaml file",
		Command: "pulsarctl functions create \n" +
			"\t--function-config-file (the path of function config yaml file) \n" +
			"\t--jar (the path of user code jar)",
	}
	examples = append(examples, createWithConf)

	createWithPkgURL := pulsar.Example{
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

	createWithLogTopic := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with log topic",
		Command: "pulsarctl functions create \n" +
			"\t--log-topic persistent://public/default/test-log-topic\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithLogTopic)

	createWithDeadLetterTopic := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with dead letter topic",
		Command: "pulsarctl functions create \n" +
			"\t--dead-letter-topic persistent://public/default/test-dead-letter-topic\n" +
			"\t--max-message-retries 10\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithDeadLetterTopic)

	createWithAutoAck := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with auto ack",
		Command: "pulsarctl functions create \n" +
			"\t--auto-ack \n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithAutoAck)

	createWithFQFN := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with FQFN",
		Command: "pulsarctl functions create \n" +
			"\t--fqfn tenant/namespace/name eg:public/default/test-fqfn-function\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithFQFN)

	createWithTopicsPattern := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with topics pattern",
		Command: "pulsarctl functions create \n" +
			"\t--topics-pattern persistent://tenant/ns/topicPattern*\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithTopicsPattern)

	createWithUserConfig := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with user config",
		Command: "pulsarctl functions create \n" +
			"\t--user-config \"{\"publishTopic\":\"publishTopic\", \"key\":\"pulsar\"}\"\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithUserConfig)

	createWithRetainOrdering := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with retain ordering",
		Command: "pulsarctl functions create \n" +
			"\t--retain-ordering \n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithRetainOrdering)

	createWithCustomSchemasInputs := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with custom schema for inputs topic",
		Command: "pulsarctl functions create \n" +
			"\t--custom-schema-inputs \"{\"topic-1\":\"schema.STRING\", \"topic-2\":\"schema.JSON\"}\"\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithCustomSchemasInputs)

	createWithSchema := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with schema type for output topic",
		Command: "pulsarctl functions create \n" +
			"\t--schema-type schema.STRING\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithSchema)

	createWithParallelism := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with parallelism",
		Command: "pulsarctl functions create \n" +
			"\t--parallelism 1\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithParallelism)

	createWithResource := pulsar.Example{
		Desc: "Create a Pulsar Function in cluster mode with resource",
		Command: "pulsarctl functions create \n" +
			"\t--ram 5656565656\n" +
			"\t--disk 8080808080808080\n" +
			"\t--cpu 5.0\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, createWithResource)

	createWithWindowFunctions := pulsar.Example{
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

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Created (the name of a Pulsar Function) successfully",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"",
		desc.ToString(),
		desc.ExampleToString(),
		"create",
	)

	functionData := &pulsar.FunctionData{}

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
			"Path to the main Python file/Python Wheel file for the function (if the function is written in Python)")

		flagSet.StringVar(
			&functionData.Go,
			"go",
			"",
			"Path to the main Go executable binary for the function (if the function is written in Go)")

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
			&functionData.LogTopic,
			"log-topic",
			"",
			"The topic to which the logs of a Pulsar Function are produced")

		flagSet.StringVarP(
			&functionData.SchemaType,
			"schema-type",
			"t",
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
			&functionData.OutputSerDeClassName,
			"output-serde-classname",
			"",
			"The SerDe class to be used for messages output by the function")

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

		flagSet.StringVar(
			&functionData.SubsName,
			"subs-name",
			"",
			"Pulsar source subscription name if user wants a specific subscription-name for input-topic consumer")

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
			&functionData.DeadLetterTopic,
			"dead-letter-topic",
			"",
			"The topic where messages that are not processed successfully are sent to")
	})
}

func doCreateFunctions(vc *cmdutils.VerbCmd, funcData *pulsar.FunctionData) error {
	err := processArgs(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}

	err = validateFunctionConfigs(funcData.FuncConf)
	if err != nil {
		vc.Command.Help()
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V3)

	if utils.IsPackageURLSupported(funcData.Jar) {
		err = admin.Functions().CreateFuncWithURL(funcData.FuncConf, funcData.Jar)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Created %s successfully\n", funcData.FuncName)
		}
	} else {
		err = admin.Functions().CreateFunc(funcData.FuncConf, funcData.UserCodeFile)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Created %s successfully\n", funcData.FuncName)
		}
	}

	return err
}
