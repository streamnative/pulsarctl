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

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func StateCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting the state of the bookie."
	desc.CommandPermission = "none"

	var exmaples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Get the state of the bookie",
		Command: "pulsarctl bk bookies state",
	}
	exmaples = append(exmaples, get)
	desc.CommandExamples = exmaples

	s := pulsar.State{
		Running:                        true,
		ReadOnly:                       true,
		ShuttingDown:                   false,
		AvailableForHighPriorityWrites: false,
	}
	o, _ := json.MarshalIndent(s, "", "   ")

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  string(o),
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"state",
		"Get the state of the bookie",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doGetState(vc)
	})
}

func doGetState(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	state, err := admin.Bookie().State()
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), state)
	}

	return err
}
