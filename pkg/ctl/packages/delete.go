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

package packages

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func deletePackagesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Delete a package."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc: "Delete a package",
		Command: "pulsarctl packages delete \n" +
			"\tfunction://public/default/test@v1",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The package 'function://public/default/test@v1' deleted successfully\n",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete a package",
		desc.ToString(),
		desc.ExampleToString(),
		"delete",
	)

	// set the run function
	vc.SetRunFuncWithNameArg(func() error {
		return doDeletePackage(vc)
	}, "the package URL is not provided")

	vc.EnableOutputFlagSet()
}

func doDeletePackage(vc *cmdutils.VerbCmd) error {
	_, err := utils.GetPackageName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(common.V3)
	err = admin.Packages().Delete(vc.NameArg)
	if err != nil {
		return err
	}

	vc.Command.Printf("The package '%s' deleted successfully\n", vc.NameArg)

	return err
}
