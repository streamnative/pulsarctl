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

func GetOffloadPoliciesCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-offload-policies", "Get offload policies", "Get offload policies", "")
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
		value, err := policies.GetOffloadPolicies(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, value, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetOffloadPoliciesCmd(vc *cmdutils.VerbCmd) {
	var global bool
	policy := utils.NewOffloadPolicies()
	vc.SetDescription("set-offload-policies", "Set offload policies", "Set offload policies", "")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("OffloadPolicies", func(set *pflag.FlagSet) {
		set.StringVarP(&policy.ManagedLedgerOffloadDriver, "driver", "", "", "offload driver")
		set.IntVarP(&policy.ManagedLedgerOffloadMaxThreads, "max-threads", "", 2, "max offload threads")
		set.Int64VarP(&policy.ManagedLedgerOffloadThresholdInBytes, "threshold-bytes", "", -1, "offload threshold in bytes")
		set.Int64VarP(&policy.ManagedLedgerOffloadDeletionLagInMillis, "deletion-lag-millis", "", 14400000, "offload deletion lag in millis")
		set.Int64VarP(&policy.ManagedLedgerOffloadAutoTriggerSizeThresholdBytes, "auto-trigger-size-threshold-bytes", "", -1, "auto trigger size threshold bytes")
		set.StringVarP(&policy.S3ManagedLedgerOffloadBucket, "bucket", "", "", "S3 bucket")
		set.StringVarP(&policy.S3ManagedLedgerOffloadRegion, "region", "", "", "S3 region")
		set.StringVarP(&policy.S3ManagedLedgerOffloadServiceEndpoint, "service-endpoint", "", "", "S3 service endpoint")
		set.StringVarP(&policy.S3ManagedLedgerOffloadCredentialID, "credential-id", "", "", "credential id")
		set.StringVarP(&policy.S3ManagedLedgerOffloadCredentialSecret, "credential-secret", "", "", "credential secret")
		set.StringVarP(&policy.S3ManagedLedgerOffloadRole, "role", "", "", "S3 role")
		set.StringVarP(&policy.S3ManagedLedgerOffloadRoleSessionName, "role-session-name", "", "", "S3 role session name")
		set.StringVarP(&policy.OffloadersDirectory, "offloaders-directory", "", "", "offloaders directory")
		set.StringToStringVarP(&policy.ManagedLedgerOffloadDriverMetadata, "driver-metadata", "", nil, "driver metadata in key=value,key2=value2 format")
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
		err = policies.SetOffloadPolicies(vc.Command.Context(), *topic, *policy)
		if err == nil {
			vc.Command.Printf("Set offload policies successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveOffloadPoliciesCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-offload-policies", "Removed offload policies", func(global bool) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.RemoveOffloadPolicies(vc.Command.Context(), *topic)
	})
}
