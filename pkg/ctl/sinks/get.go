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
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsar-admin-go/pkg/admin/config"
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getSinksCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the information about a Pulsar IO sink connector"
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example

	get := cmdutils.Example{
		Desc: "Get the information about a Pulsar IO sink connector",
		Command: "pulsarctl sink get \n" +
			"\t--tenant public\n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Sink)",
	}

	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			" \"tenant\": \"public\",\n" +
			" \"namespace\": \"default\",\n" +
			" \"name\": \"mysql-jdbc-sink\",\n" +
			" \"className\": \"org.apache.pulsar.io.jdbc.JdbcAutoSchemaSink\",\n" +
			" \"inputSpecs\": {\n" +
			"   \"test-jdbc\": {\n" +
			"     \"isRegexPattern\": false\n" +
			"   }\n" +
			" },\n" +
			" \"configs\": {\n" +
			"   \"password\": \"jdbc\",\n" +
			"   \"jdbcUrl\": \"jdbc:mysql://127.0.0.1:3306/test_jdbc\",\n" +
			"   \"userName\": \"root\",\n" +
			"   \"tableName\": \"test_jdbc\"\n" +
			" },\n" +
			" \"parallelism\": 1,\n" +
			" \"processingGuarantees\": \"ATLEAST_ONCE\",\n" +
			" \"retainOrdering\": false,\n" +
			" \"autoAck\": true\n" +
			"}",
	}

	nameNotExistOut := cmdutils.Output{
		Desc: "sink doesn't exist",
		Out:  "code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist",
	}
	out = append(out, successOut, nameNotExistOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"Get the information about a Pulsar IO sink connector",
		desc.ToString(),
		desc.ExampleToString(),
		"get",
	)

	sinkData := &utils.SinkData{}
	// set the run sink
	vc.SetRunFunc(func() error {
		return doGetSinks(vc, sinkData)
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
	})
	vc.EnableOutputFlagSet()
}

func doGetSinks(vc *cmdutils.VerbCmd, sinkData *utils.SinkData) error {
	err := processBaseArguments(sinkData)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	sinkConfig, err := admin.Sinks().GetSink(sinkData.Tenant, sinkData.Namespace, sinkData.Name)
	if err == nil {
		oc := cmdutils.NewOutputContent().WithObject(sinkConfig)
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
	}

	return err
}
