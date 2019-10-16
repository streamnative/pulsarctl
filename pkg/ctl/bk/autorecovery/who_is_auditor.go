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
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func WhoIsAuditorCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for getting who is the auditor."
	desc.CommandPermission = "none"

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Get who is the auditor",
		Command: "pulsarctl bk auto-recovery who-is-auditor",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: `{
    "Auditor": "hostname/hostAddress:Port"
}`,
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"who-is-auditor",
		"Get who is the auditor",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doWhoIsAuditor(vc)
	})
}

func doWhoIsAuditor(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	auditor, err := admin.AutoRecovery().WhoIsAuditor()
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), auditor)
	}

	return err
}
