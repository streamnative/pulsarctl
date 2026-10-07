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

package topic

import (
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetMaxConsumersPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	vc.SetDescription("get-max-consumers-per-subscription", "Get max consumers per subscription for a topic", "Get max consumers per subscription for a topic", "", "get-max-consumers-per-subscription")
	vc.SetRunFuncWithNameArg(func() error {
		return doGetMaxConsumersPerSubscription(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doGetMaxConsumersPerSubscription(vc *cmdutils.VerbCmd) error {
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	value, err := admin.Topics().GetMaxConsumersPerSubscription(*topic)
	if err == nil {
		if value == -1 {
			vc.Command.Printf("The max consumers per subscription of the topic %s is not set\n", topic.String())
			return nil
		}
		vc.Command.Printf("%d\n", value)
	}
	return err
}

func SetMaxConsumersPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	var size int
	vc.SetDescription("set-max-consumers-per-subscription", "Set max consumers per subscription for a topic", "Set max consumers per subscription for a topic", "", "set-max-consumers-per-subscription")
	vc.FlagSetGroup.InFlagSet("MaxConsumersPerSubscription", func(set *pflag.FlagSet) {
		set.IntVar(&size, "size", -1, "max consumers per subscription")
		_ = cobra.MarkFlagRequired(set, "size")
	})
	vc.SetRunFuncWithNameArg(func() error {
		return doSetMaxConsumersPerSubscription(vc, size)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doSetMaxConsumersPerSubscription(vc *cmdutils.VerbCmd, size int) error {
	if vc.NameError != nil {
		return vc.NameError
	}
	if size < 0 {
		return errors.New("the specified consumers value must bigger than 0")
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().SetMaxConsumersPerSubscription(*topic, size)
	if err == nil {
		vc.Command.Printf("Set max consumers per subscription successfully for [%s]\n", topic.String())
	}
	return err
}

func RemoveMaxConsumersPerSubscriptionCmd(vc *cmdutils.VerbCmd) {
	vc.SetDescription("remove-max-consumers-per-subscription", "Remove max consumers per subscription for a topic", "Remove max consumers per subscription for a topic", "", "remove-max-consumers-per-subscription")
	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveMaxConsumersPerSubscription(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doRemoveMaxConsumersPerSubscription(vc *cmdutils.VerbCmd) error {
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().RemoveMaxConsumersPerSubscription(*topic)
	if err == nil {
		vc.Command.Printf("Removed max consumers per subscription successfully for [%s]\n", topic.String())
	}
	return err
}
