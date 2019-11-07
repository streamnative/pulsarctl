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

package brokerstats

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func dumpLoadReport(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Dump the broker load-report"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Dump the broker load-report",
		Command: "pulsarctl broker-stats load-report",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Print the broker load-report info",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"load-report",
		"Dump the broker load-report",
		desc.ToString(),
		desc.ExampleToString(),
		"load-report")

	vc.SetRunFunc(func() error {
		return doDumpLoadReport(vc)
	})
}

func doDumpLoadReport(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	loadReport, err := admin.BrokerStats().GetLoadReport()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), loadReport)
	}
	return err
}
