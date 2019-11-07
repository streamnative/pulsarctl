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

package cluster

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func createFailureDomainCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for creating a failure domain of the (cluster-name)."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	create := cmdutils.Example{
		Desc:    "create the failure domain",
		Command: "pulsarctl clusters create-failure-domain (cluster-name) (domain-name)",
	}
	examples = append(examples, create)

	createWithBrokers := cmdutils.Example{
		Desc: "create the failure domain with brokers",
		Command: "pulsarctl clusters create-failure-domain" +
			" -b (broker-ip):(broker-port) -b (broker-ip):(broker-port) (cluster-name) (domain-name)",
	}
	examples = append(examples, createWithBrokers)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Create failure domain (domain-name) for cluster (cluster-name) succeed",
	}
	out = append(out, successOut)

	argsErrorOut := cmdutils.Output{
		Desc: "the args need to be specified as (cluster-name) (domain-name)",
		Out:  "[✖]  need specified two names for cluster and failure domain",
	}
	out = append(out, argsErrorOut)
	out = append(out, clusterNonExist)
	desc.CommandOutput = out

	vc.SetDescription(
		"create-failure-domain",
		"Create a failure domain",
		desc.ToString(),
		desc.ExampleToString(),
		"cfd")

	var failureDomainData utils.FailureDomainData

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doCreateFailureDomain(vc, &failureDomainData)
	}, checkFailureDomainArgs)

	vc.FlagSetGroup.InFlagSet("FailureDomainData", func(set *pflag.FlagSet) {
		set.StringSliceVarP(
			&failureDomainData.BrokerList,
			"brokers",
			"b",
			nil,
			"Set the failure domain clusters")
	})
}

func doCreateFailureDomain(vc *cmdutils.VerbCmd, failureDomain *utils.FailureDomainData) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	failureDomain.ClusterName = vc.NameArgs[0]
	failureDomain.DomainName = vc.NameArgs[1]

	if len(failureDomain.BrokerList) == 0 || failureDomain.BrokerList == nil {
		return errors.New("broker list must be specified")
	}

	admin := cmdutils.NewPulsarClient()
	err := admin.Clusters().CreateFailureDomain(*failureDomain)
	if err == nil {
		vc.Command.Printf(
			"Create failure domain [%s] for cluster [%s] succeed\n",
			failureDomain.DomainName, failureDomain.ClusterName)
	}
	return err
}
