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

func updateDynamicConfig(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Update dynamic-serviceConfiguration of broker"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "Update dynamic-serviceConfiguration of broker",
		Command: "pulsarctl brokers update-dynamic-config --config (config name) --value (config value)",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Update dynamic config: (configName) successful.",
	}

	failOut := cmdutils.Output{
		Desc: "Can't update non-dynamic configuration, please check `--config` arg.",
		Out:  "[✖]  code: 412 reason:  Can't update non-dynamic configuration",
	}

	out = append(out, successOut, failOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"update-dynamic-config",
		"Update dynamic-serviceConfiguration of broker",
		desc.ToString(),
		desc.ExampleToString(),
		"update-dynamic-config")

	brokerData := &utils.BrokerData{}

	vc.SetRunFunc(func() error {
		return doUpdateDynamic(vc, brokerData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("BrokerData", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&brokerData.ConfigName,
			"config",
			"",
			"service-configuration name")

		flagSet.StringVar(
			&brokerData.ConfigValue,
			"value",
			"",
			"service-configuration value")

		cobra.MarkFlagRequired(flagSet, "config")
		cobra.MarkFlagRequired(flagSet, "value")
	})
}

func doUpdateDynamic(vc *cmdutils.VerbCmd, brokerData *utils.BrokerData) error {
	admin := cmdutils.NewPulsarClient()
	err := admin.Brokers().UpdateDynamicConfiguration(brokerData.ConfigName, brokerData.ConfigValue)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Printf("Update dynamic config: %s successful\n", brokerData.ConfigName)
	}
	return err
}
