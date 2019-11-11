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

package brokerstats

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func dumpMBeans(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Dump the mbean stats"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Dump the mbean stats",
		Command: "pulsarctl broker-stats mbeans",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Get all the mbean details of this broker JVM",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"mbeans",
		"Dump the mbean stats",
		desc.ToString(),
		desc.ExampleToString(),
		"mbeans")

	vc.SetRunFunc(func() error {
		return doDumpMBeans(vc)
	})
}

func doDumpMBeans(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	metrics, err := admin.BrokerStats().GetMBeans()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), metrics)
	}
	return err
}
