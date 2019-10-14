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
//

package ledger

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func DeleteLedgerCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for deleting a ledger."
	desc.CommandPermission = "none"

	var examples []pulsar.Example
	deleteLedger := pulsar.Example{
		Desc:    "Delete the specified ledger",
		Command: "pulsarctl bookies ledger delete --ledger-id (ledger-id)",
	}
	examples = append(examples, deleteLedger)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Successfully delete the ledger (ledger-id)",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"d",
		desc.ToString(),
		desc.ExampleToString())

	var id int64

	vc.SetRunFuncWithNameArg(func() error {
		return doDeleteLedgerCmd(vc, id)
	}, "the ledger id is not specified or the ledger id is specified more than one")

	vc.FlagSetGroup.InFlagSet("Ledger", func(set *pflag.FlagSet) {
		set.Int64Var(&id, "ledger-id", -1, "the delete ledger id")
		cobra.MarkFlagRequired(set, "ledger-id")
	})
}

func doDeleteLedgerCmd(vc *cmdutils.VerbCmd, id int64) error {

	if id <= 0 {
		return errors.Errorf("invalid ledger id %d", id)
	}

	admin := cmdutils.NewBookieClient()
	err := admin.Ledger().DeleteLedger(id)
	if err == nil {
		vc.Command.Println("Successfully delete the ledger %d", id)
	}

	return err
}
