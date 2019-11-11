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

package brokers

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func getOwnedNamespacesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "List namespaces owned by the broker"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "List namespaces owned by the broker",
		Command: "pulsarctl brokers namespaces (cluster-name) --url (eg:127.0.0.1:8080)",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"  \"public/functions/0x40000000_0x80000000\": {\n" +
			"    \"broker_assignment\": \"shared\",\n" +
			"    \"is_controlled\": false,\n" +
			"    \"is_active\": true\n" +
			"  },\n" +
			"  \"pulsar/standalone/127.0.0.1:8080/0x00000000_0xffffffff\": {\n" +
			"    \"broker_assignment\": \"shared\",\n" +
			"    \"is_controlled\": false,\n" +
			"    \"is_active\": true\n" +
			"  }\n" +
			"}",
	}

	var argsError = cmdutils.Output{
		Desc: "the cluster name is not specified or the cluster name is specified more than one",
		Out:  "[✖]  the cluster name is not specified or the cluster name is specified more than one",
	}

	var urlError = cmdutils.Output{
		Desc: "The correct url is not provided, please check the `--url` arg.",
		Out:  "[✖]  Get (broker url)/admin/v2/brokers/standalone/127.0.0.1:6650/ownedNamespaces: EOF",
	}

	out = append(out, successOut, argsError, urlError)
	desc.CommandOutput = out

	vc.SetDescription(
		"namespaces",
		"List namespaces owned by the broker",
		desc.ToString(),
		desc.ExampleToString(),
		"namespaces")

	brokerData := &utils.BrokerData{}

	vc.SetRunFuncWithNameArg(func() error {
		return doOwnedNamespaces(vc, brokerData)
	}, "the cluster name is not specified or the cluster name is specified more than one")

	// register the params
	vc.FlagSetGroup.InFlagSet("BrokerData", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&brokerData.URL,
			"url",
			"",
			"broker-url")

		cobra.MarkFlagRequired(flagSet, "url")
	})
}

func doOwnedNamespaces(vc *cmdutils.VerbCmd, brokerData *utils.BrokerData) error {
	clusterName := vc.NameArg
	if clusterName == "" {
		return errors.New("should specified a cluster name")
	}

	admin := cmdutils.NewPulsarClient()
	namespaces, err := admin.Brokers().GetOwnedNamespaces(clusterName, brokerData.URL)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), namespaces)
	}
	return err
}
