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
	"github.com/streamnative/pulsar-admin-go/pkg/admin/config"
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getPackageMetadataCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get the metadata of a package"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc: "Get the metadata of a package\n",
		Command: "pulsarctl packages get-metadata \n" +
			"\tfunction://public/default/test@v1",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "{\n" +
			"   \"description\":\"test\",\n" +
			"   \"contact\":\"apache pulsar\",\n" +
			"   \"createTime\":1,\n" +
			"   \"modificationTime\":1,\n" +
			"   \"properties\":{\n" +
			"      \"foo\":\"bar\"\n" +
			"   }\n" +
			"}",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"get-metadata",
		"Get a package metadata information",
		desc.ToString(),
		desc.ExampleToString(),
		"get-metadata",
	)

	// set the run function
	vc.SetRunFuncWithNameArg(func() error {
		return doGetPackageMetadata(vc)
	}, "the package URL is not provided")

	vc.EnableOutputFlagSet()
}

func doGetPackageMetadata(vc *cmdutils.VerbCmd) error {
	_, err := utils.GetPackageName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	metadata, err := admin.Packages().GetMetadata(vc.NameArg)
	if err == nil {
		oc := cmdutils.NewOutputContent().WithObject(metadata)
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
	}

	return err
}
