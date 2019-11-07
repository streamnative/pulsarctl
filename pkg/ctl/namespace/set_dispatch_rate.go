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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func SetDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the default message dispatch rate of a namespace."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	setByMsg := cmdutils.Example{
		Desc:    "Set the default message dispatch rate by message of the namespace (namespace-name) to (rate)",
		Command: "pulsarctl namespaces set-dispatch-rate --msg-rate (rate) (namespace)",
	}

	setByByte := cmdutils.Example{
		Desc:    "Set the default message dispatch rate by byte of the namespace (namespace-name) to (rate)",
		Command: "pulsarctl namespaces set-dispatch-rate --byte-rate (rate) (namespace)",
	}

	setByTime := cmdutils.Example{
		Desc:    "Set the default message dispatch rate by time of the namespace (namespace-name) to (period)",
		Command: "pulsarctl namespaces set-dispatch-rate --period (period) (namespace)",
	}
	examples = append(examples, setByMsg, setByByte, setByTime)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Success set the default message dispatch rate of the namespace (namespace-name) to (rate)",
	}
	out = append(out, successOut)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-dispatch-rate",
		"Set the default message dispatch rate of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var rate utils.DispatchRate

	vc.SetRunFuncWithNameArg(func() error {
		return doSetDispatchRate(vc, rate)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Dispatch Rate", func(set *pflag.FlagSet) {
		set.IntVarP(&(rate.DispatchThrottlingRateInMsg), "msg-rate", "m", -1,
			"message dispatch rate (default -1)")
		set.Int64VarP(&(rate.DispatchThrottlingRateInByte), "byte-rate", "b", -1,
			"byte dispatch rate (default -1)")
		set.IntVarP(&(rate.RatePeriodInSecond), "period", "p", 1,
			"dispatch rate period (default 1 second)")
	})
}

func doSetDispatchRate(vc *cmdutils.VerbCmd, rate utils.DispatchRate) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetDispatchRate(*ns, rate)
	if err == nil {
		vc.Command.Printf("Success set the default message dispatch rate "+
			"of the namespace %s to %+v\n", ns.String(), rate)
	}

	return err
}
