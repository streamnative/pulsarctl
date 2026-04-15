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
	"context"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	ctlutils "github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

func GetMessageTTLCmd(vc *cmdutils.VerbCmd) {
	getOptionalIntPolicyCmd(
		vc,
		"get-message-ttl",
		"Get message TTL for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*int, error) {
			return policies.GetMessageTTL(ctx, topic, applied)
		},
	)
}

func SetMessageTTLCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var ttl string
	vc.SetDescription("set-message-ttl", "Set message TTL for a topic", "Set message TTL for a topic", "", "set-message-ttl")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("MessageTTL", func(set *pflag.FlagSet) {
		set.StringVarP(&ttl, "ttl", "t", "", "message TTL for topic with optional time unit suffix")
		_ = cobra.MarkFlagRequired(set, "ttl")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		seconds, err := ctlutils.ParseRelativeTimeInSeconds(ttl)
		if err != nil {
			return err
		}
		messageTTL := -1
		if seconds != -1 {
			messageTTL = int(seconds.Seconds())
		}
		err = policies.SetMessageTTL(vc.Command.Context(), *topic, messageTTL)
		if err == nil {
			vc.Command.Printf("Set message TTL successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveMessageTTLCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-message-ttl", "Removed message TTL for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveMessageTTL(vc.Command.Context(), *topic)
	})
}

func GetMaxUnackMessagesPerConsumerCmd(vc *cmdutils.VerbCmd) {
	getOptionalIntPolicyCmd(
		vc,
		"get-max-unacked-messages-per-consumer",
		"Get max unacked messages per consumer for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*int, error) {
			return policies.GetMaxUnackMessagesPerConsumer(ctx, topic, applied)
		},
	)
}

func SetMaxUnackMessagesPerConsumerCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var maxNum int
	vc.SetDescription("set-max-unacked-messages-per-consumer", "Set max unacked messages per consumer for a topic", "Set max unacked messages per consumer for a topic", "", "set-max-unacked-messages-per-consumer")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("MaxUnackedMessagesPerConsumer", func(set *pflag.FlagSet) {
		set.IntVarP(&maxNum, "maxNum", "m", 0, "max unacked messages num on consumer")
		_ = cobra.MarkFlagRequired(set, "maxNum")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetMaxUnackMessagesPerConsumer(vc.Command.Context(), *topic, maxNum)
		if err == nil {
			vc.Command.Printf("Set max unacked messages per consumer successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveMaxUnackMessagesPerConsumerCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-max-unacked-messages-per-consumer", "Removed max unacked messages per consumer for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveMaxUnackMessagesPerConsumer(vc.Command.Context(), *topic)
	})
}

func GetMaxUnackMessagesPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	getOptionalIntPolicyCmd(
		vc,
		"get-max-unacked-messages-per-subscription",
		"Get max unacked messages per subscription for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*int, error) {
			return policies.GetMaxUnackMessagesPerSubscription(ctx, topic, applied)
		},
	)
}

func SetMaxUnackMessagesPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var maxNum int
	vc.SetDescription("set-max-unacked-messages-per-subscription", "Set max unacked messages per subscription for a topic", "Set max unacked messages per subscription for a topic", "", "set-max-unacked-messages-per-subscription")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("MaxUnackedMessagesPerSubscription", func(set *pflag.FlagSet) {
		set.IntVarP(&maxNum, "maxNum", "m", 0, "max unacked messages num on subscription")
		_ = cobra.MarkFlagRequired(set, "maxNum")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetMaxUnackMessagesPerSubscription(vc.Command.Context(), *topic, maxNum)
		if err == nil {
			vc.Command.Printf("Set max unacked messages per subscription successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveMaxUnackMessagesPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-max-unacked-messages-per-subscription", "Removed max unacked messages per subscription for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveMaxUnackMessagesPerSubscription(vc.Command.Context(), *topic)
	})
}

func GetMaxConsumersPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	getOptionalIntPolicyCmd(
		vc,
		"get-max-consumers-per-subscription",
		"Get max consumers per subscription for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*int, error) {
			return policies.GetMaxConsumersPerSubscription(ctx, topic, applied)
		},
	)
}

func SetMaxConsumersPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var maxConsumers int
	vc.SetDescription("set-max-consumers-per-subscription", "Set max consumers per subscription for a topic", "Set max consumers per subscription for a topic", "", "set-max-consumers-per-subscription")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("MaxConsumersPerSubscription", func(set *pflag.FlagSet) {
		set.IntVarP(&maxConsumers, "max-consumers-per-subscription", "c", 0, "max consumers per subscription for a topic")
		_ = cobra.MarkFlagRequired(set, "max-consumers-per-subscription")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetMaxConsumersPerSubscription(vc.Command.Context(), *topic, maxConsumers)
		if err == nil {
			vc.Command.Printf("Set max consumers per subscription successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveMaxConsumersPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-max-consumers-per-subscription", "Removed max consumers per subscription for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveMaxConsumersPerSubscription(vc.Command.Context(), *topic)
	})
}

func GetMaxConsumersCmd(vc *cmdutils.VerbCmd) {
	getOptionalIntPolicyCmd(
		vc,
		"get-max-consumers",
		"Get max consumers for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*int, error) {
			return policies.GetMaxConsumers(ctx, topic, applied)
		},
	)
}

func SetMaxConsumersCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var maxConsumers int
	vc.SetDescription("set-max-consumers", "Set max consumers for a topic", "Set max consumers for a topic", "", "set-max-consumers")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("MaxConsumers", func(set *pflag.FlagSet) {
		set.IntVarP(&maxConsumers, "max-consumers", "c", 0, "max consumers for a topic")
		_ = cobra.MarkFlagRequired(set, "max-consumers")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetMaxConsumers(vc.Command.Context(), *topic, maxConsumers)
		if err == nil {
			vc.Command.Printf("Set max consumers successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveMaxConsumersCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-max-consumers", "Removed max consumers for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveMaxConsumers(vc.Command.Context(), *topic)
	})
}

func GetMaxProducersCmd(vc *cmdutils.VerbCmd) {
	getOptionalIntPolicyCmd(
		vc,
		"get-max-producers",
		"Get max producers for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*int, error) {
			return policies.GetMaxProducers(ctx, topic, applied)
		},
	)
}

func SetMaxProducersCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var maxProducers int
	vc.SetDescription("set-max-producers", "Set max producers for a topic", "Set max producers for a topic", "", "set-max-producers")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("MaxProducers", func(set *pflag.FlagSet) {
		set.IntVarP(&maxProducers, "max-producers", "p", 0, "max producers for a topic")
		_ = cobra.MarkFlagRequired(set, "max-producers")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetMaxProducers(vc.Command.Context(), *topic, maxProducers)
		if err == nil {
			vc.Command.Printf("Set max producers successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveMaxProducersCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-max-producers", "Removed max producers for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveMaxProducers(vc.Command.Context(), *topic)
	})
}

func GetCompactionThresholdCmd(vc *cmdutils.VerbCmd) {
	getOptionalInt64PolicyCmd(
		vc,
		"get-compaction-threshold",
		"Get compaction threshold for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*int64, error) {
			return policies.GetCompactionThreshold(ctx, topic, applied)
		},
	)
}

func SetCompactionThresholdCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var threshold string
	vc.SetDescription("set-compaction-threshold", "Set compaction threshold for a topic", "Set compaction threshold for a topic", "", "set-compaction-threshold")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("CompactionThreshold", func(set *pflag.FlagSet) {
		set.StringVarP(&threshold, "threshold", "t", "", "maximum backlog before compaction is triggered")
		_ = cobra.MarkFlagRequired(set, "threshold")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		value, err := ctlutils.ValidateSizeString(threshold)
		if err != nil {
			return err
		}
		err = policies.SetCompactionThreshold(vc.Command.Context(), *topic, value)
		if err == nil {
			vc.Command.Printf("Set compaction threshold successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveCompactionThresholdCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-compaction-threshold", "Removed compaction threshold for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveCompactionThreshold(vc.Command.Context(), *topic)
	})
}

func GetDeduplicationCmd(vc *cmdutils.VerbCmd) {
	getOptionalBoolPolicyCmd(
		vc,
		"get-deduplication",
		"Get deduplication status for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*bool, error) {
			return policies.GetDeduplicationStatus(ctx, topic, applied)
		},
	)
}

func SetDeduplicationCmd(vc *cmdutils.VerbCmd) {
	setEnableDisablePolicyCmd(
		vc,
		"set-deduplication",
		"Set deduplication status for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, enabled bool) error {
			return policies.SetDeduplicationStatus(ctx, topic, enabled)
		},
	)
}

func RemoveDeduplicationCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-deduplication", "Removed deduplication status for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveDeduplicationStatus(vc.Command.Context(), *topic)
	})
}

func GetSchemaValidationEnforcedCmd(vc *cmdutils.VerbCmd) {
	getOptionalBoolPolicyCmd(
		vc,
		"get-schema-validation-enforced",
		"Get schema validation enforced for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, applied bool) (*bool, error) {
			return policies.GetSchemaValidationEnforced(ctx, topic, applied)
		},
	)
}

func SetSchemaValidationEnforcedCmd(vc *cmdutils.VerbCmd) {
	setEnableDisablePolicyCmd(
		vc,
		"set-schema-validation-enforced",
		"Set schema validation enforced for a topic",
		func(ctx context.Context, policies admin.TopicPolicies, topic utils.TopicName, enabled bool) error {
			return policies.SetSchemaValidationEnforced(ctx, topic, enabled)
		},
	)
}

func RemoveSchemaValidationEnforcedCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-schema-validation-enforced", "Removed schema validation enforced for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveSchemaValidationEnforced(vc.Command.Context(), *topic)
	})
}
