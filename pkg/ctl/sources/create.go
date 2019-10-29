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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func createSourcesCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Submit a Pulsar IO source connector to run in a Pulsar cluster."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []pulsar.Example
	create := pulsar.Example{
		Desc: "Create a Pulsar Source in cluster mode",
		Command: "pulsarctl sources create \n" +
			"\t--tenant public \n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Sources) \n" +
			"\t--destination-topic-name kafka-topic  \n" +
			"\t--classname org.apache.pulsar.io.kafka.KafkaBytesSource \n" +
			"\t--archive pulsar-io-kafka-2.4.0.nar \n" +
			"\t--source-config-file conf/kafkaSourceConfig.yaml \n" +
			"\t--parallelism 1",
	}

	createWithPkgURL := pulsar.Example{
		Desc: "Create a Pulsar Source in cluster mode with pkg URL",
		Command: "pulsarctl source create \n" +
			"\t--tenant public \n" +
			"\t--namespace default \n" +
			"\t--name (the name of Pulsar Source) \n" +
			"\t--destination-topic-name kafka-topic  \n" +
			"\t--classname org.apache.pulsar.io.kafka.KafkaBytesSource \n" +
			"\t--archive file://(or http://) + /examples/api-examples.nar",
	}

	createWithSchema := pulsar.Example{
		Desc: "Create a Pulsar Source in cluster mode with schema type",
		Command: "pulsarctl source create \n" +
			"\t--schema-type schema.STRING\n" +
			"\t// Other source parameters ",
	}

	createWithParallelism := pulsar.Example{
		Desc: "Create a Pulsar Source in cluster mode with parallelism",
		Command: "pulsarctl source create \n" +
			"\t--parallelism 1\n" +
			"\t// Other source parameters ",
	}

	createWithResource := pulsar.Example{
		Desc: "Create a Pulsar Source in cluster mode with resource",
		Command: "pulsarctl source create \n" +
			"\t--ram 5656565656\n" +
			"\t--disk 8080808080808080\n" +
			"\t--cpu 5.0\n" +
			"\t// Other source parameters ",
	}

	createWithSourceConfig := pulsar.Example{
		Desc: "Create a Pulsar Source in cluster mode with source config",
		Command: "pulsarctl source create \n" +
			"\t--source-config \"{\"publishTopic\":\"publishTopic\", \"key\":\"pulsar\"}\"\n" +
			"\t// Other source parameters ",
	}

	createWithProcessingGuarantees := pulsar.Example{
		Desc: "Create a Pulsar Source in cluster mode with processing guarantees",
		Command: "pulsarctl source create \n" +
			"\t--processing-guarantees EFFECTIVELY_ONCE\n" +
			"\t// Other source parameters ",
	}

	examples = append(examples, create, createWithPkgURL, createWithSchema, createWithParallelism,
		createWithResource, createWithSourceConfig, createWithProcessingGuarantees)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Created (the name of a Pulsar Sources) successfully",
	}

	failureOut := pulsar.Output{
		Desc: "source archive not specified, please check --archive arg",
		Out:  "[✖]  Source archive not specified",
	}

	//Cannot specify both archive and source-type
	sourceTypeOut := pulsar.Output{
		Desc: "Cannot specify both archive and source-type, please check --archive and --source-type args",
		Out:  "[✖]  Cannot specify both archive and source-type",
	}

	out = append(out, successOut, failureOut, sourceTypeOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Submit a Pulsar IO source connector to run in a Pulsar cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"create",
	)

	sourceData := &pulsar.SourceData{}

	// set the run source
	vc.SetRunFunc(func() error {
		return doCreateSources(vc, sourceData)
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

		flagSet.StringVarP(
			&sourceData.SourceType,
			"source-type",
			"t",
			"",
			"The source's connector provider")

		flagSet.StringVar(
			&sourceData.ProcessingGuarantees,
			"processing-guarantees",
			"",
			"The processing guarantees (aka delivery semantics) applied to the source")

		flagSet.StringVar(
			&sourceData.DestinationTopicName,
			"destination-topic-name",
			"",
			"The Pulsar topic to which data is sent")

		flagSet.StringVar(
			&sourceData.DeserializationClassName,
			"deserialization-classname",
			"",
			"The SerDe classname for the source")

		flagSet.StringVar(
			&sourceData.SchemaType,
			"schema-type",
			"",
			"The schema type (either a builtin schema like 'avro', 'json', etc.. or custom Schema class name "+
				"to be used to encode messages emitted from the source")

		flagSet.IntVar(
			&sourceData.Parallelism,
			"parallelism",
			0,
			"The source's parallelism factor (i.e. the number of source instances to run)")

		flagSet.StringVarP(
			&sourceData.Archive,
			"archive",
			"a",
			"",
			"The path to the NAR archive for the Source. It also supports url-path [http/https/file "+
				"(file protocol assumes that file already exists on worker host)] from which worker can download the package")

		flagSet.StringVar(
			&sourceData.ClassName,
			"classname",
			"",
			"The source's class name if archive is file-url-path (file://)")

		flagSet.StringVar(
			&sourceData.SourceConfigFile,
			"source-config-file",
			"",
			"he path to a YAML config file specifying the ")

		flagSet.Float64Var(
			&sourceData.CPU,
			"cpu",
			0.0,
			"The CPU (in cores) that needs to be allocated per source instance (applicable only to Docker runtime)")

		flagSet.Int64Var(
			&sourceData.RAM,
			"ram",
			0,
			"The RAM (in bytes) that need to be allocated per source instance (applicable only to the "+
				"process and Docker runtimes)")

		flagSet.Int64Var(
			&sourceData.Disk,
			"disk",
			0,
			"The disk (in bytes) that need to be allocated per source instance (applicable only to Docker runtime)")

		flagSet.StringVar(
			&sourceData.SourceConfigFile,
			"source-config",
			"",
			"Source config key/values")
	})
}

func doCreateSources(vc *cmdutils.VerbCmd, sourceData *pulsar.SourceData) error {
	err := processArguments(sourceData)
	if err != nil {
		vc.Command.Help()
		return err
	}

	err = validateSourceConfigs(sourceData.SourceConf)
	if err != nil {
		vc.Command.Help()
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V3)
	if utils.IsPackageURLSupported(sourceData.Archive) {
		err = admin.Sources().CreateSourceWithURL(sourceData.SourceConf, sourceData.Archive)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Created instanceID[%s] of Pulsar Sources[%s] successfully\n",
				sourceData.InstanceID, sourceData.Name)
		}
	} else {
		err = admin.Sources().CreateSource(sourceData.SourceConf, sourceData.Archive)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Printf("Created %s successfully\n", sourceData.Name)
		}
	}

	return err
}
