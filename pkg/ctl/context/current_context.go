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

package context

import (
	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func currentContextCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Displays the current-context"
	desc.CommandPermission = "This command does not need any permission"

	var examples []cmdutils.Example
	currentContext := cmdutils.Example{
		Desc:    "Provisions a new context",
		Command: "pulsarctl context current",
	}
	examples = append(examples, currentContext)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "(current context name)",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"current",
		"Displays the current context",
		desc.ToString(),
		desc.ExampleToString(),
		"current")

	ops := new(currentContextOptions)
	ops.access = cmdutils.NewDefaultClientConfigLoadingRules()

	// set the run function without name argument
	vc.SetRunFunc(func() error {
		return doRunCurrentContext(vc, ops)
	})
}

type currentContextOptions struct {
	access cmdutils.ConfigAccess
}

func doRunCurrentContext(vc *cmdutils.VerbCmd, ops *currentContextOptions) error {
	config, err := ops.access.GetStartingConfig()
	if err != nil {
		return err
	}

	if config.CurrentContext == "" {
		return errors.New("current-context is not set")
	}

	vc.Command.Printf("%s\n", config.CurrentContext)
	return nil
}
