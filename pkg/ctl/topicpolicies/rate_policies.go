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

package topicpolicies

import (
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-dispatch-rate", "Get message dispatch rate for a topic", "Get message dispatch rate for a topic", "", "get-dispatch-rate")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		rate, err := policies.GetDispatchRate(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, rate, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	data := utils.DispatchRateData{}
	vc.SetDescription("set-dispatch-rate", "Set message dispatch rate for a topic", "Set message dispatch rate for a topic", "", "set-dispatch-rate")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("DispatchRate", func(set *pflag.FlagSet) {
		set.Int64VarP(&data.DispatchThrottlingRateInMsg, "msg-dispatch-rate", "", -1, "message dispatch rate")
		set.Int64VarP(&data.DispatchThrottlingRateInByte, "byte-dispatch-rate", "", -1, "byte dispatch rate")
		set.Int64VarP(&data.RatePeriodInSecond, "dispatch-rate-period", "", 1, "dispatch rate period in seconds")
		set.BoolVarP(&data.RelativeToPublishRate, "relative-to-publish-rate", "", false, "relative to publish rate")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetDispatchRate(vc.Command.Context(), *topic, data)
		if err == nil {
			vc.Command.Printf("Set dispatch rate successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveDispatchRateCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-dispatch-rate", "Removed dispatch rate for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveDispatchRate(vc.Command.Context(), *topic)
	})
}

func GetSubscriptionDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-subscription-dispatch-rate", "Get subscription dispatch rate for a topic", "Get subscription dispatch rate for a topic", "", "get-subscription-dispatch-rate")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		rate, err := policies.GetSubscriptionDispatchRate(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, rate, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetSubscriptionDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	data := utils.DispatchRateData{}
	vc.SetDescription("set-subscription-dispatch-rate", "Set subscription dispatch rate for a topic", "Set subscription dispatch rate for a topic", "", "set-subscription-dispatch-rate")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("SubscriptionDispatchRate", func(set *pflag.FlagSet) {
		set.Int64VarP(&data.DispatchThrottlingRateInMsg, "msg-dispatch-rate", "", -1, "message dispatch rate")
		set.Int64VarP(&data.DispatchThrottlingRateInByte, "byte-dispatch-rate", "", -1, "byte dispatch rate")
		set.Int64VarP(&data.RatePeriodInSecond, "dispatch-rate-period", "", 1, "dispatch rate period in seconds")
		set.BoolVarP(&data.RelativeToPublishRate, "relative-to-publish-rate", "", false, "relative to publish rate")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetSubscriptionDispatchRate(vc.Command.Context(), *topic, data)
		if err == nil {
			vc.Command.Printf("Set subscription dispatch rate successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveSubscriptionDispatchRateCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-subscription-dispatch-rate", "Removed subscription dispatch rate for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveSubscriptionDispatchRate(vc.Command.Context(), *topic)
	})
}

func GetPublishRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-publish-rate", "Get publish rate for a topic", "Get publish rate for a topic", "", "get-publish-rate")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		rate, err := policies.GetPublishRate(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, rate, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetPublishRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	data := utils.PublishRateData{}
	vc.SetDescription("set-publish-rate", "Set publish rate for a topic", "Set publish rate for a topic", "", "set-publish-rate")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("PublishRate", func(set *pflag.FlagSet) {
		set.Int64VarP(&data.PublishThrottlingRateInMsg, "msg-publish-rate", "", -1, "message publish rate")
		set.Int64VarP(&data.PublishThrottlingRateInByte, "byte-publish-rate", "", -1, "byte publish rate")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetPublishRate(vc.Command.Context(), *topic, data)
		if err == nil {
			vc.Command.Printf("Set publish rate successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemovePublishRateCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-publish-rate", "Removed publish rate for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemovePublishRate(vc.Command.Context(), *topic)
	})
}

func GetSubscribeRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-subscribe-rate", "Get subscribe rate for a topic", "Get subscribe rate for a topic", "", "get-subscribe-rate")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		rate, err := policies.GetSubscribeRate(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, rate, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetSubscribeRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	data := *utils.NewSubscribeRate()
	vc.SetDescription("set-subscribe-rate", "Set subscribe rate for a topic", "Set subscribe rate for a topic", "", "set-subscribe-rate")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("SubscribeRate", func(set *pflag.FlagSet) {
		set.IntVarP(&data.SubscribeThrottlingRatePerConsumer, "subscribe-rate", "m", -1, "message dispatch rate")
		set.IntVarP(&data.RatePeriodInSecond, "period", "p", 30, "dispatch rate period")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetSubscribeRate(vc.Command.Context(), *topic, data)
		if err == nil {
			vc.Command.Printf("Set subscribe rate successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveSubscribeRateCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-subscribe-rate", "Removed subscribe rate for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveSubscribeRate(vc.Command.Context(), *topic)
	})
}
