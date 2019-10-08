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
)

func ClearOffloadDeletionLagCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for clearing offload deletion lag of a namespace."
	desc.CommandPermission = "This command requires super-user permissions and broker has write policies permission."

	var examples []pulsar.Example
	clear := pulsar.Example{
		Desc:    "Clear offload deletion lag of the namespace (namespace-name)",
		Command: "pulsarctl namespaces clear-offload-deletion-lag (namespace-name)",
	}
	examples = append(examples, clear)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Successfully clear the offload deletion lag of the namespace (namespace-name)",
	}
	out = append(out, successOut, ArgError, NsNotExistError)
	out = append(out, NsErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"clear-offload-deletion-lag",
		"Clear offload deletion lag of a namespace",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doClearOffloadDeletionLag(vc)
	})
}

func doClearOffloadDeletionLag(vc *cmdutils.VerbCmd) error {
	ns, err := pulsar.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().ClearOffloadDeleteLag(*ns)
	if err == nil {
		vc.Command.Printf("Successfully clear the offload deletion lag of the namespace %s\n", ns.String())
	}

	return err
}
