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

package functionsworker

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getCluster(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get all workers belonging to this cluster"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	stats := cmdutils.Example{
		Desc:    "Get all workers belonging to this cluster",
		Command: "pulsarctl functions-worker get-cluster",
	}
	examples = append(examples, stats)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "[\n" +
			"  {\n" +
			"    \"workerId\": \"c-standalone-fw-127.0.0.1-8080\",\n" +
			"    \"workerHostname\": \"127.0.0.1\",\n" +
			"    \"port\": 8080\n" +
			"  }\n" +
			"]",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-cluster",
		"Get all workers belonging to this cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"get-cluster",
	)

	// set the run function
	vc.SetRunFunc(func() error {
		return doGetCluster(vc)
	})
}

func doGetCluster(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	workersInfo, err := admin.FunctionsWorker().GetCluster()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), workersInfo)
	}

	return err
}
