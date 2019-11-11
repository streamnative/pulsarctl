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
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"

	util "github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func setBacklogQuota(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set a backlog quota policy for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	setBacklog := cmdutils.Example{
		Desc: "Set a backlog quota policy for a namespace",
		Command: "pulsarctl namespaces set-backlog-quota tenant/namespace \n" +
			"\t--limit 2G \n" +
			"\t--policy producer_request_hold",
	}
	examples = append(examples, setBacklog)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set backlog quota successfully for [tenant/namespace]",
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

	noSupportPolicyType := cmdutils.Output{
		Desc: "invalid retention policy type, please check --policy arg",
		Out:  "invalid retention policy type: (policy type)",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError, noSupportPolicyType)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-backlog-quota",
		"Set a backlog quota policy for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"set-backlog-quota",
	)

	var namespaceData util.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSetBacklogQuota(vc, namespaceData)
	}, "the namespace name is not specified or the namespace name is specified more than one")

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

func doSetBacklogQuota(vc *cmdutils.VerbCmd, data util.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()

	sizeLimit, err := utils.ValidateSizeString(data.LimitStr)
	if err != nil {
		return err
	}

	var policy util.RetentionPolicy
	switch data.PolicyStr {
	case "producer_request_hold":
		policy = util.ProducerRequestHold
	case "producer_exception":
		policy = util.ProducerException
	case "consumer_backlog_eviction":
		policy = util.ConsumerBacklogEviction
	default:
		return fmt.Errorf("invalid retention policy type: %v", data.PolicyStr)
	}

	err = admin.Namespaces().SetBacklogQuota(ns, util.NewBacklogQuota(sizeLimit, policy))
	if err == nil {
		vc.Command.Printf("Set backlog quota successfully for [%s]\n", ns)
	}
	return err
}
