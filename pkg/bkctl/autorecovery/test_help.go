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

package autorecovery

import (
	"bytes"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/spf13/cobra"
)

func testAutoRecoveryCommands(newVerb func(cmd *cmdutils.VerbCmd), args []string) (out *bytes.Buffer,
	execErr, nameErr, err error) {

	var execError error
	cmdutils.ExecErrorHandler = func(err error) {
		execError = err
	}

	var nameError error
	cmdutils.CheckNameArgError = func(err error) {
		nameError = err
	}

	var rootCmd = &cobra.Command{}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs(append([]string{"auto-recovery"}, args...))
	rootCmd.PersistentFlags().AddFlagSet(cmdutils.PulsarCtlConfig.FlagSet())

	resourceCmd := cmdutils.NewResourceCmd(
		"auto-recovery",
		"Operations about auto recovering",
		"")
	flagGrouping := cmdutils.NewGrouping()
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, newVerb)
	rootCmd.AddCommand(resourceCmd)
	err = rootCmd.Execute()

	return buf, execError, nameError, err
}
