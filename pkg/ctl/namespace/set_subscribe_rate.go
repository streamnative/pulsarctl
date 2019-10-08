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

func SetSubscribeRateCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for setting the default subscribe rate per consumer of a namespace."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	setBySub := pulsar.Example{
		Desc:    "Set the default subscribe rate by subscribe of the namespace (namespace-name) (rate)",
		Command: "pulsarctl namespaces set-subscribe-rate --subscribe-rate (rate) (namespace)",
	}

	setByTime := pulsar.Example{
		Desc:    "Set the default subscribe rate by time of the namespace (namespace-name) (period)",
		Command: "pulsarctl namespaces set-subscribe-rate --period (period) (namespace)",
	}
	examples = append(examples, setBySub, setByTime)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Success set the default subscribe rate of the namespace (namespace-name) to (rate)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-subscribe-rate",
		"Set the default subscribe rate per consumer of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var rate pulsar.SubscribeRate

	vc.SetRunFuncWithNameArg(func() error {
		return doSetSubscribeRate(vc, rate)
	})

	vc.FlagSetGroup.InFlagSet("Subscription Rate", func(set *pflag.FlagSet) {
		set.IntVarP(&(rate.SubscribeThrottlingRatePerConsumer), "subscribe-rate", "m", -1,
			"message dispatch rate (default -1)")
		set.IntVarP(&(rate.RatePeriodInSecond), "period", "p", 30,
			"dispatch rate period (default 30 second)")
	})
}

func doSetSubscribeRate(vc *cmdutils.VerbCmd, rate pulsar.SubscribeRate) error {
	ns, err := pulsar.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetSubscribeRate(*ns, rate)
	if err == nil {
		vc.Command.Printf("Success set the default subscribe rate of the namespace %s to %+v", ns.String(), rate)
	}

	return err
}
