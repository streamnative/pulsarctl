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

package functionsworker

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func functionsStats(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Dump all functions stats running on this broker"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	stats := cmdutils.Example{
		Desc:    "Dump all functions stats running on this broker",
		Command: "pulsarctl functions-worker function-stats",
	}
	examples = append(examples, stats)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "[ ]",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"function-stats",
		"Dump all functions stats running on this broker",
		desc.ToString(),
		desc.ExampleToString(),
		"function-stats",
	)

	// set the run function
	vc.SetRunFunc(func() error {
		return doFunctionStats(vc)
	})
}

func doFunctionStats(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	fnStats, err := admin.FunctionsWorker().GetFunctionsStats()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), fnStats)
	}

	return err
}
