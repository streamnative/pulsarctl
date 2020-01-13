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

package cmdutils

import (
	"os"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

// VerbCmd holds attributes that most of the commands use
type VerbCmd struct {
	Command      *cobra.Command
	FlagSetGroup *NamedFlagSetGroup
	NameArg      string
	NameArgs     []string
	OutputConfig  *OutputConfig

	// for testing
	NameError error
}

// AddVerbCmd create a registers a new command under the given resource command
func AddVerbCmd(flagGrouping *FlagGrouping, parentResourceCmd *cobra.Command, newVerbCmd func(*VerbCmd)) {
	verb := &VerbCmd{
		Command: &cobra.Command{},
		OutputConfig: &GlobalOutputConfig,
	}
	verb.FlagSetGroup = flagGrouping.New(verb.Command)
	newVerbCmd(verb)
	verb.FlagSetGroup.AddTo(verb.Command)
	parentResourceCmd.AddCommand(verb.Command)
}

func AddVerbCmds(flagGrouping *FlagGrouping, parentResourceCmd *cobra.Command, newVerbCmd ...func(cmd *VerbCmd)) {
	for _, cmd := range newVerbCmd {
		AddVerbCmd(flagGrouping, parentResourceCmd, cmd)
	}
}

// SetDescription sets usage along with short and long descriptions as well as aliases
func (vc *VerbCmd) SetDescription(use, short, long, example string, aliases ...string) {
	vc.Command.Use = use
	vc.Command.Short = short
	vc.Command.Long = long
	vc.Command.Aliases = aliases
	vc.Command.Example = example
}

// SetRunFunc registers a command function
func (vc *VerbCmd) SetRunFunc(cmd func() error) {
	vc.Command.Run = func(_ *cobra.Command, _ []string) {
		run(cmd)
	}
}

// SetRunFuncWithNameArg registers a command function with an optional name argument
func (vc *VerbCmd) SetRunFuncWithNameArg(cmd func() error, errMsg string) {
	vc.Command.Run = func(_ *cobra.Command, args []string) {
		vc.NameArg, vc.NameError = GetNameArg(args, errMsg)
		run(cmd)
	}
}

func (vc *VerbCmd) SetRunFuncWithMultiNameArgs(cmd func() error, checkArgs func(args []string) error) {
	vc.Command.Run = func(_ *cobra.Command, args []string) {
		vc.NameArgs, vc.NameError = GetNameArgs(args, checkArgs)
		run(cmd)
	}
}

var ExecErrorHandler = defaultExecErrorHandler

var defaultExecErrorHandler = func(err error) {
	logger.Critical("%s\n", err.Error())
	os.Exit(1)
}

func run(cmd func() error) {
	if err := cmd(); err != nil {
		ExecErrorHandler(err)
	}
}
