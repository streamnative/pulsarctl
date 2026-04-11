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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetSubscribeRateCmd(vc *cmdutils.VerbCmd) {
	vc.SetDescription("get-subscribe-rate", "Get subscribe rate for a topic", "Get subscribe rate for a topic", "", "get-subscribe-rate")
	vc.SetRunFuncWithNameArg(func() error {
		return doGetSubscribeRate(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
	vc.EnableOutputFlagSet()
}

func doGetSubscribeRate(vc *cmdutils.VerbCmd) error {
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	rate, err := admin.Topics().GetSubscribeRate(*topic)
	if err == nil {
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), cmdutils.NewOutputContent().WithObject(rate))
	}
	return err
}

func SetSubscribeRateCmd(vc *cmdutils.VerbCmd) {
	data := *utils.NewSubscribeRate()
	vc.SetDescription("set-subscribe-rate", "Set subscribe rate for a topic", "Set subscribe rate for a topic", "", "set-subscribe-rate")
	vc.FlagSetGroup.InFlagSet("SubscribeRate", func(set *pflag.FlagSet) {
		set.IntVarP(&data.SubscribeThrottlingRatePerConsumer, "subscribe-rate", "m", -1, "message dispatch rate")
		set.IntVarP(&data.RatePeriodInSecond, "period", "p", 30, "dispatch rate period")
	})
	vc.SetRunFuncWithNameArg(func() error {
		return doSetSubscribeRate(vc, data)
	}, "the topic name is not specified or the topic name is specified more than one")
	vc.EnableOutputFlagSet()
}

func doSetSubscribeRate(vc *cmdutils.VerbCmd, rate utils.SubscribeRate) error {
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetSubscribeRate(*topic, rate)
	if err == nil {
		vc.Command.Printf("Set subscribe rate successfully for [%s]\n", topic.String())
	}
	return err
}

func RemoveSubscribeRateCmd(vc *cmdutils.VerbCmd) {
	vc.SetDescription("remove-subscribe-rate", "Remove subscribe rate for a topic", "Remove subscribe rate for a topic", "", "remove-subscribe-rate")
	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveSubscribeRate(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doRemoveSubscribeRate(vc *cmdutils.VerbCmd) error {
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().RemoveSubscribeRate(*topic)
	if err == nil {
		vc.Command.Printf("Removed subscribe rate successfully for [%s]\n", topic.String())
	}
	return err
}
