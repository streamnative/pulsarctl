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
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"topic-policies",
		"Operations about topic policies",
		"",
		"topic-policy",
	)

	cmdutils.AddVerbCmds(flagGrouping, resourceCmd,
		GetMessageTTLCmd,
		SetMessageTTLCmd,
		RemoveMessageTTLCmd,
		GetMaxUnackMessagesPerConsumerCmd,
		SetMaxUnackMessagesPerConsumerCmd,
		RemoveMaxUnackMessagesPerConsumerCmd,
		GetMaxConsumersPerSubscriptionCmd,
		SetMaxConsumersPerSubscriptionCmd,
		RemoveMaxConsumersPerSubscriptionCmd,
		GetRetentionCmd,
		SetRetentionCmd,
		RemoveRetentionCmd,
		GetBacklogQuotaCmd,
		SetBacklogQuotaCmd,
		RemoveBacklogQuotaCmd,
		GetMaxProducersCmd,
		SetMaxProducersCmd,
		RemoveMaxProducersCmd,
		GetDeduplicationCmd,
		SetDeduplicationCmd,
		RemoveDeduplicationCmd,
		GetPersistenceCmd,
		SetPersistenceCmd,
		RemovePersistenceCmd,
		GetSubscriptionDispatchRateCmd,
		SetSubscriptionDispatchRateCmd,
		RemoveSubscriptionDispatchRateCmd,
		GetPublishRateCmd,
		SetPublishRateCmd,
		RemovePublishRateCmd,
		GetCompactionThresholdCmd,
		SetCompactionThresholdCmd,
		RemoveCompactionThresholdCmd,
		GetSubscribeRateCmd,
		SetSubscribeRateCmd,
		RemoveSubscribeRateCmd,
		GetMaxConsumersCmd,
		SetMaxConsumersCmd,
		RemoveMaxConsumersCmd,
		GetDelayedDeliveryCmd,
		SetDelayedDeliveryCmd,
		RemoveDelayedDeliveryCmd,
		GetDispatchRateCmd,
		SetDispatchRateCmd,
		RemoveDispatchRateCmd,
		GetMaxUnackMessagesPerSubscriptionCmd,
		SetMaxUnackMessagesPerSubscriptionCmd,
		RemoveMaxUnackMessagesPerSubscriptionCmd,
		GetInactiveTopicPoliciesCmd,
		SetInactiveTopicPoliciesCmd,
		RemoveInactiveTopicPoliciesCmd,
		GetMaxMessageSizeCmd,
		SetMaxMessageSizeCmd,
		RemoveMaxMessageSizeCmd,
		GetMaxSubscriptionsPerTopicCmd,
		SetMaxSubscriptionsPerTopicCmd,
		RemoveMaxSubscriptionsPerTopicCmd,
		GetDeduplicationSnapshotIntervalCmd,
		SetDeduplicationSnapshotIntervalCmd,
		RemoveDeduplicationSnapshotIntervalCmd,
		GetReplicatorDispatchRateCmd,
		SetReplicatorDispatchRateCmd,
		RemoveReplicatorDispatchRateCmd,
		GetOffloadPoliciesCmd,
		SetOffloadPoliciesCmd,
		RemoveOffloadPoliciesCmd,
		GetAutoSubscriptionCreationCmd,
		SetAutoSubscriptionCreationCmd,
		RemoveAutoSubscriptionCreationCmd,
		GetSchemaCompatibilityStrategyCmd,
		SetSchemaCompatibilityStrategyCmd,
		RemoveSchemaCompatibilityStrategyCmd,
		GetReplicationClustersCmd,
		SetReplicationClustersCmd,
		RemoveReplicationClustersCmd,
	)

	return resourceCmd
}
