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

func SetEncryptionRequiredCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for enabling or disabling messages encryption for a namespace."
	desc.CommandPermission = "This command requires tenant admin and " +
		"a broker needs the read-write operations of the global zookeeper."

	var examples []cmdutils.Example
	enable := cmdutils.Example{
		Desc:    "Enable messages encryption for the namespace (namespace-name)",
		Command: "pulsarctl namespaces messages-encryption (namespace-name)",
	}

	disable := cmdutils.Example{
		Desc:    "Disable messages encryption for the namespace (namespace-name)",
		Command: "pulsarct. namespaces messages-encryption --disable (namespace-name)",
	}
	examples = append(examples, enable, disable)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Enable/Disable message encryption for the namespace (namespace-name)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"messages-encryption",
		"Enable or disable messages encryption of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	var d bool

	vc.SetRunFuncWithNameArg(func() error {
		return doSetEncryptionRequired(vc, d)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Messages Encryption", func(set *pflag.FlagSet) {
		set.BoolVar(&d, "disable", false, "Disable messages encryption")
	})
}

func doSetEncryptionRequired(vc *cmdutils.VerbCmd, disable bool) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetEncryptionRequiredStatus(*ns, !disable)
	if err == nil {
		var out string
		if !disable {
			out = "Enable"
		} else {
			out = "Disable"
		}
		vc.Command.Printf("%s messages encryption of the namespace %s", out, ns.String())
	}

	return nil
}
