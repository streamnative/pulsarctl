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
    `fmt`
    `github.com/spf13/pflag`
    `github.com/streamnative/pulsarctl/pkg/cmdutils`
    `github.com/streamnative/pulsarctl/pkg/pulsar`
)

func createFunctionsCmd(vc *cmdutils.VerbCmd)  {
    vc.SetDescription(
        "create",
        "",
        "Create a Pulsar Function in cluster mode (deploy it on a Pulsar cluster)",
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
            "Path to the JAR file for the function (if the function is written in Java). " +
                "It also supports URL path [http/https/file (file protocol assumes that file " +
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
            "The topic pattern to consume from list of topics under a namespace that match the pattern. " +
                "[--input] and [--topic-pattern] are mutually exclusive. Add SerDe class name for a pattern in " +
                "--custom-serde-inputs (supported for java fun only)")

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

        flagSet.IntVar(
            &functionData.ProcessingGuarantees,
            "processing-guarantees",
            0,
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

        flagSet.BoolVar(
            &functionData.AutoAck,
            "auto-ack",
            false,
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
        return err
    }

    admin := cmdutils.NewPulsarClient()

    if isFunctionPackageUrlSupported(funcData.Jar) {
        err = admin.Functions().CreateFuncWithUrl(funcData.FuncConf, funcData.Jar)
        if err != nil {
            cmdutils.PrintError(vc.Command.OutOrStderr(), err)
            return err
        }
    } else {
        err = admin.Functions().CreateFunc(funcData.FuncConf, funcData.UserCodeFile)
        if err != nil {
            cmdutils.PrintError(vc.Command.OutOrStderr(), err)
            return err
        }
    }

    fmt.Println("Created successfully")

    return nil
}
