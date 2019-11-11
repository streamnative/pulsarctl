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
)

func getPolicies(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the configuration policies of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	police := cmdutils.Example{
		Desc:    "Get the configuration policies of a namespace",
		Command: "pulsarctl namespaces policies (tenant/namespace)",
	}
	examples = append(examples, police)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"AuthPolicies\": {},\n" +
			"  \"ReplicationClusters\": null,\n" +
			"  \"Bundles\": {\n" +
			"    \"boundaries\": [\n" +
			"      \"0x00000000\",\n" +
			"      \"0x40000000\",\n" +
			"      \"0x80000000\",\n" +
			"      \"0xc0000000\",\n" +
			"      \"0xffffffff\"\n" +
			"    ],\n" +
			"    \"numBundles\": 4\n" +
			"  },\n" +
			"  \"BacklogQuotaMap\": null,\n" +
			"  \"TopicDispatchRate\": {\n" +
			"    \"standalone\": {\n" +
			"      \"DispatchThrottlingRateInMsg\": 0,\n" +
			"      \"DispatchThrottlingRateInByte\": 0,\n" +
			"      \"RatePeriodInSecond\": 1\n" +
			"    }\n" +
			"  },\n" +
			"  \"SubscriptionDispatchRate\": {\n" +
			"    \"standalone\": {\n" +
			"      \"DispatchThrottlingRateInMsg\": 0,\n" +
			"      \"DispatchThrottlingRateInByte\": 0,\n" +
			"      \"RatePeriodInSecond\": 1\n" +
			"    }\n" +
			"  },\n" +
			"  \"ClusterSubscribeRate\": {\n" +
			"    \"standalone\": {\n" +
			"      \"SubscribeThrottlingRatePerConsumer\": 0,\n" +
			"      \"RatePeriodInSecond\": 30\n" +
			"    }\n" +
			"  },\n" +
			"  \"Persistence\": {\n" +
			"    \"BookkeeperEnsemble\": 0,\n" +
			"    \"BookkeeperWriteQuorum\": 0,\n" +
			"    \"BookkeeperAckQuorum\": 0,\n" +
			"    \"ManagedLedgerMaxMarkDeleteRate\": 0\n" +
			"  },\n" +
			"  \"DeduplicationEnabled\": false,\n" +
			"  \"LatencyStatsSampleRate\": null,\n" +
			"  \"MessageTTLInSeconds\": 0,\n" +
			"  \"RetentionPolicies\": {\n" +
			"    \"RetentionTimeInMinutes\": 0,\n" +
			"    \"RetentionSizeInMB\": 0\n" +
			"  },\n" +
			"  \"Deleted\": false,\n" +
			"  \"AntiAffinityGroup\": \"\",\n" +
			"  \"EncryptionRequired\": false,\n" +
			"  \"SubscriptionAuthMode\": \"\",\n" +
			"  \"MaxProducersPerTopic\": 0,\n" +
			"  \"MaxConsumersPerTopic\": 0,\n" +
			"  \"MaxConsumersPerSubscription\": 0,\n" +
			"  \"CompactionThreshold\": 0,\n" +
			"  \"OffloadThreshold\": 0,\n" +
			"  \"OffloadDeletionLagMs\": 0,\n" +
			"  \"SchemaCompatibilityStrategy\": \"\",\n" +
			"  \"SchemaValidationEnforced\": false\n" +
			"}",
	}

	noNamespaceName := cmdutils.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}

	tenantNotExistError := cmdutils.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := cmdutils.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"policies",
		"Get the configuration policies of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"policies",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetPolicies(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetPolicies(vc *cmdutils.VerbCmd) error {
	namespace := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	policies, err := admin.Namespaces().GetPolicies(namespace)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), policies)
	}
	return err
}
