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
	"sort"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/spf13/pflag"
)

func listCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for listing all the ledgers."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "List all the ledgers",
		Command: "pulsarctl bookkeeper ledger list",
	}

	showMeta := cmdutils.Example{
		Desc:    "List all the ledgers and the metadata of the ledger",
		Command: "pulsarctl bookkeeper ledger list --show-metadata",
	}
	examples = append(examples, list, showMeta)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "[1,2,3,4]",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"list all the ledgers",
		desc.ToString(),
		desc.ExampleToString())

	var show bool

	vc.SetRunFunc(func() error {
		return doListCmd(vc, show)
	})

	vc.FlagSetGroup.InFlagSet("List Ledgers", func(set *pflag.FlagSet) {
		set.BoolVarP(&show, "show-metadata", "p", false,
			"Show the metadata of the ledgers")
	})
	vc.EnableOutputFlagSet()
}

func doListCmd(vc *cmdutils.VerbCmd, showMeta bool) error {
	admin := cmdutils.NewBookieClient()
	ledgers, err := admin.Ledger().List(showMeta)
	if err == nil {
		if !showMeta {
			ledgerList := make([]int64, 0)
			for k := range ledgers {
				ledgerList = append(ledgerList, k)
			}
			sort.Slice(ledgerList, func(i, j int) bool {
				return ledgerList[i] < ledgerList[j]
			})
			vc.Command.Println(ledgerList)
		} else {
			cmdutils.PrintJSON(vc.Command.OutOrStdout(), ledgers)
		}
	}
	return err
}
