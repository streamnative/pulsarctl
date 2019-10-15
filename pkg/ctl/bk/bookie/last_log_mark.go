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
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func LastLogMarkCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting the last log marker."
	desc.CommandPermission = "none"

	var exmaples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Get the last log marker",
		Command: "pulsarctl bk bookies last-log-marker",
	}
	exmaples = append(exmaples, get)
	desc.CommandExamples = exmaples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: `{
    JournalId1 : position1,
    JournalId2 : position2,
    ...
}`,
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"last-log-marker",
		"Get the last log marker",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doGetLastLogMark(vc)
	})
}

func doGetLastLogMark(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	marker, err := admin.Bookie().LastLogMark()
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), marker)
	}

	return err
}
