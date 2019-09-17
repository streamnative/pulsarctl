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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func setBacklogQuota(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Set a backlog quota policy for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	setBacklog := pulsar.Example{
		Desc: "Set a backlog quota policy for a namespace",
		Command: "pulsarctl namespaces set-backlog-quota tenant/namespace \n" +
			"\t--limit 2G \n" +
			"\t--policy producer_request_hold",
	}
	examples = append(examples, setBacklog)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Set backlog quota successfully for [tenant/namespace]",
	}

	notTenantName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}

	notExistTenantName := pulsar.Output{
		Desc: "the tenant name not exist, please check the tenant name",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	notExistNsName := pulsar.Output{
		Desc: "the namespace not exist, please check namespace name",
		Out:  "[✖]  code: 404 reason: Namespace <tenant/namespace> does not exist",
	}

	noSupportPolicyType := pulsar.Output{
		Desc: "invalid retention policy type, please check --policy arg",
		Out:  "invalid retention policy type: <policy type>",
	}

	out = append(out, successOut, notTenantName, notExistTenantName, notExistNsName, noSupportPolicyType)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-backlog-quota",
		"Set a backlog quota policy for a namespace",
		desc.ToString(),
		"set-backlog-quota",
	)

	var namespaceData pulsar.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSetBacklogQuota(vc, namespaceData)
	})

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&namespaceData.LimitStr,
			"limit",
			"l",
			"",
			"Size limit (eg: 10M, 16G)")

		flagSet.StringVarP(
			&namespaceData.PolicyStr,
			"policy",
			"p",
			"",
			"Retention policy to enforce when the limit is reached.\n"+
				"Valid options are: [producer_request_hold, producer_exception, consumer_backlog_eviction]")
		cobra.MarkFlagRequired(flagSet, "limit")
		cobra.MarkFlagRequired(flagSet, "policy")
	})
}

func doSetBacklogQuota(vc *cmdutils.VerbCmd, data pulsar.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()

	sizeLimit, err := validateSizeString(data.LimitStr)
	if err != nil {
		return err
	}

	var policy pulsar.RetentionPolicy
	switch data.PolicyStr {
	case "producer_request_hold":
		policy = pulsar.ProducerRequestHold
	case "producer_exception":
		policy = pulsar.ProducerException
	case "consumer_backlog_eviction":
		policy = pulsar.ConsumerBacklogEviction
	default:
		return fmt.Errorf("invalid retention policy type: %v", data.PolicyStr)
	}

	err = admin.Namespaces().SetBacklogQuota(ns, pulsar.NewBacklogQuota(sizeLimit, policy))
	if err == nil {
		vc.Command.Printf("Set backlog quota successfully for [%s]", ns)
	}
	return err
}
