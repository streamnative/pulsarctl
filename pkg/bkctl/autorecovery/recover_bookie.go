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
	"errors"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/spf13/pflag"
)

func recoverBookieCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for recovering the ledger data of a failed bookie."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	rb := cmdutils.Example{
		Desc:    "Recover the ledger data of a failed bookie.",
		Command: "pulsarctl bookkeeper auto-recovery recover-bookie (bookie-1) (bookie-2)",
	}
	examples = append(examples, rb)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Recover the bookies successfully.",
		Out:  "Successfully recover the bookies (bookie-1) (bookie-2).",
	}

	IDNotSpecified := cmdutils.Output{
		Desc: "The recover bookie id is not specified.",
		Out:  "[✖]  you need to specify the recover bookies id",
	}
	out = append(out, successOut, IDNotSpecified)
	desc.CommandOutput = out

	vc.SetDescription(
		"recover-bookie",
		"Recover the ledger data of a failed bookie.",
		desc.ToString(),
		desc.ExampleToString())

	var deleteCookie bool

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doRecoverBookie(vc, deleteCookie)
	}, func(args []string) error {
		if len(args) == 0 {
			return errors.New("you need to specify the recover bookies id")
		}
		return nil
	})

	vc.FlagSetGroup.InFlagSet("Recover bookie", func(set *pflag.FlagSet) {
		set.BoolVar(&deleteCookie, "delete-cookie", false,
			"Delete cookie when recovering the failed bookies.")
	})
}

func doRecoverBookie(vc *cmdutils.VerbCmd, deleteCookie bool) error {
	admin := cmdutils.NewBookieClient()
	err := admin.AutoRecovery().RecoverBookie(vc.NameArgs, deleteCookie)
	if err == nil {
		if deleteCookie {
			vc.Command.Printf("Successfully recover the bookies %v and delete the cookie.\n", vc.NameArgs)
		} else {
			vc.Command.Printf("Successfully recover the bookie %v.\n", vc.NameArgs)
		}
	}

	return err
}
