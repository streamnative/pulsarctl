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
	"errors"
	"strings"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetReplicationClustersCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-replication-clusters", "Get replication clusters for a topic", "Get replication clusters for a topic", "")
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
		clusters, err := policies.GetReplicationClusters(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, clusters, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetReplicationClustersCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var clusterIDs string
	vc.SetDescription("set-replication-clusters", "Set replication clusters for a topic", "Set replication clusters for a topic", "")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("ReplicationClusters", func(set *pflag.FlagSet) {
		set.StringVarP(&clusterIDs, "clusters", "c", "", "comma separated cluster names")
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
		err = policies.SetReplicationClusters(vc.Command.Context(), *topic, strings.Split(clusterIDs, ","))
		if err == nil {
			vc.Command.Printf("Set replication clusters successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveReplicationClustersCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-replication-clusters", "Removed replication clusters", func(global bool) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.RemoveReplicationClusters(vc.Command.Context(), *topic)
	})
}

func GetAutoSubscriptionCreationCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-auto-subscription-creation", "Get auto subscription creation policy", "Get auto subscription creation policy", "")
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
		override, err := policies.GetAutoSubscriptionCreation(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, override, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetAutoSubscriptionCreationCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var enable bool
	var disable bool
	vc.SetDescription("set-auto-subscription-creation", "Set auto subscription creation policy", "Set auto subscription creation policy", "")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("AutoSubscriptionCreation", func(set *pflag.FlagSet) {
		set.BoolVarP(&enable, "enable", "e", false, "enable auto subscription creation")
		set.BoolVarP(&disable, "disable", "d", false, "disable auto subscription creation")
	})
	vc.SetRunFuncWithNameArg(func() error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		if enable == disable {
			return errors.New("need to specify either --enable or --disable")
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		override := utils.AutoSubscriptionCreationOverride{AllowAutoSubscriptionCreation: enable}
		err = policies.SetAutoSubscriptionCreation(vc.Command.Context(), *topic, override)
		if err == nil {
			vc.Command.Printf("Set auto subscription creation successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveAutoSubscriptionCreationCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-auto-subscription-creation", "Removed auto subscription creation policy", func(global bool) error {
		topic, err := topicName(vc)
		if err != nil {
			return err
		}
		policies, err := topicPolicies(global)
		if err != nil {
			return err
		}
		return policies.RemoveAutoSubscriptionCreation(vc.Command.Context(), *topic)
	})
}
