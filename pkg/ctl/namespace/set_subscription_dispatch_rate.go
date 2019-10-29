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
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func SetSubscriptionDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for setting the default subscription message dispatch rate of a namespace."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	setByMsg := pulsar.Example{
		Desc:    "Set the default subscription message dispatch rate by message of the namespace (namespace-name) to (rate)",
		Command: "pulsarctl namespaces set-subscription-dispatch-rate --msg-rate <rate> <namespace",
	}

	setByByte := pulsar.Example{
		Desc:    "Set the default subscription message dispatch rate by byte of the namespace (namespace-name) to (rate)",
		Command: "pulsarctl namespaces set-subscription-dispatch-rate --byte-rate (rate) (namespace)",
	}

	setByTime := pulsar.Example{
		Desc:    "Set the default subscriptions message dispatch rate by time of the namespace (namespace-name) to (rate)",
		Command: "pulsarctl namespaces set-subscription-dispatch-rate --period (period) (namespace)",
	}
	examples = append(examples, setByMsg, setByByte, setByTime)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Success set the default subscription message dispatch rate of the namespace (namespace-name) to (rate)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-subscription-dispatch-rate",
		"Set the default subscription message dispatch rate of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var rate pulsar.DispatchRate

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSubscriptionDispatchRate(vc, rate)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Subscription Dispatch Rate", func(set *pflag.FlagSet) {
		set.IntVarP(&(rate.DispatchThrottlingRateInMsg), "msg-rate", "m", -1,
			"message dispatch rate (default -1)")
		set.Int64VarP(&(rate.DispatchThrottlingRateInByte), "byte-rate", "b", -1,
			"byte dispatch rate (default -1)")
		set.IntVarP(&(rate.RatePeriodInSecond), "period", "p", 1,
			"dispatch rate period (default 1 second)")
	})
}

func doSetSubscriptionDispatchRate(vc *cmdutils.VerbCmd, rate pulsar.DispatchRate) error {
	ns, err := pulsar.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSubscriptionDispatchRate(*ns, rate)
	if err == nil {
		vc.Command.Printf("Success set the default subscription message dispatch rate "+
			"of the namespace %s to %+v\n", ns.String(), rate)
	}

	return err
}
