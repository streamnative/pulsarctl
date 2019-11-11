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

package brokers

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func deleteDynamicConfigCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Delete dynamic-serviceConfiguration of broker"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "Delete dynamic-serviceConfiguration of broker",
		Command: "pulsarctl brokers delete-dynamic-config --config (config name)",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Deleted dynamic config: (config name) successful.",
	}

	failOut := cmdutils.Output{
		Desc: "Can't update non-dynamic configuration, please check `--config` arg.",
		Out:  "[✖]  code: 412 reason:  Can't update non-dynamic configuration",
	}

	out = append(out, successOut, failOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete-dynamic-config",
		"Delete dynamic-serviceConfiguration of broker",
		desc.ToString(),
		desc.ExampleToString(),
		"delete-dynamic-config")

	brokerData := &utils.BrokerData{}

	vc.SetRunFunc(func() error {
		return doDeleteDynamicConf(vc, brokerData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("BrokerData", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&brokerData.ConfigName,
			"config",
			"",
			"service-configuration name")
		cobra.MarkFlagRequired(flagSet, "config")
	})
}

func doDeleteDynamicConf(vc *cmdutils.VerbCmd, brokerData *utils.BrokerData) error {
	admin := cmdutils.NewPulsarClient()
	err := admin.Brokers().DeleteDynamicConfiguration(brokerData.ConfigName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Deleted dynamic config: %s successful\n", brokerData.ConfigName)
	}
	return err
}
