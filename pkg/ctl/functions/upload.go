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

package functions

import (
	"fmt"
	"strings"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func uploadFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for uploading a local file to Pulsar."
	desc.CommandPermission = "This command requires super-user permission."

	var examples []cmdutils.Example
	uploadExample := cmdutils.Example{
		Desc:    "Upload a local file to Pulsar",
		Command: `pulsarctl functions upload --source-file <file-path> --path public/default/test`,
	}
	examples = append(examples, uploadExample)
	desc.CommandExamples = examples

	vc.SetDescription(
		"upload",
		"Upload a local file to Pulsar",
		desc.ToString(),
		desc.ExampleToString(),
		"upload")

	var sourceFile, path string
	vc.SetRunFunc(func() error {
		return doUploadFunction(vc, sourceFile, path)
	})

	vc.FlagSetGroup.InFlagSet("Upload", func(set *pflag.FlagSet) {
		set.StringVar(&sourceFile, "source-file", "",
			"The file whose content will be uploaded")
		_ = cobra.MarkFlagRequired(set, "source-file")
		set.StringVar(&path, "path", "", "Path where the contents will to be stored")
		_ = cobra.MarkFlagRequired(set, "path")
	})

}

func doUploadFunction(vc *cmdutils.VerbCmd, sourceFile, path string) error {
	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	if strings.TrimSpace(sourceFile) == "" || strings.TrimSpace(path) == "" {
		return fmt.Errorf("the source file or the path can not be specified as empty")
	}
	err := admin.Functions().Upload(sourceFile, path)
	if err != nil {
		return err
	}
	vc.Command.Printf("Upload file %s to Pulsar path %s successfully", sourceFile, path)
	return nil
}
