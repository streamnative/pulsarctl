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

func triggerAuditCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for triggering audit by resetting the lostBookieRecoveryDelay."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	trigger := cmdutils.Example{
		Desc:    "Trigger audit by resetting the lostBookieRecoveryDelay",
		Command: "pulsarctl bookkeeper autorecovery triggeraudit",
	}
	examples = append(examples, trigger)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully trigger audit by resetting the lostBookieRecoveryDelay",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"triggeraudit",
		"Trigger audit by resetting the lostBookieRecoveryDelay",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doTriggerAudit(vc)
	})
}

func doTriggerAudit(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	err := admin.AutoRecovery().TriggerAudit()
	if err == nil {
		vc.Command.Println("Successfully trigger audit by resetting the lostBookieRecoveryDelay")
	}

	return err
}
