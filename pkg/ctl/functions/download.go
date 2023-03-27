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
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsar-admin-go/pkg/admin/config"
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func downloadFunctionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for download File Data from Pulsar."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	download := cmdutils.Example{
		Desc: "Download File Data from Pulsar",
		Command: "pulsarctl functions download \n" +
			"\t--destination-file public\n" +
			"\t--path default\n",
	}
	downloadByNs := cmdutils.Example{
		Desc: "Download File Data from Pulsar",
		Command: "pulsarctl functions download \n" +
			"\t--destination-file public\n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name <function-name>\n",
	}
	examples = append(examples, download, downloadByNs)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Downloaded <the name of a Pulsar Function> successfully",
	}
	failOut := cmdutils.Output{
		Desc: "You must specify a name for the Pulsar Functions or a FQFN, please check the --name args",
		Out:  "[✖]  you must specify a name for the function or a Fully Qualified Function Name (FQFN)",
	}

	out = append(out, successOut, failOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"download",
		"Download File Data from Pulsar",
		desc.ToString(),
		desc.ExampleToString(),
		"download",
	)

	functionData := &utils.FunctionData{}

	// set the run function
	vc.SetRunFunc(func() error {
		return doDownloadFunctions(vc, functionData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("FunctionsConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&functionData.DestinationFile,
			"destination-file",
			"",
			"The file to store downloaded content")

		flagSet.StringVar(
			&functionData.Path,
			"path",
			"",
			"Path to store the content")

		flagSet.StringVar(
			&functionData.FQFN,
			"fqfn",
			"",
			"The Fully Qualified Function Name (FQFN) for the function")

		flagSet.StringVar(
			&functionData.Tenant,
			"tenant",
			"",
			"Tenant name")

		flagSet.StringVar(
			&functionData.Namespace,
			"namespace",
			"",
			"Namespace name")

		flagSet.StringVar(
			&functionData.FuncName,
			"name",
			"",
			"Function name")
	})
	vc.EnableOutputFlagSet()
}

func doDownloadFunctions(vc *cmdutils.VerbCmd, funcData *utils.FunctionData) error {
	if funcData.Path == "" {
		err := processBaseArguments(funcData)
		if err != nil {
			return err
		}
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)

	if funcData.Path != "" {
		err := admin.Functions().DownloadFunction(funcData.Path, funcData.DestinationFile)
		if err != nil {
			return err
		}
	} else {
		err := admin.Functions().DownloadFunctionByNs(funcData.DestinationFile, funcData.Tenant,
			funcData.Namespace, funcData.FuncName)
		if err != nil {
			return err
		}
	}

	vc.Command.Printf("Downloaded %s successfully\n", funcData.DestinationFile)
	return nil
}
