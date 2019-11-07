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

func dumpAllocatorStats(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Dump the allocator stats"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Dump the allocator stats",
		Command: "pulsarctl broker-stats allocator-stats",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Print allocator stats info",
	}

	failOut := cmdutils.Output{
		Desc: "the namespace name is not specified or the namespace name is specified more than one",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}
	out = append(out, successOut, failOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"allocator-stats",
		"Dump the allocator stats",
		desc.ToString(),
		desc.ExampleToString(),
		"allocator-stats")

	vc.SetRunFuncWithNameArg(func() error {
		return doDumpAllocatorStats(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doDumpAllocatorStats(vc *cmdutils.VerbCmd) error {
	allocatorName := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	stats, err := admin.BrokerStats().GetAllocatorStats(allocatorName)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), stats)
	}
	return err
}
