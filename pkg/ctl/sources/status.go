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

package sources

import (
	"strconv"

	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func statusSourcesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Check the current status of a Pulsar Source."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example
	status := cmdutils.Example{
		Desc: "Check the current status of a Pulsar Source",
		Command: "pulsarctl source status \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Source)",
	}
	examples = append(examples, status)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"numInstances\" : 1,\n" +
			"  \"numRunning\" : 1,\n" +
			"  \"instances\" : [ {\n" +
			"    \"instanceId\" : 0,\n" +
			"    \"status\" : {\n" +
			"      \"running\" : true,\n" +
			"      \"error\" : \"\",\n" +
			"      \"numRestarts\" : 0,\n" +
			"      \"numReceivedFromSource\" : 0,\n" +
			"      \"numSystemExceptions\" : 0,\n" +
			"      \"latestSystemExceptions\" : [ ],\n" +
			"      \"numSourceExceptions\" : 0,\n" +
			"      \"latestSourceExceptions\" : [ ],\n" +
			"      \"numWritten\" : 0,\n" +
			"      \"lastReceivedTime\" : 0,\n" +
			"      \"workerId\" : \"c-standalone-fw-7e0cf1b3bf9d-8080\"\n" +
			"    }\n" +
			"  } ]\n" +
			"}",
	}

	failOut := cmdutils.Output{
		Desc: "Update contains no change",
		Out:  "[✖]  code: 400 reason: Update contains no change",
	}

	failOutWithNameNotExist := cmdutils.Output{
		Desc: "The name of Pulsar Source doesn't exist, please check the --name args",
		Out:  "[✖]  code: 404 reason: Source (your source name) doesn't exist",
	}

	out = append(out, successOut, failOut, failOutWithNameNotExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"status",
		"Check the current status of a Pulsar Source",
		desc.ToString(),
		desc.ExampleToString(),
		"getstatus",
	)

	sourceData := &utils.SourceData{}
	// set the run source
	vc.SetRunFunc(func() error {
		return doStatusSource(vc, sourceData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("SourceConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&sourceData.Tenant,
			"tenant",
			"",
			"The source's tenant")

		flagSet.StringVar(
			&sourceData.Namespace,
			"namespace",
			"",
			"The source's namespace")

		flagSet.StringVar(
			&sourceData.Name,
			"name",
			"",
			"The source's name")

		flagSet.StringVar(
			&sourceData.InstanceID,
			"instance-id",
			"",
			"The source instanceId (stop all instances if instance-id is not provided)")
	})
}

func doStatusSource(vc *cmdutils.VerbCmd, sourceData *utils.SourceData) error {
	err := processBaseArguments(sourceData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	if sourceData.InstanceID != "" {
		instanceID, err := strconv.Atoi(sourceData.InstanceID)
		if err != nil {
			return err
		}
		sourceInstanceStatusData, err := admin.Sources().GetSourceStatusWithID(
			sourceData.Tenant, sourceData.Namespace, sourceData.Name, instanceID)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		}
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), sourceInstanceStatusData)
	} else {
		sourceStatus, err := admin.Sources().GetSourceStatus(sourceData.Tenant, sourceData.Namespace, sourceData.Name)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		}
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), sourceStatus)
	}

	return err
}
