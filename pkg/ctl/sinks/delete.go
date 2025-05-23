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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func deleteSinksCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for deleting a Pulsar IO sink connector."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []cmdutils.Example

	del := cmdutils.Example{
		Desc: "Delete a Pulsar IO sink connector",
		Command: "pulsarctl sink delete \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Sink)",
	}
	examples = append(examples, del)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Deleted (the name of a Pulsar Sink) successfully",
	}

	nameNotExistOut := cmdutils.Output{
		Desc: "sink doesn't exist",
		Out:  "code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist",
	}

	out = append(out, successOut, nameNotExistOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete a Pulsar IO Sink connector",
		desc.ToString(),
		desc.ExampleToString(),
		"delete",
	)

	sinkData := &utils.SinkData{}

	// set the run sink
	vc.SetRunFunc(func() error {
		return doDeleteSink(vc, sinkData)
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

func doDeleteSink(vc *cmdutils.VerbCmd, sinkData *utils.SinkData) error {
	err := processBaseArguments(sinkData)
	if err != nil {
		_ = vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	err = admin.Sinks().DeleteSink(sinkData.Tenant, sinkData.Namespace, sinkData.Name)
	if err != nil {
		return err
	}

	vc.Command.Printf("Deleted %s successfully\n", sinkData.Name)
	return nil
}
