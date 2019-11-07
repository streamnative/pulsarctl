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
	"strconv"

	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func statusSinksCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the current status of a Pulsar Sink."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example
	status := cmdutils.Example{
		Desc: "Get the current status of a Pulsar Sink",
		Command: "pulsarctl sink status \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Sink)",
	}
	examples = append(examples, status)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			" \"numInstances\" : 1,\n" +
			" \"numRunning\" : 1,\n" +
			" \"instances\" : [ {\n" +
			"   \"instanceId\" : 0,\n" +
			"   \"status\" : {\n" +
			"     \"running\" : true,\n" +
			"     \"error\" : \"\",\n" +
			"     \"numRestarts\" : 0,\n" +
			"     \"numReadFromPulsar\" : 0,\n" +
			"     \"numSystemExceptions\" : 0,\n" +
			"     \"latestSystemExceptions\" : [ ],\n" +
			"     \"numSinkExceptions\" : 0,\n" +
			"     \"latestSinkExceptions\" : [ ],\n" +
			"     \"numWrittenToSink\" : 0,\n" +
			"     \"lastReceivedTime\" : 0,\n" +
			"     \"workerId\" : \"c-standalone-fw-tengdeMBP.lan-8080\"\n" +
			"   }\n" +
			" } ]\n" +
			"}",
	}

	failOut := cmdutils.Output{
		Desc: "Update contains no change",
		Out:  "[✖]  code: 400 reason: Update contains no change",
	}

	failOutWithNameNotExist := cmdutils.Output{
		Desc: "The name of Pulsar Sink doesn't exist, please check the --name args",
		Out:  "[✖]  code: 404 reason: Sink (your sink name) doesn't exist",
	}

	out = append(out, successOut, failOut, failOutWithNameNotExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"status",
		"Get the current status of a Pulsar Sink",
		desc.ToString(),
		desc.ExampleToString(),
		"getstatus",
	)

	sinkData := &utils.SinkData{}
	// set the run sink
	vc.SetRunFunc(func() error {
		return doStatusSink(vc, sinkData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("SinkConfig", func(flagSet *pflag.FlagSet) {
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

		flagSet.StringVar(
			&sinkData.InstanceID,
			"instance-id",
			"",
			"The sink instanceId (stop all instances if instance-id is not provided)")
	})
}

func doStatusSink(vc *cmdutils.VerbCmd, sinkData *utils.SinkData) error {
	err := processBaseArguments(sinkData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	if sinkData.InstanceID != "" {
		instanceID, err := strconv.Atoi(sinkData.InstanceID)
		if err != nil {
			return err
		}
		sinkInstanceStatusData, err := admin.Sinks().GetSinkStatusWithID(
			sinkData.Tenant, sinkData.Namespace, sinkData.Name, instanceID)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		}
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), sinkInstanceStatusData)
	} else {
		sinkStatus, err := admin.Sinks().GetSinkStatus(sinkData.Tenant, sinkData.Namespace, sinkData.Name)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		}
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), sinkStatus)
	}

	return err
}
