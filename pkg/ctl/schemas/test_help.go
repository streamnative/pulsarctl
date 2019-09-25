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

package schemas

import (
	"bytes"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

func TestSchemasCommands(newVerb func(cmd *cmdutils.VerbCmd), args []string) (out *bytes.Buffer, execErr, err error) {
	var rootCmd = &cobra.Command{
		Use:   "pulsarctl [command]",
		Short: "a CLI for Apache Pulsar",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}

	var execError error
	cmdutils.ExecErrorHandler = func(err error) {
		execError = err
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs(append([]string{"schemas"}, args...))

	resourceCmd := cmdutils.NewResourceCmd(
		"schemas",
		"Operations about schemas",
		"",
		"schemas")
	flagGrouping := cmdutils.NewGrouping()
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, newVerb)
	rootCmd.AddCommand(resourceCmd)
	err = rootCmd.Execute()

	return buf, execError, err
}
