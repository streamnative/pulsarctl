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

func downloadPackagesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Download a package"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example

	list := cmdutils.Example{
		Desc: "Download a package",
		Command: "pulsarctl packages download \n" +
			"\tfunction://public/default/test@v1 \n" +
			"\t--path /pulsar/examples/test.jar",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "The package 'function://public/default/test@v1' downloaded to path '/pulsar/examples/test.jar' successfully\n",
	}

	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"download",
		"Download a package",
		desc.ToString(),
		desc.ExampleToString(),
		"download",
	)

	var path string

	// set the run function
	vc.SetRunFuncWithNameArg(func() error {
		return doDownloadPackage(vc, &path)
	}, "the package URL is not provided")

	// register the params
	vc.FlagSetGroup.InFlagSet("Download Package", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&path,
			"path",
			"",
			"",
			"download destination path of the package")
		_ = cobra.MarkFlagRequired(flagSet, "path")
	})
	vc.EnableOutputFlagSet()
}

func doDownloadPackage(vc *cmdutils.VerbCmd, path *string) error {
	var _, err = utils.GetPackageName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	err = admin.Packages().Download(vc.NameArg, *path)
	if err != nil {
		return err
	}

	vc.Command.Printf("The package '%s' downloaded to path '%s' successfully\n", vc.NameArg, *path)

	return err
}
