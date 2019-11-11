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

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func startSourcesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for starting a stopped source instance."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example

	start := cmdutils.Example{
		Desc: "Start source instance",
		Command: "pulsarctl source start \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Source)",
	}
	examples = append(examples, start)

	startWithInstanceID := cmdutils.Example{
		Desc: "Starts a stopped source instance with instance ID",
		Command: "pulsarctl source start \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Source)\n" +
			"\t--instance-id 1",
	}
	examples = append(examples, startWithInstanceID)
	desc.CommandExamples = examples
	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Started (the name of a Pulsar Source) successfully",
	}

	nameNotExistOut := cmdutils.Output{
		Desc: "source doesn't exist",
		Out:  "code: 404 reason: Source (the name of a Pulsar Source) doesn't exist",
	}
	out = append(out, successOut, nameNotExistOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"start",
		"Start source instance",
		desc.ToString(),
		desc.ExampleToString(),
		"start",
	)

	sourceData := &utils.SourceData{}

	// set the run source
	vc.SetRunFunc(func() error {
		return doStartSource(vc, sourceData)
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

func doStartSource(vc *cmdutils.VerbCmd, sourceData *utils.SourceData) error {
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
		err = admin.Sources().StartSourceWithID(sourceData.Tenant, sourceData.Namespace, sourceData.Name, instanceID)
		if err != nil {
			return err
		}
		vc.Command.Printf("Started instanceID[%s] of Pulsar Source[%s] successfully\n",
			sourceData.InstanceID, sourceData.Name)
	} else {
		err = admin.Sources().StartSource(sourceData.Tenant, sourceData.Namespace, sourceData.Name)
		if err != nil {
			return err
		}

		vc.Command.Printf("Started %s successfully\n", sourceData.Name)
	}

	return nil
}
