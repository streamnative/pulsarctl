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

package bookies

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func diskUsageInfoCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting the bookie disk usage info of a cluster."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "Get the bookie disk usage info of a cluster.",
		Command: "pulsar bookkeeper bookie disk-usage-info",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "Successfully get the bookie disk usage info of a cluster.",
		Out: `{
    "bookieAddress" : {free: xxx, total: xxx},
    "bookieAddress" : {free: xxx, total: xxx},
    ...
    "clusterInfo" : {total_free: xxx, total: xxx}
}`,
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"disk-usage-info",
		"Get the bookie disk usage info of a cluster.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doGetInfo(vc)
	})
}

func doGetInfo(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewBookieClient()
	info, err := admin.Bookies().DiskUsageInfo()
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), info)
	}

	return err
}
