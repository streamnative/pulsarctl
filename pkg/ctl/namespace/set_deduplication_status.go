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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/spf13/pflag"
)

func setDeduplication(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Enable or disable deduplication for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	enableDeduplication := cmdutils.Example{
		Desc:    "Enable or disable deduplication for a namespace",
		Command: "pulsarctl namespaces set-deduplication tenant/namespace (--enable)",
	}

	examples = append(examples, enableDeduplication)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set deduplication is [true or false] successfully for public/default",
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

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-deduplication",
		"Enable or disable deduplication for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"set-deduplication",
	)

	var data utils.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSetDeduplication(vc, data)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.BoolVarP(
			&data.Enable,
			"enable",
			"e",
			false,
			"Enable deduplication")
	})
}

func doSetDeduplication(vc *cmdutils.VerbCmd, data utils.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()

	err := admin.Namespaces().SetDeduplicationStatus(ns, data.Enable)
	if err == nil {
		vc.Command.Printf("Set deduplication is [%v] successfully for %s\n", data.Enable, ns)
	}

	return err
}
