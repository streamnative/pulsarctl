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

func expandStorageCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for expanding storage for a bookie."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	expand := cmdutils.Example{
		Desc:    "Expanding storage for a bookie.",
		Command: "pulsarctl bookkeeper bookie expand-storage",
	}
	examples = append(examples, expand)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Expand the storage successfully.",
		Out:  "Successfully expand the storage",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"expand-storage",
		"Expand storage for a bookie.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doExpandStorage(vc)
	})
}

func doExpandStorage(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	err := admin.Bookie().ExpandStorage()
	if err == nil {
		vc.Command.Println("Successfully expand the storage.")
	}

	return err
}
