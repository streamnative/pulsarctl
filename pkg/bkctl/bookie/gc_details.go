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
	"encoding/json"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper/bkdata"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func gcDetailsCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the garbage collection details of a bookie."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the garbage collection details of a bookie.",
		Command: "pulsarctl bookkeeper bookie gc-details",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	details := []bkdata.GCStatus{
		{
			ForceCompacting:         false,
			MajorCompacting:         false,
			MinorCompacting:         false,
			LastMajorCompactionTime: 1,
			LastMinorCompactionTime: 1,
			MajorCompactionCounter:  1,
			MinorCompactionCounter:  1,
		},
	}
	d, _ := json.MarshalIndent(details, "", "")

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Successfully get the garbage collection details of a bookie.",
		Out:  string(d),
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"gc-details",
		"Get the garbage collection details of a bookie.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doGetGCDetails(vc)
	})
}

func doGetGCDetails(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	details, err := admin.Bookie().GCDetails()
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), details)
	}

	return err
}
