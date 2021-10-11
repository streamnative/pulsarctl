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
	"errors"
	"fmt"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func useContextCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "use-context CONTEXT_NAME"
	desc.CommandPermission = "This command does not need any permission"

	var examples []cmdutils.Example
	currentContext := cmdutils.Example{
		Desc:    "Use the context for the `test` cluster",
		Command: "pulsarctl context use test",
	}
	examples = append(examples, currentContext)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Switched to context (context name)",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"use",
		"use-context CONTEXT_NAME",
		desc.ToString(),
		desc.ExampleToString(),
		"use")

	ops := new(useContextOptions)
	ops.access = cmdutils.NewDefaultClientConfigLoadingRules()

	// set the run function with name argument
	vc.SetRunFuncWithNameArg(func() error {
		return doRunUseContext(vc, ops)
	}, "the context name is not specified or the context name is specified more than one")
}

type useContextOptions struct {
	access cmdutils.ConfigAccess
}

func doRunUseContext(vc *cmdutils.VerbCmd, ops *useContextOptions) error {
	config, err := ops.access.GetStartingConfig()
	if err != nil {
		return err
	}

	err = validate(vc, config)
	if err != nil {
		return err
	}

	config.CurrentContext = vc.NameArg

	err = cmdutils.ModifyConfig(ops.access, *config)
	if err == nil {
		vc.Command.Printf("Switched to context %q.\n", vc.NameArg)
	}

	return err
}

func validate(vc *cmdutils.VerbCmd, config *cmdutils.Config) error {
	if len(vc.NameArg) == 0 {
		return errors.New("empty context names are not allowed")
	}

	for name := range config.Contexts {
		if name == vc.NameArg {
			return nil
		}
	}

	return fmt.Errorf("no context exists with the name: %q", vc.NameArg)
}
