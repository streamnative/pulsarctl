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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getSourcesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Gets the information about a Pulsar IO source connector"
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example

	get := cmdutils.Example{
		Desc: "Gets the information about a Pulsar IO source connector",
		Command: "pulsarctl source get \n" +
			"\t--tenant public\n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Source)",
	}

	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"tenant\": \"public\",\n" +
			"  \"namespace\": \"default\",\n" +
			"  \"name\": \"kafka\",\n" +
			"  \"className\": \"org.apache.pulsar.io.kafka.KafkaBytesSource\",\n" +
			"  \"topicName\": \"my-topic\",\n" +
			"  \"configs\": {\n" +
			"    \"bootstrapServers\": \"pulsar-kafka:9092\",\n" +
			"    \"groupId\": \"test-pulsar-io1\",\n" +
			"    \"topic\": \"my-topic\",\n" +
			"    \"sessionTimeoutMs\": \"10000\",\n" +
			"    \"autoCommitEnabled\": \"false\"\n" +
			"  },\n" +
			"  \"parallelism\": 1,\n" +
			"  \"processingGuarantees\": \"ATLEAST_ONCE\"\n" +
			"}\n",
	}

	nameNotExistOut := cmdutils.Output{
		Desc: "source doesn't exist",
		Out:  "code: 404 reason: Source (the name of a Pulsar Source) doesn't exist",
	}
	out = append(out, successOut, nameNotExistOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"Gets the information about a Pulsar IO source connector",
		desc.ToString(),
		desc.ExampleToString(),
		"get",
	)

	sourceData := &utils.SourceData{}
	// set the run source
	vc.SetRunFunc(func() error {
		return doGetSources(vc, sourceData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("SourcesConfig", func(flagSet *pflag.FlagSet) {
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
	})
	vc.EnableOutputFlagSet()
}

func doGetSources(vc *cmdutils.VerbCmd, sourceData *utils.SourceData) error {
	err := processBaseArguments(sourceData)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	sourceConfig, err := admin.Sources().GetSource(sourceData.Tenant, sourceData.Namespace, sourceData.Name)
	if err == nil {
		oc := cmdutils.NewOutputContent().WithObject(sourceConfig)
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
	}

	return err
}
