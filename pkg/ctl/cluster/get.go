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

package cluster

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getClusterDataCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for getting the cluster data of the specified cluster."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	get := cmdutils.Example{
		Desc:    "getting the (cluster-name) data",
		Command: "pulsarctl clusters get (cluster-name)",
	}
	examples = append(examples, get)

	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n  " +
			"\"serviceUrl\": \"http://localhost:8080\",\n  " +
			"\"serviceUrlTls\": \"\",\n  " +
			"\"brokerServiceUrl\": \"pulsar://localhost:6650\",\n  " +
			"\"brokerServiceUrlTls\": \"\",\n  " +
			"\"peerClusterNames\": null\n" +
			"}",
	}
	out = append(out, successOut)
	out = append(out, argsError)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"get",
		"Get the configuration data for the specified cluster",
		desc.ToString(),
		desc.ExampleToString(),
		"get")

	vc.SetRunFuncWithNameArg(func() error {
		return doGetClusterData(vc)
	}, "the cluster name is not specified or the cluster name is specified more than one")
}

func doGetClusterData(vc *cmdutils.VerbCmd) error {
	clusterName := vc.NameArg
	if clusterName == "" {
		return errors.New("should specified a cluster name")
	}

	admin := cmdutils.NewPulsarClient()
	clusterData, err := admin.Clusters().Get(clusterName)
	if err == nil {
		s, err := json.MarshalIndent(clusterData, "", "    ")
		if err != nil {
			return err
		}
		vc.Command.Println(string(s))
	}

	return err
}
