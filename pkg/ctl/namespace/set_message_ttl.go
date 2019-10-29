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
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func setMessageTTL(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Set Message TTL for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	setMsgTTL := pulsar.Example{
		Desc:    "Set Message TTL for a namespace",
		Command: "pulsarctl namespaces set-message-ttl tenant/namespace -ttl 10",
	}
	examples = append(examples, setMsgTTL)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Set message TTL successfully for [tenant/namespace]",
	}

	noNamespaceName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}

	tenantNotExistError := pulsar.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := pulsar.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	failOut := pulsar.Output{
		Desc: "Invalid value for message TTL, please check -ttl arg",
		Out:  "code: 412 reason: Invalid value for message TTL",
	}
	out = append(out, successOut, failOut, noNamespaceName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-message-ttl",
		"Set Message TTL for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"set-message-ttl",
	)

	var namespaceData pulsar.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSetMessageTTL(vc, namespaceData)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.IntVarP(
			&namespaceData.MessageTTL,
			"messageTTL",
			"t",
			0,
			"Message TTL in seconds")
		cobra.MarkFlagRequired(flagSet, "messageTTL")
	})
}

func doSetMessageTTL(vc *cmdutils.VerbCmd, data pulsar.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	err := admin.Namespaces().SetNamespaceMessageTTL(ns, data.MessageTTL)
	if err == nil {
		vc.Command.Printf("Set message TTL successfully for [%s]\n", ns)
	}
	return err
}
