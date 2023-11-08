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
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func setReadonlyStateCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the readonly state of a bookie."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Set the readonly state of the bookie.",
		Command: "pulsarctl bookkeeper bookie set-readonly <true/false>",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Successfully set the readonly state of a bookie.",
		Out:  "Successfully set the readonly state of a bookie",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-readonly",
		"Set the readonly state of a bookie.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doSetReadonlyState(vc)
	}, "the readonly state is boolean")
}

func doSetReadonlyState(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	readonly, err := strconv.ParseBool(vc.NameArg)
	if err != nil {
		return err
	}
	err = admin.Bookie().SetReadonlyState(readonly)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), "Successfully set the readonly state of a bookie")
	}

	return err
}
