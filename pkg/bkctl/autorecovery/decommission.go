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
)

func decommissionCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for decommissioning a bookie."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	c := cmdutils.Example{
		Desc:    "Decommission a bookie.",
		Command: "pulsarctl bookkeeper auto-recovery decommission (bk-ip:bk-port)",
	}
	examples = append(examples, c)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Successfully decommission a bookie.",
		Out:  "Successfully decommission the bookie (bookie-ip:bookie-port)",
	}

	argError := cmdutils.Output{
		Desc: "The bookie address is not specified or the bookie address is specified more than one.",
		Out:  "[✖]  the bookie address is not specified or the bookie address is specified more than one",
	}
	out = append(out, successOut, argError)
	desc.CommandOutput = out

	vc.SetDescription(
		"decommission",
		"Decommission a bookie.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doDecommission(vc)
	}, "the bookie address is not specified or the bookie address is specified more than one")
}

func doDecommission(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	err := admin.AutoRecovery().Decommission(vc.NameArg)
	if err == nil {
		vc.Command.Printf("Successfully decommission the bookie %s.\n", vc.NameArg)
	}

	return err
}
