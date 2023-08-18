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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func uploadPackagesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Upload a package"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc: "Upload a package",
		Command: "pulsarctl packages upload \n" +
			"\tfunction://public/default/test@v1 \n" +
			"\t--path /pulsar/examples/test.jar \n" +
			"\t--description test",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The package 'function://public/default/test@v1' uploaded from path '/pulsar/examples/test.jar' successfully\n",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"upload",
		"Upload a package",
		desc.ToString(),
		desc.ExampleToString(),
		"upload",
	)

	var path string
	packageMetadata := &utils.PackageMetadata{}

	// set the run function
	vc.SetRunFuncWithNameArg(func() error {
		return doUploadPackage(vc, &path, packageMetadata)
	}, "the package URL is not provided")

	vc.FlagSetGroup.InFlagSet("Upload Package", func(set *pflag.FlagSet) {
		set.StringVarP(
			&packageMetadata.Description,
			"description",
			"",
			"",
			"descriptions of a package")
		set.StringVarP(
			&packageMetadata.Contact,
			"contact",
			"",
			"",
			"contact info of a package")
		set.StringToStringVarP(
			&packageMetadata.Properties,
			"properties",
			"P",
			nil,
			"external information of a package")
		set.StringVarP(
			&path,
			"path",
			"",
			"",
			"file path of the package")
		cobra.MarkFlagRequired(set, "description")
		cobra.MarkFlagRequired(set, "path")
	})

	vc.EnableOutputFlagSet()
}

func doUploadPackage(vc *cmdutils.VerbCmd, path *string, packageMetadata *utils.PackageMetadata) error {
	_, err := utils.GetPackageName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	err = admin.Packages().Upload(vc.NameArg, *path, packageMetadata.Description,
		packageMetadata.Contact, packageMetadata.Properties)
	if err != nil {
		return err
	}

	vc.Command.Printf("The package '%s' uploaded from path '%s' successfully\n", vc.NameArg, *path)

	return err
}
