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

	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	ctlutils "github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

type backlogQuotaArgs struct {
	limitSize string
	limitTime int64
	policy    string
	quotaType string
}

type inactiveTopicPoliciesArgs struct {
	enableDeleteWhileInactive  bool
	disableDeleteWhileInactive bool
	maxInactiveDuration        string
	deleteMode                 string
}

func GetRetentionCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-retention", "Get retention policy for a topic", "Get retention policy for a topic", "", "get-retention")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		retention, err := policies.GetRetention(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, retention, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetRetentionCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var timeStr string
	var sizeStr string
	vc.SetDescription("set-retention", "Set retention policy for a topic", "Set retention policy for a topic", "", "set-retention")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("Retention", func(set *pflag.FlagSet) {
		set.StringVarP(&timeStr, "time", "t", "", "retention time with optional time unit suffix")
		set.StringVarP(&sizeStr, "size", "s", "", "retention size limit")
		_ = cobra.MarkFlagRequired(set, "time")
		_ = cobra.MarkFlagRequired(set, "size")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		retentionTimeInSeconds, err := ctlutils.ParseRelativeTimeInSeconds(timeStr)
		if err != nil {
			return err
		}
		sizeLimit, err := ctlutils.ValidateSizeString(sizeStr)
		if err != nil {
			return err
		}
		retentionTimeInMin := -1
		if retentionTimeInSeconds != -1 {
			retentionTimeInMin = int(retentionTimeInSeconds.Minutes())
		}
		retentionSizeInMB := -1
		if sizeLimit != -1 {
			retentionSizeInMB = int(sizeLimit / (1024 * 1024))
		}
		err = policies.SetRetention(vc.Command.Context(), *topic, util.NewRetentionPolicies(retentionTimeInMin, retentionSizeInMB))
		if err == nil {
			vc.Command.Printf("Set retention policy successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveRetentionCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-retention", "Removed retention policy for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveRetention(vc.Command.Context(), *topic)
	})
}

func GetBacklogQuotaCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-backlog-quota", "Get backlog quota for a topic", "Get backlog quota for a topic", "", "get-backlog-quota")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		quota, err := policies.GetBacklogQuotaMap(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, quota, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetBacklogQuotaCmd(vc *cmdutils.VerbCmd) {
	var global bool
	args := backlogQuotaArgs{}
	vc.SetDescription("set-backlog-quota", "Set backlog quota for a topic", "Set backlog quota for a topic", "", "set-backlog-quota")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("BacklogQuota", func(set *pflag.FlagSet) {
		set.StringVarP(&args.limitSize, "limit-size", "", "", "size limit")
		set.Int64VarP(&args.limitTime, "limit-time", "", -1, "time limit in seconds")
		set.StringVarP(&args.policy, "policy", "p", "", "retention policy")
		set.StringVarP(&args.quotaType, "type", "t", string(util.DestinationStorage), "backlog quota type")
		_ = cobra.MarkFlagRequired(set, "policy")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		sizeLimit := int64(-1)
		if args.limitSize != "" {
			sizeLimit, err = ctlutils.ValidateSizeString(args.limitSize)
			if err != nil {
				return err
			}
		}
		policy, err := util.ParseRetentionPolicy(args.policy)
		if err != nil {
			return err
		}
		quotaType, err := util.ParseBacklogQuotaType(args.quotaType)
		if err != nil {
			return err
		}
		err = policies.SetBacklogQuota(
			vc.Command.Context(),
			*topic,
			util.NewBacklogQuota(sizeLimit, args.limitTime, policy),
			quotaType,
		)
		if err == nil {
			vc.Command.Printf("Set backlog quota successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveBacklogQuotaCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var quotaType string
	vc.SetDescription("remove-backlog-quota", "Remove backlog quota for a topic", "Remove backlog quota for a topic", "", "remove-backlog-quota")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("BacklogQuota", func(set *pflag.FlagSet) {
		set.StringVarP(&quotaType, "type", "t", string(util.DestinationStorage), "backlog quota type")
	})
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		parsedType, err := util.ParseBacklogQuotaType(quotaType)
		if err != nil {
			return err
		}
		err = policies.RemoveBacklogQuota(vc.Command.Context(), *topic, parsedType)
		if err == nil {
			vc.Command.Printf("Removed backlog quota successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func GetPersistenceCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-persistence", "Get persistence policy for a topic", "Get persistence policy for a topic", "", "get-persistence")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		persistence, err := policies.GetPersistence(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, persistence, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetPersistenceCmd(vc *cmdutils.VerbCmd) {
	var global bool
	data := util.PersistenceData{}
	vc.SetDescription("set-persistence", "Set persistence policy for a topic", "Set persistence policy for a topic", "", "set-persistence")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("Persistence", func(set *pflag.FlagSet) {
		set.Int64VarP(&data.BookkeeperEnsemble, "bookkeeper-ensemble", "e", 0, "number of bookies")
		set.Int64VarP(&data.BookkeeperWriteQuorum, "bookkeeper-write-quorum", "w", 0, "bookkeeper write quorum")
		set.Int64VarP(&data.BookkeeperAckQuorum, "bookkeeper-ack-quorum", "a", 0, "bookkeeper ack quorum")
		set.Float64VarP(&data.ManagedLedgerMaxMarkDeleteRate, "ml-mark-delete-max-rate", "r", 0.0, "managed ledger max mark delete rate")
	})
	vc.SetRunFuncWithNameArg(func() error {
		if data.BookkeeperEnsemble <= 0 || data.BookkeeperWriteQuorum <= 0 || data.BookkeeperAckQuorum <= 0 {
			return errors.New("[--bookkeeper-ensemble], [--bookkeeper-write-quorum] and [--bookkeeper-ack-quorum] must greater than 0")
		}
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = policies.SetPersistence(vc.Command.Context(), *topic, data)
		if err == nil {
			vc.Command.Printf("Set persistence policy successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemovePersistenceCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-persistence", "Removed persistence policy for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemovePersistence(vc.Command.Context(), *topic)
	})
}

func GetDelayedDeliveryCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-delayed-delivery", "Get delayed delivery policy for a topic", "Get delayed delivery policy for a topic", "", "get-delayed-delivery")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		delayed, err := policies.GetDelayedDelivery(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, delayed, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetDelayedDeliveryCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var enable bool
	var disable bool
	var tickTime string
	var maxDelay string
	vc.SetDescription("set-delayed-delivery", "Set delayed delivery policy for a topic", "Set delayed delivery policy for a topic", "", "set-delayed-delivery")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("DelayedDelivery", func(set *pflag.FlagSet) {
		set.BoolVarP(&enable, "enable", "e", false, "enable delayed delivery messages")
		set.BoolVarP(&disable, "disable", "d", false, "disable delayed delivery messages")
		set.StringVarP(&tickTime, "time", "t", "1s", "tick time for delayed delivery")
		set.StringVarP(&maxDelay, "max-delay", "", "0s", "max allowed delay for delayed delivery")
	})
	vc.SetRunFuncWithNameArg(func() error {
		if enable == disable {
			return errors.New("need to specify either --enable or --disable")
		}
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		tickTimeInSeconds, err := ctlutils.ParseRelativeTimeInSeconds(tickTime)
		if err != nil {
			return err
		}
		maxDelayInSeconds, err := ctlutils.ParseRelativeTimeInSeconds(maxDelay)
		if err != nil {
			return err
		}
		data := util.NewDelayedDeliveryDataWithMaxDelay(
			tickTimeInSeconds.Seconds()*1000,
			enable,
			int64(maxDelayInSeconds.Seconds()*1000),
		)
		err = policies.SetDelayedDelivery(vc.Command.Context(), *topic, *data)
		if err == nil {
			vc.Command.Printf("Set delayed delivery policy successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveDelayedDeliveryCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-delayed-delivery", "Removed delayed delivery policy for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveDelayedDelivery(vc.Command.Context(), *topic)
	})
}

func GetInactiveTopicPoliciesCmd(vc *cmdutils.VerbCmd) {
	var global bool
	var applied bool
	vc.SetDescription("get-inactive-topic-policies", "Get inactive topic policies for a topic", "Get inactive topic policies for a topic", "", "get-inactive-topic-policies")
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		inactive, err := policies.GetInactiveTopicPolicies(vc.Command.Context(), *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, inactive, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func SetInactiveTopicPoliciesCmd(vc *cmdutils.VerbCmd) {
	var global bool
	args := inactiveTopicPoliciesArgs{}
	vc.SetDescription("set-inactive-topic-policies", "Set inactive topic policies for a topic", "Set inactive topic policies for a topic", "", "set-inactive-topic-policies")
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("InactiveTopicPolicies", func(set *pflag.FlagSet) {
		set.BoolVarP(&args.enableDeleteWhileInactive, "enable-delete-while-inactive", "e", false, "enable delete while inactive")
		set.BoolVarP(&args.disableDeleteWhileInactive, "disable-delete-while-inactive", "d", false, "disable delete while inactive")
		set.StringVarP(&args.maxInactiveDuration, "max-inactive-duration", "t", "", "max inactive duration")
		set.StringVarP(&args.deleteMode, "delete-mode", "m", "", "delete mode")
		_ = cobra.MarkFlagRequired(set, "max-inactive-duration")
		_ = cobra.MarkFlagRequired(set, "delete-mode")
	})
	vc.SetRunFuncWithNameArg(func() error {
		if args.enableDeleteWhileInactive == args.disableDeleteWhileInactive {
			return errors.New("need to specify either --enable-delete-while-inactive or --disable-delete-while-inactive")
		}
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		mode, err := util.ParseInactiveTopicDeleteMode(args.deleteMode)
		if err != nil {
			return err
		}
		maxInactiveDuration, err := ctlutils.ParseRelativeTimeInSeconds(args.maxInactiveDuration)
		if err != nil {
			return err
		}
		body := util.NewInactiveTopicPolicies(&mode, int(maxInactiveDuration.Seconds()), args.enableDeleteWhileInactive)
		err = policies.SetInactiveTopicPolicies(vc.Command.Context(), *topic, body)
		if err == nil {
			vc.Command.Printf("Set inactive topic policies successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func RemoveInactiveTopicPoliciesCmd(vc *cmdutils.VerbCmd) {
	removePolicyCmd(vc, "remove-inactive-topic-policies", "Removed inactive topic policies for a topic", func(global bool) error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		return policies.RemoveInactiveTopicPolicies(vc.Command.Context(), *topic)
	})
}
