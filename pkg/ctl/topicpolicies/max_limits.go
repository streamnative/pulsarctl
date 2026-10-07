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

func getIntPolicyCmd(vc *cmdutils.VerbCmd, use, short string, getter func(admin bool, applied bool) (int, error)) {
	var global bool
	var applied bool
	vc.SetDescription(use, short, short, "", use)
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		value, err := getter(global, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, value, "%d\n", value)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func setIntPolicyCmd(vc *cmdutils.VerbCmd, use, short, flagName string, setter func(global bool, value int) error) {
	var global bool
	var value int
	vc.SetDescription(use, short, short, "", use)
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("Policy", func(set *pflag.FlagSet) {
		set.IntVarP(&value, flagName, "m", 0, flagName)
	})
	vc.SetRunFuncWithNameArg(func() error {
		err := setter(global, value)
		if err == nil {
			vc.Command.Printf("%s successfully for [%s]\n", short, vc.NameArg)
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func removePolicyCmd(vc *cmdutils.VerbCmd, use, short string, remover func(global bool) error) {
	var global bool
	vc.SetDescription(use, short, short, "", use)
	addScopeFlags(vc, &global, nil)
	vc.SetRunFuncWithNameArg(func() error {
		err := remover(global)
		if err == nil {
			vc.Command.Printf("%s successfully for [%s]\n", short, vc.NameArg)
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func GetMaxMessageSizeCmd(vc *cmdutils.VerbCmd) {
	getIntPolicyCmd(vc, "get-max-message-size", "Get max message size for a topic", func(global bool, applied bool) (int, error) {
		topic, err := topicName(vc)
		if err != nil {
			return 0, err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return 0, err
		}
		value, err := policies.GetMaxMessageSize(vc.Command.Context(), *topic, applied)
		if err != nil {
			return 0, err
		}
		if value == nil {
			return -1, nil
		}
		return *value, nil
	})
}

func SetMaxMessageSizeCmd(vc *cmdutils.VerbCmd) {
	setIntPolicyCmd(vc, "set-max-message-size", "Set max message size for a topic", "max-message-size", func(global bool, value int) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.SetMaxMessageSize(vc.Command.Context(), *topic, value)
	})
}

func RemoveMaxMessageSizeCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-max-message-size", "Removed max message size for a topic", func(global bool) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.RemoveMaxMessageSize(vc.Command.Context(), *topic)
	})
}

func GetMaxSubscriptionsPerTopicCmd(vc *cmdutils.VerbCmd) {
	getIntPolicyCmd(vc, "get-max-subscriptions-per-topic", "Get max subscriptions per topic", func(global bool, applied bool) (int, error) {
		topic, err := topicName(vc)
		if err != nil {
			return 0, err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return 0, err
		}
		value, err := policies.GetMaxSubscriptionsPerTopic(vc.Command.Context(), *topic, applied)
		if err != nil {
			return 0, err
		}
		if value == nil {
			return -1, nil
		}
		return *value, nil
	})
}

func SetMaxSubscriptionsPerTopicCmd(vc *cmdutils.VerbCmd) {
	setIntPolicyCmd(vc, "set-max-subscriptions-per-topic", "Set max subscriptions per topic", "max-subscriptions-per-topic", func(global bool, value int) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.SetMaxSubscriptionsPerTopic(vc.Command.Context(), *topic, value)
	})
}

func RemoveMaxSubscriptionsPerTopicCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-max-subscriptions-per-topic", "Removed max subscriptions per topic", func(global bool) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.RemoveMaxSubscriptionsPerTopic(vc.Command.Context(), *topic)
	})
}

func GetDeduplicationSnapshotIntervalCmd(vc *cmdutils.VerbCmd) {
	getIntPolicyCmd(vc, "get-deduplication-snapshot-interval", "Get deduplication snapshot interval", func(global bool, applied bool) (int, error) {
		topic, err := topicName(vc)
		if err != nil {
			return 0, err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return 0, err
		}
		value, err := policies.GetDeduplicationSnapshotInterval(vc.Command.Context(), *topic, applied)
		if err != nil {
			return 0, err
		}
		if value == nil {
			return -1, nil
		}
		return *value, nil
	})
}

func SetDeduplicationSnapshotIntervalCmd(vc *cmdutils.VerbCmd) {
	setIntPolicyCmd(vc, "set-deduplication-snapshot-interval", "Set deduplication snapshot interval", "interval", func(global bool, value int) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.SetDeduplicationSnapshotInterval(vc.Command.Context(), *topic, value)
	})
}

func RemoveDeduplicationSnapshotIntervalCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-deduplication-snapshot-interval", "Removed deduplication snapshot interval", func(global bool) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.RemoveDeduplicationSnapshotInterval(vc.Command.Context(), *topic)
	})
}

func GetReplicatorDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-replicator-dispatch-rate", "Get replicator dispatch rate for a topic", "Get replicator dispatch rate for a topic", "")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		rate, err := policies.GetReplicatorDispatchRate(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, rate, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetReplicatorDispatchRateCmd(vc *cmdutils.VerbCmd) {
	var global bool
	data := utils.DispatchRateData{}
	vc.SetDescription("set-replicator-dispatch-rate", "Set replicator dispatch rate for a topic", "Set replicator dispatch rate for a topic", "")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("ReplicatorDispatchRate", func(set *pflag.FlagSet) {
		set.Int64VarP(&data.DispatchThrottlingRateInMsg, "msg-dispatch-rate", "", -1, "message dispatch rate")
		set.Int64VarP(&data.DispatchThrottlingRateInByte, "byte-dispatch-rate", "", -1, "byte dispatch rate")
		set.Int64VarP(&data.RatePeriodInSecond, "dispatch-rate-period", "", 1, "dispatch rate period in seconds")
		set.BoolVarP(&data.RelativeToPublishRate, "relative-to-publish-rate", "", false, "relative to publish rate")
	})
	vc.SetRunFuncWithNameArg(func() error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		err = policies.SetReplicatorDispatchRate(vc.Command.Context(), *topic, data)
		if err == nil {
			vc.Command.Printf("Set replicator dispatch rate successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveReplicatorDispatchRateCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-replicator-dispatch-rate", "Removed replicator dispatch rate", func(global bool) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.RemoveReplicatorDispatchRate(vc.Command.Context(), *topic)
	})
}
