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

	"github.com/spf13/pflag"
)

func unload(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Unload a namespace from the current serving broker"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	unload := pulsar.Example{
		Desc:    "Unload a namespace from the current serving broker",
		Command: "pulsarctl namespaces unload tenant/namespace",
	}

	unloadWithBundle := pulsar.Example{
		Desc:    "Unload a namespace with bundle from the current serving broker",
		Command: "pulsarctl namespaces unload tenant/namespace --bundle ({start-boundary}_{end-boundary})",
	}
	examples = append(examples, unload, unloadWithBundle)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Unload namespace (tenant/namespace) (with bundle ({start-boundary}_{end-boundary})) successfully ",
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

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"unload",
		"Unload a namespace from the current serving broker",
		desc.ToString(),
		desc.ExampleToString(),
		"unload",
	)

	var data pulsar.NamespacesData
	vc.SetRunFuncWithNameArg(func() error {
		return doUnload(vc, data)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&data.Bundle,
			"bundle",
			"b",
			"",
			"{start-boundary}_{end-boundary}(e.g. 0x00000000_0xffffffff)")
	})
}

func doUnload(vc *cmdutils.VerbCmd, data pulsar.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	if data.Bundle == "" {
		err := admin.Namespaces().Unload(ns)
		if err == nil {
			vc.Command.Printf("Unload namespace %s successfully\n", ns)
		}
		return err
	}

	err := admin.Namespaces().UnloadNamespaceBundle(ns, data.Bundle)
	if err == nil {
		vc.Command.Printf("Unload namespace %s with bundle %s successfully\n", ns, data.Bundle)
	}
	return err
}
