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

package topic

import (
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	ctlUtil "github.com/streamnative/pulsarctl/pkg/ctl/utils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func SetDelayedDeliveryCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set delayed delivery policy for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Set delayed delivery policy for a topic",
		Command: "pulsarctl topics set-delayed-delivery topic -t 10s -e",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set delayed delivery policy successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-delayed-delivery",
		"Set delayed delivery policy for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"set-delayed-delivery",
	)
	delayedDeliveryData := &utils.DelayedDeliveryCmdData{}
	vc.SetRunFuncWithNameArg(func() error {
		return doSetDelayedDelivery(vc, delayedDeliveryData)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Persistence", func(set *pflag.FlagSet) {
		set.BoolVarP(
			&delayedDeliveryData.Enable,
			"enable",
			"e",
			false,
			"Enable delayed delivery messages")
		set.BoolVarP(
			&delayedDeliveryData.Disable,
			"disable",
			"d",
			false,
			"Disable delayed delivery messages")
		set.StringVarP(
			&delayedDeliveryData.DelayedDeliveryTimeStr,
			"time",
			"t",
			"1s",
			"The tick time for when retrying on delayed delivery messages, affecting the"+
				" accuracy of the delivery time compared to the scheduled time. (eg: 1s, 10s, 1m, 5h, 3d)")
	})
	vc.EnableOutputFlagSet()
}

func doSetDelayedDelivery(vc *cmdutils.VerbCmd, delayedDeliveryCmdData *utils.DelayedDeliveryCmdData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	delayedDeliveryData := &utils.DelayedDeliveryData{}
	if delayedDeliveryCmdData.Enable == delayedDeliveryCmdData.Disable {
		msg := "Need to specify either --enable or --disable"
		vc.Command.Printf(msg)
		return errors.Errorf(msg)
	}
	if delayedDeliveryCmdData.Enable {
		retentionTimeInSecond, err := ctlUtil.ParseRelativeTimeInSeconds(delayedDeliveryCmdData.DelayedDeliveryTimeStr)
		if err != nil {
			return err
		}
		delayedDeliveryData.TickTime = tickTimeInSecond.Seconds()
		delayedDeliveryData.Active = true
	}
	err = admin.Topics().SetDelayedDelivery(*topic, *delayedDeliveryData)
	if err == nil {
		vc.Command.Printf("Set delayed delivery policy successfully for [%s]\n", topic.String())
	}
	return err
}
