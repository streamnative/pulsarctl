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

package plugin

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/plugin"
)

func listCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription

	desc.CommandUsedFor = "This command is used for listing all plugins."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	list := cmdutils.Example{
		Desc:    "List all the plugins.",
		Command: "pulsarctl plugin list",
	}
	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "List all the plugins successfully.",
		Out:  "",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all the plugins.",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFunc(func() error {
		return doListPlugins(vc)
	})

	vc.EnableOutputFlagSet()
}

func doListPlugins(vc *cmdutils.VerbCmd) error {
	paths := filepath.SplitList(os.Getenv("PATH"))
	plugins := []string{}

	for _, dir := range paths {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if !hasValidPrefix(f.Name(), plugin.ValidPluginFilenamePrefixes) {
				continue
			}
			plugins = append(plugins, trimPrefix(f.Name(), plugin.ValidPluginFilenamePrefixes))
		}
	}

	oc := cmdutils.NewOutputContent().
		WithObject(plugins).
		WithTextFunc(func(w io.Writer) error {
			table := tablewriter.NewWriter(w)
			table.SetHeader([]string{"plugins"})
			for _, v := range plugins {
				table.Append([]string{v})
			}
			table.Render()
			return nil
		})

	return vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
}

func hasValidPrefix(filepath string, validPrefixes []string) bool {
	for _, prefix := range validPrefixes {
		if !strings.HasPrefix(filepath, prefix+"-") {
			continue
		}
		return true
	}
	return false
}

func trimPrefix(fileName string, validPrefixes []string) string {
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(fileName, prefix) {
			return strings.TrimPrefix(fileName, prefix+"-")
		}
	}
	return ""
}
