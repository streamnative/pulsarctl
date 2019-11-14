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
	"encoding/json"
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper/bkdata"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/pkg/errors"
)

func getCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the metadata of a ledger."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the metadata of the specified ledger",
		Command: "pulsarctl bookkeeper ledger get (ledger-i)",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	metadata := bkdata.LedgerMetadata{
		MetadataFormatVersion: 1,
		Ensemble:              1,
		WriteQuorum:           1,
		AckQuorum:             1,
		Length:                1,
		LastEntryID:           1,
		Ctime:                 1,
		CToken:                0,
		State:                 "CLOSED",
		DigestType:            "MAC",
		Ensembles: map[int64][]bkdata.BookieSocketAddress{
			1: {
				bkdata.BookieSocketAddress{
					HostName: "www.examples.com",
					Port:     8080,
				},
			},
		},
		CurrentEnsemble: []bkdata.BookieSocketAddress{
			{
				HostName: "www.example.com",
				Port:     8080,
			},
		},
		Password:       make([]byte, 0),
		CustomMetadata: map[string][]byte{},
	}
	meta, _ := json.MarshalIndent(metadata, "", "    ")

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  string(meta),
	}
	out = append(out, successOut, argError)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"Get the metadata of a ledger",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGet(vc)
	}, "the ledger id is not specified or the ledger id is specified more than one")
}

func doGet(vc *cmdutils.VerbCmd) error {
	id, err := strconv.ParseInt(vc.NameArg, 10, 64)
	if err != nil || id < 0 {
		return errors.Errorf("invalid ledger id %s", vc.NameArg)
	}

	admin := cmdutils.NewBookieClient()
	metadata, err := admin.Ledger().Get(id)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), metadata)
	}

	return err
}
