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

package namespace

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
)

func GetDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var desc LongDescription
	desc.CommandUsedFor = "This command is used to get message-dispatch-rate for all topics of a namespace."
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []Example
	get := Example{
		Desc:    "Get the message-dispatch-rate for all topics of namespace <namespace-name>",
		Command: "pulsarctl namespaces get-dispatch-rate <namespace",
	}
	desc.CommandExamples = append(examples, get)

	var out []Output
	successOut := Output{
		Desc: "normal output",
		Out:  "{\n  \"dispatchThrottlingRateInMsg\" : 0,\n  \"dispatchThrottlingRateInByte\" : 0,\n  \"ratePeriodInSecond\" : 1\n}",
	}
	out = append(out, successOut)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-dispatch-rate",
		"Get the message-dispatch-rate for all topics of a namespace",
		desc.ToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetDispatchRate(vc)
	})
}

func doGetDispatchRate(vc *cmdutils.VerbCmd) error {
	ns, err := GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	rate, err := admin.Namespaces().GetDispatchRate(*ns)
	if err == nil {
		cmdutils.PrintJson(vc.Command.OutOrStdout(), rate)
	}

	return err
}
