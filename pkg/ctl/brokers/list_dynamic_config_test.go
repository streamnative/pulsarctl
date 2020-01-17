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

package brokers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDynamicConfigListNameCmd(t *testing.T) {
	args := []string{"list-dynamic-config"}
	listOut, execErr, _, _ := TestBrokersCommands(getDynamicConfigListNameCmd, args)
	assert.Nil(t, execErr)
	expectedOut := `+-------------------------------------------------+\n|
DYNAMIC CONFIG NAMES               |
\n+-------------------------------------------------+\n|
dispatchThrottlingRatePerTopicInMsg             |\n|
brokerPublisherThrottlingTickTimeMillis         |\n|
loadBalancerSheddingEnabled                     |\n|
brokerClientAuthenticationParameters            |\n|
dispatchThrottlingRatePerReplicatorInByte       |\n|
loadBalancerBrokerMaxTopics                     |\n|
maxConcurrentTopicLoadRequest                   |\n|
brokerShutdownTimeoutMs                         |\n|
preferLaterVersions                             |\n|
subscribeThrottlingRatePerConsumer              |\n|
brokerClientAuthenticationPlugin                |\n|
dispatchThrottlingRatePerTopicInByte            |\n|
dispatcherMaxReadBatchSize                      |\n|
dispatcherMinReadBatchSize                      |\n|
loadBalancerReportUpdateThresholdPercentage     |\n|
dispatchThrottlingOnNonBacklogConsumerEnabled   |\n|
superUserRoles                                  |\n|
dispatchThrottlingRatePerReplicatorInMsg        |\n|
brokerPublisherThrottlingMaxByteRate            |\n|
loadManagerClassName                            |\n|
autoSkipNonRecoverableData                      |\n|
subscriptionKeySharedEnable                     |\n|
brokerPublisherThrottlingMaxMessageRate         |\n|
loadBalancerBrokerOverloadedThresholdPercentage |\n|
loadBalancerReportUpdateMaxIntervalMinutes      |\n|
dispatchThrottlingRatePerSubscriptionInByte     |\n|
dispatchThrottlingRateRelativeToPublishRate     |\n|
maxConcurrentLookupRequest                      |\n|
dispatcherMaxRoundRobinBatchSize                |\n|
subscriptionRedeliveryTrackerEnabled            |\n|
topicPublisherThrottlingTickTimeMillis          |\n|
failureDomainsEnabled                           |\n|
loadBalancerAutoBundleSplitEnabled              |\n|
brokerClientTlsEnabled                          |\n|
subscribeRatePeriodPerConsumerInSecond          |\n|
dispatchThrottlingRatePerSubscriptionInMsg      |\n|
clientLibraryVersionCheckEnabled                |\n|
loadBalancerAutoUnloadSplitBundlesEnabled       |\n
+-------------------------------------------------+\n`
	assert.Equal(t, expectedOut, listOut.String())
}
