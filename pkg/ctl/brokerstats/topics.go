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

package brokerstats

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func dumpTopics(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Dump the topics stats"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Dump the topics stats",
		Command: "pulsarctl broker-stats topics",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"    \"public/functions\": {\n" +
			"        \"0x40000000_0x80000000\": {\n" +
			"            \"persistent\": {\n" +
			"                \"persistent://public/functions/metadata\": {\n" +
			"                    \"publishers\": [\n" +
			"                        {\n" +
			"                            \"msgRateIn\": 0.0,\n" +
			"                            \"msgThroughputIn\": 0.0,\n" +
			"                            \"averageMsgSize\": 0.0,\n" +
			"                            \"address\": \"/127.0.0.1:49751\",\n" +
			"                            \"producerId\": 1,\n" +
			"                            \"producerName\": \"standalone-0-1\",\n" +
			"                            \"connectedSince\": \"2019-10-14T13:51:36.399+08:00\",\n" +
			"                            \"clientVersion\": \"2.4.1\",\n" +
			"                            \"metadata\": {}\n" +
			"                        }\n" +
			"                    ],\n" +
			"                    \"replication\": {},\n" +
			"                    \"subscriptions\": {\n" +
			"                        \"reader-3283723b08\": {\n" +
			"                            \"consumers\": [\n" +
			"                                {\n" +
			"                                    \"address\": \"/127.0.0.1:49751\",\n" +
			"                                    \"consumerName\": \"4e268\",\n" +
			"                                    \"availablePermits\": 1000,\n" +
			"                                    \"connectedSince\": \"2019-10-14T13:51:36.928+08:00\",\n" +
			"                                    \"msgRateOut\": 0.0,\n" +
			"                                    \"msgThroughputOut\": 0.0,\n" +
			"                                    \"msgRateRedeliver\": 0.0,\n" +
			"                                    \"clientVersion\": \"2.4.1\",\n" +
			"                                    \"metadata\": {}\n" +
			"                                }\n" +
			"                            ],\n" +
			"                            \"msgBacklog\": 0,\n" +
			"                            \"msgRateExpired\": 0.0,\n" +
			"                            \"msgRateOut\": 0.0,\n" +
			"                            \"msgThroughputOut\": 0.0,\n" +
			"                            \"msgRateRedeliver\": 0.0,\n" +
			"                            \"numberOfEntriesSinceFirstNotAckedMessage\": 1,\n" +
			"                            \"totalNonContiguousDeletedMessagesRange\": 0,\n" +
			"                            \"type\": \"Exclusive\"\n" +
			"                        }\n" +
			"                    },\n" +
			"                    \"producerCount\": 1,\n" +
			"                    \"averageMsgSize\": 0.0,\n" +
			"                    \"msgRateIn\": 0.0,\n" +
			"                    \"msgRateOut\": 0.0,\n" +
			"                    \"msgThroughputIn\": 0.0,\n" +
			"                    \"msgThroughputOut\": 0.0,\n" +
			"                    \"storageSize\": 0,\n" +
			"                    \"pendingAddEntriesCount\": 0\n" +
			"                },\n" +
			"                \"persistent://public/functions/coordinate\": {\n" +
			"                    \"publishers\": [],\n" +
			"                    \"replication\": {},\n" +
			"                    \"subscriptions\": {\n" +
			"                        \"participants\": {\n" +
			"                            \"consumers\": [\n" +
			"                                {\n" +
			"                                    \"address\": \"/127.0.0.1:49751\",\n" +
			"                                    \"consumerName\": \"89176\",\n" +
			"                                    \"availablePermits\": 1000,\n" +
			"                                    \"connectedSince\": \"2019-10-14T13:51:36.54+08:00\",\n" +
			"                                    \"msgRateOut\": 0.0,\n" +
			"                                    \"msgThroughputOut\": 0.0,\n" +
			"                                    \"msgRateRedeliver\": 0.0,\n" +
			"                                    \"clientVersion\": \"2.4.1\",\n" +
			"                                    \"metadata\": {\n" +
			"                                        \"id\": \"c-standalone-fw-127.0.0.1-8080:127.0.0.1:8080\"\n" +
			"                                    }\n" +
			"                                }\n" +
			"                            ],\n" +
			"                            \"msgBacklog\": 0,\n" +
			"                            \"msgRateExpired\": 0.0,\n" +
			"                            \"msgRateOut\": 0.0,\n" +
			"                            \"msgThroughputOut\": 0.0,\n" +
			"                            \"msgRateRedeliver\": 0.0,\n" +
			"                            \"numberOfEntriesSinceFirstNotAckedMessage\": 1,\n" +
			"                            \"totalNonContiguousDeletedMessagesRange\": 0,\n" +
			"                            \"type\": \"Failover\"\n" +
			"                        }\n" +
			"                    },\n" +
			"                    \"producerCount\": 0,\n" +
			"                    \"averageMsgSize\": 0.0,\n" +
			"                    \"msgRateIn\": 0.0,\n" +
			"                    \"msgRateOut\": 0.0,\n" +
			"                    \"msgThroughputIn\": 0.0,\n" +
			"                    \"msgThroughputOut\": 0.0,\n" +
			"                    \"storageSize\": 0,\n" +
			"                    \"pendingAddEntriesCount\": 0\n" +
			"                },\n" +
			"                \"persistent://public/functions/assignments\": {\n" +
			"                    \"publishers\": [\n" +
			"                        {\n" +
			"                            \"msgRateIn\": 0.0,\n" +
			"                            \"msgThroughputIn\": 0.0,\n" +
			"                            \"averageMsgSize\": 0.0,\n" +
			"                            \"address\": \"/127.0.0.1:49751\",\n" +
			"                            \"producerId\": 0,\n" +
			"                            \"producerName\": \"standalone-0-0\",\n" +
			"                            \"connectedSince\": \"2019-10-14T13:51:36.368+08:00\",\n" +
			"                            \"clientVersion\": \"2.4.1\",\n" +
			"                            \"metadata\": {}\n" +
			"                        }\n" +
			"                    ],\n" +
			"                    \"replication\": {},\n" +
			"                    \"subscriptions\": {\n" +
			"                        \"reader-b310b4910b\": {\n" +
			"                            \"consumers\": [\n" +
			"                                {\n" +
			"                                    \"address\": \"/127.0.0.1:49751\",\n" +
			"                                    \"consumerName\": \"25d1a\",\n" +
			"                                    \"availablePermits\": 1000,\n" +
			"                                    \"connectedSince\": \"2019-10-14T13:51:36.949+08:00\",\n" +
			"                                    \"msgRateOut\": 0.0,\n" +
			"                                    \"msgThroughputOut\": 0.0,\n" +
			"                                    \"msgRateRedeliver\": 0.0,\n" +
			"                                    \"clientVersion\": \"2.4.1\",\n" +
			"                                    \"metadata\": {}\n" +
			"                                }\n" +
			"                            ],\n" +
			"                            \"msgBacklog\": 0,\n" +
			"                            \"msgRateExpired\": 0.0,\n" +
			"                            \"msgRateOut\": 0.0,\n" +
			"                            \"msgThroughputOut\": 0.0,\n" +
			"                            \"msgRateRedeliver\": 0.0,\n" +
			"                            \"numberOfEntriesSinceFirstNotAckedMessage\": 1,\n" +
			"                            \"totalNonContiguousDeletedMessagesRange\": 0,\n" +
			"                            \"type\": \"Exclusive\"\n" +
			"                        }\n" +
			"                    },\n" +
			"                    \"producerCount\": 1,\n" +
			"                    \"averageMsgSize\": 0.0,\n" +
			"                    \"msgRateIn\": 0.0,\n" +
			"                    \"msgRateOut\": 0.0,\n" +
			"                    \"msgThroughputIn\": 0.0,\n" +
			"                    \"msgThroughputOut\": 0.0,\n" +
			"                    \"storageSize\": 0,\n" +
			"                    \"pendingAddEntriesCount\": 0\n" +
			"                }\n" +
			"            }\n" +
			"        }\n" +
			"    }\n" +
			"}",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"topics",
		"Dump the topics stats",
		desc.ToString(),
		desc.ExampleToString(),
		"topics")

	vc.SetRunFunc(func() error {
		return doDumpTopics(vc)
	})
}

func doDumpTopics(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	topicsStats, err := admin.BrokerStats().GetTopics()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), topicsStats)
	}
	return err
}
