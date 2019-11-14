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

package bookie

import (
	"github.com/streamnative/pulsarctl/pkg/bookkeeper/bkdata"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func listDiskFileCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for getting all the files on the disk of the current bookie."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "Get all the specified fileType (e.g. journal, entrylog, index) files on the disk of the current bookie",
		Command: "pulsarctl bookkeeper bookie listdiskfile (file-type)",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: `{
    "journal files" : "filename1 filename2 ...",
    "entrylog files" : "filename1 filename2...",
    "index files" : "filename1 filename2 ..."
}`,
	}

	argError := cmdutils.Output{
		Desc: "the file type is not specified or the file type is specified more than one",
		Out:  "[✖]  the file type is not specified or the file type is specified more than one",
	}

	typeError := cmdutils.Output{
		Desc: "the specified file type is invalid",
		Out: "[✖]  invalid file type %s, the file type only can be specified as 'journal', " +
			"'entrylog', 'index'",
	}
	out = append(out, successOut, argError, typeError)
	desc.CommandOutput = out

	vc.SetDescription(
		"listdiskfile",
		"Get all the files on the disk of the current bookie",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doListDiskFile(vc)
	}, "the file type is not specified or the file type is specified more than one")
}

func doListDiskFile(vc *cmdutils.VerbCmd) error {
	t, err := bkdata.ParseFileType(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewBookieClient()
	files, err := admin.Bookie().ListDiskFile(t)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), files)
	}

	return err
}
