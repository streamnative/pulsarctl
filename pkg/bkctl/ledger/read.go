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
	"github.com/spf13/pflag"
)

func readCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for reading a range of entries of a ledger."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	r := cmdutils.Example{
		Desc:    "Read a range of entries of the specified ledger",
		Command: "pulsar bookkeeper ledger read (ledger-id)",
	}

	rs := cmdutils.Example{
		Desc:    "Read the entries of the specified ledger started from the given entry id",
		Command: "pulsar bookkeeper ledger --start (entry-id) (ledger-id)",
	}

	rse := cmdutils.Example{
		Desc:    "Read the specified range of entries of the specified ledger",
		Command: "pulsar bookkeeper ledger --start (entry-id) --end (entry-id) (ledger-id)",
	}

	examples = append(examples, r, rs, rse)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: `{
	"ledger-id", "message"
}`,
	}
	out = append(out, successOut, argError)
	desc.CommandOutput = out

	vc.SetDescription(
		"read",
		"Read a range of entries of a ledger",
		desc.ToString(),
		desc.ExampleToString())

	var start int64
	var end int64

	vc.SetRunFuncWithNameArg(func() error {
		return doRead(vc, start, end)
	}, "the ledger id is not specified or the ledger id is specified more than one")

	vc.FlagSetGroup.InFlagSet("Read Ledger", func(set *pflag.FlagSet) {
		set.Int64VarP(&start, "start", "b", -1,
			"")
		set.Int64VarP(&end, "end", "e", -1, "")
	})
	vc.EnableOutputFlagSet()
}

func doRead(vc *cmdutils.VerbCmd, start, end int64) error {
	id, err := strconv.ParseInt(vc.NameArg, 10, 64)
	if err != nil || id < 0 {
		return errors.Errorf("invalid ledger id %s", vc.NameArg)
	}

	if start != -1 && start < 0 {
		return errors.Errorf("invalid start ledger id %d", start)
	}

	if end != -1 && end < 0 {
		return errors.Errorf("invalid end ledger id %d", end)
	}

	admin := cmdutils.NewBookieClient()
	info, err := admin.Ledger().Read(id, start, end)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), info)
	}

	return err
}
