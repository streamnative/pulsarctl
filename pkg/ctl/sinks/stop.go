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

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func stopSinksCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "This command is used for stopping sink instance."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []pulsar.Example

	stop := pulsar.Example{
		Desc: "Stops function instance",
		Command: "pulsarctl sink stop \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Sink)",
	}
	examples = append(examples, stop)

	stopWithInstanceID := pulsar.Example{
		Desc: "Stops function instance with instance ID",
		Command: "pulsarctl sink stop \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name (the name of Pulsar Sink)\n" +
			"\t--instance-id 1",
	}
	examples = append(examples, stopWithInstanceID)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Stopped (the name of a Pulsar Sink) successfully",
	}

	nameNotExistOut := pulsar.Output{
		Desc: "sink doesn't exist",
		Out:  "code: 404 reason: Sink (the name of a Pulsar Sink) doesn't exist",
	}

	out = append(out, successOut, nameNotExistOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"stop",
		"Stops sink instance",
		desc.ToString(),
		desc.ExampleToString(),
		"stop",
	)

	sinkData := &pulsar.SinkData{}
	// set the run sink
	vc.SetRunFunc(func() error {
		return doStopSinks(vc, sinkData)
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

func doStopSinks(vc *cmdutils.VerbCmd, sinkData *pulsar.SinkData) error {
	err := processBaseArguments(sinkData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V3)
	if sinkData.InstanceID != "" {
		instanceID, err := strconv.Atoi(sinkData.InstanceID)
		if err != nil {
			return err
		}
		err = admin.Sinks().StopSinkWithID(sinkData.Tenant, sinkData.Namespace, sinkData.Name, instanceID)
		if err != nil {
			return err
		}
		vc.Command.Printf("Stopped instanceID[%s] of Pulsar Sinks[%s] successfully\n", sinkData.InstanceID, sinkData.Name)
	} else {
		err = admin.Sinks().StopSink(sinkData.Tenant, sinkData.Namespace, sinkData.Name)
		if err != nil {
			return err
		}

		vc.Command.Printf("Stopped %s successfully\n", sinkData.Name)
	}

	return nil
}
