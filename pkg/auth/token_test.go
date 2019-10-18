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

// +build token

package auth

import (
	"bytes"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/cluster"
	"github.com/stretchr/testify/assert"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

func TestUseToken(t *testing.T) {
	args := []string{"--token",
		"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXVzZXIifQ.Yb52IE0B5wzooAdSlIlskEgb6_HBXST8k3lINZS5wwg",
		"clusters",
		"list",
	}
	execErr, err := BaseCmd(cluster.ListClustersCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	args = []string{
		"clusters",
		"list",
	}
	execErr, _ = BaseCmd(cluster.ListClustersCmd, args)
	assert.NotNil(t, execErr)
	assert.True(t, strings.Contains(execErr.Error(), "401"))
}

func BaseCmd(cmd func(cmd *cmdutils.VerbCmd), args []string) (execErr, err error) {
	cmdutils.ExecErrorHandler = func(err error) {
		execErr = err
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
	rootCmd.SetArgs(args)
	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")

	flagGrouping := cmdutils.NewGrouping()

	rootCmd.AddCommand(resourceCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, cmd)
	rootCmd.PersistentFlags().AddFlagSet(cmdutils.PulsarCtlConfig.FlagSet())
	err = rootCmd.Execute()
	return execErr, err
}
