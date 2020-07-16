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

package cluster

import (
	"bytes"
	"os"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

func TestClusterCommands(newVerb func(cmd *cmdutils.VerbCmd), args []string) (out *bytes.Buffer,
	execErr, nameErr, err error) {
	var execError error
	cmdutils.ExecErrorHandler = func(err error) {
		execError = err
	}

	var nameError error
	cmdutils.CheckNameArgError = func(err error) {
		nameError = err
	}

	var rootCmd = &cobra.Command{
		Use:   "pulsarctl [command]",
		Short: "a CLI for Apache Pulsar",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs(append([]string{"clusters"}, args...))

	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")
	flagGrouping := cmdutils.NewGrouping()
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, newVerb)
	rootCmd.AddCommand(resourceCmd)
	err = rootCmd.Execute()

	return buf, execError, nameError, err
}

var (
	flag bool
	// nolint
	basePath string
)

func TestTLSHelp(newVerb func(cmd *cmdutils.VerbCmd), args []string) (out *bytes.Buffer, execErr, err error) {
	var rootCmd = &cobra.Command{
		Use:   "pulsarctl [command]",
		Short: "a CLI for Apache Pulsar",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}

	cmdutils.ExecErrorHandler = func(err error) {
		execErr = err
	}

	if !flag {
		basePath, err = os.Getwd()
		if err != nil {
			return nil, nil, err
		}

		err = os.Chdir("../../../")
		if err != nil {
			return nil, nil, err
		}

		basePath, err = os.Getwd()
		if err != nil {
			return nil, nil, err
		}
		flag = true
	}

	baseArgs := []string{
		"--admin-service-url", "https://localhost:8443",
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs(append(baseArgs, args...))
	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")

	flagGrouping := cmdutils.NewGrouping()

	rootCmd.AddCommand(resourceCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, newVerb)
	rootCmd.PersistentFlags().AddFlagSet(cmdutils.PulsarCtlConfig.FlagSet())
	err = rootCmd.Execute()
	return buf, execErr, err
}
