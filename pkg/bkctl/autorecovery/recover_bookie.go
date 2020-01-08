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

func recoverBookieCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for recovering the ledger data of a failed bookie."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	rb := cmdutils.Example{
		Desc:    "Recover the ledger data of a failed bookie",
		Command: "pulsarctl bookkeeper autorecovery recoverbookie (bookie-1) (bookie-2)",
	}
	examples = append(examples, rb)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully recover the bookies (bookie-1) (bookie-2)",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"recoverbookie",
		"Recover the ledger data of a failed bookie",
		desc.ToString(),
		desc.ExampleToString())

	var deleteCookie bool

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doRecoverBookie(vc, deleteCookie)
	}, func(args []string) error {
		return nil
	})

	vc.FlagSetGroup.InFlagSet("Recover Bookie", func(set *pflag.FlagSet) {
		set.BoolVar(&deleteCookie, "delelte-cookie", false, "delete cookie")
	})
}

func doRecoverBookie(vc *cmdutils.VerbCmd, deleteCookie bool) error {
	admin := cmdutils.NewBookieClient()
	err := admin.AutoRecovery().RecoverBookie(vc.NameArgs, deleteCookie)
	if err == nil {
		if deleteCookie {
			vc.Command.Printf("Successfully recover the bookies %v and delete the cookie\n", vc.NameArgs)
		} else {
			vc.Command.Printf("Successfully recover the bookie %v\n", vc.NameArgs)
		}
	}

	return err
}
