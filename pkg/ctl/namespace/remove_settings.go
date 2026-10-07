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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func newNamespaceRemoveCmd(use, short string, run func(*cmdutils.VerbCmd) error) func(*cmdutils.VerbCmd) {
	return func(vc *cmdutils.VerbCmd) {
		var desc cmdutils.LongDescription
		desc.CommandUsedFor = short
		desc.CommandPermission = "This command requires tenant admin permissions."
		desc.CommandOutput = append(desc.CommandOutput, ArgError)
		desc.CommandOutput = append(desc.CommandOutput, NsErrors...)
		desc.CommandOutput = append(desc.CommandOutput, NsNotExistError)

		vc.SetDescription(use, short, desc.ToString(), desc.ExampleToString())
		vc.SetRunFuncWithNameArg(func() error {
			return run(vc)
		}, "the namespace name is not specified or the namespace name is specified more than one")
	}
}

var removeMessageTTL = newNamespaceRemoveCmd(
	"remove-message-ttl",
	"Remove message TTL for a namespace",
	func(vc *cmdutils.VerbCmd) error {
		admin := cmdutils.NewPulsarClient()
		err := admin.Namespaces().RemoveNamespaceMessageTTL(vc.NameArg)
		if err == nil {
			vc.Command.Printf("Removed message TTL successfully for [%s]\n", vc.NameArg)
		}
		return err
	},
)

var removeRetention = newNamespaceRemoveCmd(
	"remove-retention",
	"Remove retention for a namespace",
	func(vc *cmdutils.VerbCmd) error {
		admin := cmdutils.NewPulsarClient()
		err := admin.Namespaces().RemoveRetention(vc.NameArg)
		if err == nil {
			vc.Command.Printf("Removed retention successfully for [%s]\n", vc.NameArg)
		}
		return err
	},
)

var removePersistence = newNamespaceRemoveCmd(
	"remove-persistence",
	"Remove persistence for a namespace",
	func(vc *cmdutils.VerbCmd) error {
		admin := cmdutils.NewPulsarClient()
		err := admin.Namespaces().RemovePersistence(vc.NameArg)
		if err == nil {
			vc.Command.Printf("Removed persistence successfully for [%s]\n", vc.NameArg)
		}
		return err
	},
)

func removeMaxConsumersPerSubscription(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().RemoveMaxConsumersPerSubscription(*ns)
	if err == nil {
		vc.Command.Printf("Removed max consumers per subscription successfully for [%s]\n", ns.String())
	}
	return err
}

var RemoveMaxConsumersPerSubscriptionCmd = newNamespaceRemoveCmd(
	"remove-max-consumers-per-subscription",
	"Remove the max consumers per subscription of a namespace",
	removeMaxConsumersPerSubscription,
)

func removeMaxConsumersPerTopic(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().RemoveMaxConsumersPerTopic(*ns)
	if err == nil {
		vc.Command.Printf("Removed max consumers per topic successfully for [%s]\n", ns.String())
	}
	return err
}

var RemoveMaxConsumersPerTopicCmd = newNamespaceRemoveCmd(
	"remove-max-consumers-per-topic",
	"Remove the max consumers per topic of a namespace",
	removeMaxConsumersPerTopic,
)

func removeMaxProducersPerTopic(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}
	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().RemoveMaxProducersPerTopic(*ns)
	if err == nil {
		vc.Command.Printf("Removed max producers per topic successfully for [%s]\n", ns.String())
	}
	return err
}

var RemoveMaxProducersPerTopicCmd = newNamespaceRemoveCmd(
	"remove-max-producers-per-topic",
	"Remove the max producers per topic of a namespace",
	removeMaxProducersPerTopic,
)
