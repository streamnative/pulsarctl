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

package ledger

import (
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/pkg/errors"
)

func deleteCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for deleting a ledger."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	deleteLedger := cmdutils.Example{
		Desc:    "Delete the specified ledger",
		Command: "pulsarctl bookkeeper ledger delete (ledger-id)",
	}
	examples = append(examples, deleteLedger)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully delete the ledger (ledger-id)",
	}
	out = append(out, successOut, argError)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete a ledger",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteCmd(vc)
	}, "the ledger id is not specified or the ledger id is specified more than one")
}

func doDeleteCmd(vc *cmdutils.VerbCmd) error {
	id, err := strconv.ParseInt(vc.NameArg, 10, 64)
	if err != nil || id < 0 {
		return errors.Errorf("invalid ledger id %s", vc.NameArg)
	}

	admin := cmdutils.NewBookieClient()
	err = admin.Ledger().Delete(id)
	if err == nil {
		vc.Command.Printf("Successfully delete the ledger %d\n", id)
	}

	return err
}
