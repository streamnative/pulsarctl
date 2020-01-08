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

func getLostBookieRecoveryDelayCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the lost bookie recovery delay in second of a bookie."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the lost Bookie Recovery Delay of a bookie.",
		Command: "pulsarctl bookkeeper auto-recovery get-lost-bookie-recovery-delay",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Get the lost bookie recovery delay of a bookie. ",
		Out:  "lostBookieRecoveryDelay value: (delay)",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-lost-bookie-recovery-delay",
		"Get the lost bookie recovery delay of a bookie.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doGetLostBookieRecoveryDelay(vc)
	})
}

func doGetLostBookieRecoveryDelay(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	out, err := admin.AutoRecovery().GetLostBookieRecoveryDelay()
	if err == nil {
		vc.Command.Println(out)
	}

	return err
}
