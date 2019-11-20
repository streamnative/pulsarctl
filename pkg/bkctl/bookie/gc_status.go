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

package bookie

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func gcStatusCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the status of the running GC."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the status of the running GC",
		Command: "pulsarctl bookkeeper bookie gc-status",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: `{
    "is_in_force_gc" : "false"
}`,
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"gc-status",
		"Get the status of the running GC",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doGetGCStatus(vc)
	})
}

func doGetGCStatus(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	status, err := admin.Bookie().GCStatus()
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), status)
	}

	return err
}
