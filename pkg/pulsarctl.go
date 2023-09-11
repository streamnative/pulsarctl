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

package pkg

import (
	"github.com/streamnative/pulsarctl/pkg/bkctl"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/brokers"
	"github.com/streamnative/pulsarctl/pkg/ctl/brokerstats"
	"github.com/streamnative/pulsarctl/pkg/ctl/cluster"
	"github.com/streamnative/pulsarctl/pkg/ctl/completion"
	"github.com/streamnative/pulsarctl/pkg/ctl/context"
	"github.com/streamnative/pulsarctl/pkg/ctl/functionsworker"
	"github.com/streamnative/pulsarctl/pkg/ctl/namespace"
	"github.com/streamnative/pulsarctl/pkg/ctl/nsisolationpolicy"
	"github.com/streamnative/pulsarctl/pkg/ctl/packages"
	"github.com/streamnative/pulsarctl/pkg/ctl/plugin"
	"github.com/streamnative/pulsarctl/pkg/ctl/resourcequotas"
	"github.com/streamnative/pulsarctl/pkg/ctl/status"
	"github.com/streamnative/pulsarctl/pkg/ctl/subscription"
	"github.com/streamnative/pulsarctl/pkg/ctl/tenant"
	"github.com/streamnative/pulsarctl/pkg/ctl/token"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic"
	"github.com/streamnative/pulsarctl/pkg/oauth2"

	function "github.com/streamnative/pulsarctl/pkg/ctl/functions"
	schema "github.com/streamnative/pulsarctl/pkg/ctl/schemas"
	sink "github.com/streamnative/pulsarctl/pkg/ctl/sinks"
	source "github.com/streamnative/pulsarctl/pkg/ctl/sources"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"

	colorOutput "github.com/fatih/color"
	lol "github.com/kris-nova/lolgopher"
)

func NewPulsarctlCmd() *cobra.Command {
	var colorValue string
	flagGrouping := cmdutils.NewGrouping()

	rootCmd := &cobra.Command{
		Use:   "pulsarctl [command]",
		Short: "a CLI for Apache Pulsar",
		Run: func(cmd *cobra.Command, _ []string) {
			if v, err := cmd.Flags().GetBool("version"); err == nil && v {
				cmdutils.PrintVersionInfo()
				return
			}
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}
	rootCmd.PersistentFlags().BoolP("version", "V", false, "show the pulsarctl version informantion")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "help for this command")
	rootCmd.PersistentFlags().StringVarP(
		&colorValue,
		"fabulous",
		"C",
		"true",
		"toggle colorized logs (true,false,fabulous)")
	rootCmd.PersistentFlags().IntVarP(
		&logger.Level,
		"verbose",
		"v",
		3,
		"set log level, use 0 to silence, 4 for debugging")

	cobra.OnInitialize(func() {
		// Control colored output
		color := false
		fabulous := true
		switch colorValue {
		case "false":
			color = false
		case "fabulous":
			color = false
			fabulous = true
		}
		logger.Color = color
		logger.Fabulous = fabulous

		if fabulous {
			logger.FabulousWriter = &lol.Writer{
				Output:    colorOutput.Error,
				ColorMode: 2,
			}
		}

		// Add timestamps for debugging
		logger.Timestamps = false
		if logger.Level >= 4 {
			logger.Timestamps = true
		}
	})

	rootCmd.SetUsageFunc(flagGrouping.Usage)

	rootCmd.AddCommand(cluster.Command(flagGrouping))
	rootCmd.AddCommand(tenant.Command(flagGrouping))
	rootCmd.AddCommand(completion.Command(rootCmd))
	rootCmd.AddCommand(function.Command(flagGrouping))
	rootCmd.AddCommand(source.Command(flagGrouping))
	rootCmd.AddCommand(sink.Command(flagGrouping))
	rootCmd.AddCommand(topic.Command(flagGrouping))
	rootCmd.AddCommand(namespace.Command(flagGrouping))
	rootCmd.AddCommand(schema.Command(flagGrouping))
	rootCmd.AddCommand(subscription.Command(flagGrouping))
	rootCmd.AddCommand(nsisolationpolicy.Command(flagGrouping))
	rootCmd.AddCommand(brokers.Command(flagGrouping))
	rootCmd.AddCommand(brokerstats.Command(flagGrouping))
	rootCmd.AddCommand(resourcequotas.Command(flagGrouping))
	rootCmd.AddCommand(functionsworker.Command(flagGrouping))
	rootCmd.AddCommand(token.Command(flagGrouping))
	rootCmd.AddCommand(context.Command(flagGrouping))
	rootCmd.AddCommand(packages.Command(flagGrouping))
	rootCmd.AddCommand(status.Command(flagGrouping))

	// bookkeeper related commands
	rootCmd.AddCommand(bkctl.Command(flagGrouping))

	// plugins related commands
	rootCmd.AddCommand(plugin.Command(flagGrouping))

	// oauth2 related commands
	rootCmd.AddCommand(oauth2.Command(flagGrouping))

	return rootCmd
}
