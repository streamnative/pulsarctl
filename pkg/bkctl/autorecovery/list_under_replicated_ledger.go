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

package autorecovery

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/spf13/pflag"
)

func listUnderReplicatedLedgerCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting all the under-replicated ledgers which have been marked " +
		"for re-replication."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "Get all the under-replicated ledgers which have been marked for re-replication.",
		Command: "pulsarctl bookkeeper auto-recovery list-under-replicated-ledger",
	}

	li := cmdutils.Example{
		Desc:    "Get all the under-replicated ledgers of a bookie which have been marked for re-replication.",
		Command: "pulsarctl bookkeeper auto-recovery list-under-replicated-ledger --include (bookie-ip:bookie-port)",
	}

	le := cmdutils.Example{
		Desc:    "Get all the under-replicated ledgers except a bookie which have been marked for re-replication.",
		Command: "pulsarctl bookkeeper auto-recovery list-under-replicated-ledger --exclude (bookie-ip:bookie-port)",
	}
	examples = append(examples, list, li, le)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Get the under-replicated ledgers successfully.",
		Out: `{
    [ledgerId1, ledgerId2...]
}`,
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list-under-replicated-ledger",
		"Get all the under-replicated ledgers which have been marked for re-replication.",
		desc.ToString(),
		desc.ExampleToString())

	var include string
	var exclude string
	var show bool

	vc.SetRunFunc(func() error {
		return doListUnderReplicatedLedger(vc, include, exclude, show)
	})

	vc.FlagSetGroup.InFlagSet("List under replicated ledgers", func(set *pflag.FlagSet) {
		set.StringVar(&include, "include", "", "Show the under-replicated ledger of the bookie.")
		set.StringVar(&exclude, "exclude", "", "Show the under-replicated ledger exclude the bookie.")
		set.BoolVar(&show, "show", false, "Show the replicate ledger list.")
	})
	vc.EnableOutputFlagSet()
}

func doListUnderReplicatedLedger(vc *cmdutils.VerbCmd, include, exclude string, show bool) error {
	admin := cmdutils.NewBookieClient()
	var l interface{}
	var err error
	if show {
		l, err = admin.AutoRecovery().PrintListUnderReplicatedLedger(include, exclude)
	} else {
		l, err = admin.AutoRecovery().ListUnderReplicatedLedger(include, exclude)
	}

	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), l)
	}

	return err
}
