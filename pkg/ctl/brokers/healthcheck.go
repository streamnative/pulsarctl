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

package brokers

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func healthCheckCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Run a health check against the broker"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "Run a health check against the broker",
		Command: "pulsarctl brokers healthcheck",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "ok",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"healthcheck",
		"Run a health check against the broker",
		desc.ToString(),
		desc.ExampleToString(),
		"healthcheck")

	vc.SetRunFunc(func() error {
		return doHealthCheck(vc)
	})
}

func doHealthCheck(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	err := admin.Brokers().HealthCheck()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Println("ok")
	}
	return err
}
