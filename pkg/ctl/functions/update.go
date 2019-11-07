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
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	util "github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func updateFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Update a Pulsar Function that has been deployed to a Pulsar cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	update := cmdutils.Example{
		Desc: "Change the output topic of a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--tenant public \n" +
			"\t--namespace default \n" +
			"\t--name update-function \n" +
			"\t--output test-output-topic",
	}
	examples = append(examples, update)

	updateWithConf := cmdutils.Example{
		Desc: "Update a Pulsar Function using a function config yaml file",
		Command: "pulsarctl functions update \n" +
			"\t--function-config-file (the path of function config yaml file) \n" +
			"\t--jar (the path of user code jar)",
	}
	examples = append(examples, updateWithConf)

	updateWithLogTopic := cmdutils.Example{
		Desc: "Change the log topic of a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--log-topic persistent://public/default/test-log-topic\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, updateWithLogTopic)

	updateWithDeadLetterTopic := cmdutils.Example{
		Desc: "Change the dead letter topic of a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--dead-letter-topic persistent://public/default/test-dead-letter-topic\n" +
			"\t--max-message-retries 10\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, updateWithDeadLetterTopic)

	updateWithUserConfig := cmdutils.Example{
		Desc: "Update the user configs of a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--user-config \"{\"publishTopic\":\"publishTopic\", \"key\":\"pulsar\"}\"\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, updateWithUserConfig)

	updateWithCustomSchemasInputs := cmdutils.Example{
		Desc: "Change the schemas of the input topics for a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--custom-schema-inputs \"{\"topic-1\":\"schema.STRING\", \"topic-2\":\"schema.JSON\"}\"\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, updateWithCustomSchemasInputs)

	updateWithSchema := cmdutils.Example{
		Desc: "Change the schema type of the input topic for a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--schema-type schema.STRING\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, updateWithSchema)

	updateWithParallelism := cmdutils.Example{
		Desc: "Change the parallelism of a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--parallelism 1\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, updateWithParallelism)

	updateWithResource := cmdutils.Example{
		Desc: "Change the resource usage for a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--ram 5656565656\n" +
			"\t--disk 8080808080808080\n" +
			"\t--cpu 5.0\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, updateWithResource)

	updateWithWindowFunctions := cmdutils.Example{
		Desc: "Update the window configurations for a Pulsar Function",
		Command: "pulsarctl functions update \n" +
			"\t--window-length-count 10\n" +
			"\t--window-length-duration-ms 1000\n" +
			"\t--sliding-interval-count 3\n" +
			"\t--sliding-interval-duration-ms 1000\n" +
			"\t// Other function parameters ",
	}
	examples = append(examples, updateWithWindowFunctions)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Updated (the name of a Pulsar Function) successfully",
	}

	failOut := cmdutils.Output{
		Desc: "Update contains no change",
		Out:  "[✖]  code: 400 reason: Update contains no change",
	}

	failOutWithNameNotExist := cmdutils.Output{
		Desc: "The name of Pulsar Functions doesn't exist, please check the --name args",
		Out:  "[✖]  code: 404 reason: Function (your function name) doesn't exist",
	}

	out = append(out, successOut, failOut, failOutWithNameNotExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"update",
		"Update a Pulsar Function that has been deployed to a Pulsar cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"update",
	)

	functionData := &util.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doUpdateFunctions(vc, functionData)
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
			"Path to the JAR file for the function (if the function is written in Java). "+
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

		flagSet.StringVar(
			&functionData.TopicsPattern,
			"topics-pattern",
			"",
			"The topic pattern to consume from list of topics under a namespace that match the pattern. "+
				"[--input] and [--topic-pattern] are mutually exclusive. Add SerDe class name for a pattern "+
				"in --custom-serde-inputs (supported for java fun only)")

		flagSet.StringVar(
			&functionData.Inputs,
			"inputs",
			"",
			"The input topic or topics (multiple topics can be specified as a comma-separated list) of a Pulsar Function")

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
			&functionData.UserConfig,
			"user-config",
			"",
			"User-defined config key/values")

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
			"The ram in bytes that need to be allocated per function instance(applicable only to process/docker runtime)")

		flagSet.Int64Var(
			&functionData.Disk,
			"disk",
			0,
			"The disk in bytes that need to be allocated per function instance(applicable only to docker runtime)")

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

		flagSet.BoolVar(
			&functionData.UpdateAuthData,
			"update-auth-data",
			false,
			"Whether or not to update the auth data")
	})
}

func doUpdateFunctions(vc *cmdutils.VerbCmd, funcData *util.FunctionData) error {
	err := processArgs(funcData)
	if err != nil {
		vc.Command.Help()
		return err
	}

	err = checkArgsForUpdate(funcData.FuncConf)
	if err != nil {
		vc.Command.Help()
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)

	updateOptions := util.NewUpdateOptions()
	updateOptions.UpdateAuthData = funcData.UpdateAuthData

	if utils.IsPackageURLSupported(funcData.Jar) {
		err = admin.Functions().UpdateFunctionWithURL(funcData.FuncConf, funcData.Jar, updateOptions)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Updated instanceID[%s] of Pulsar Functions[%s] successfully\n",
				funcData.InstanceID, funcData.FuncName)
		}
	} else {
		err = admin.Functions().UpdateFunction(funcData.FuncConf, funcData.UserCodeFile, updateOptions)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Updated %s successfully\n", funcData.FuncName)
		}
	}

	return err
}
