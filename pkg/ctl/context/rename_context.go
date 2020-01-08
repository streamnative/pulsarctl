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
	"github.com/streamnative/pulsarctl/pkg/ctl/context/internal"
)

func RenameContextCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "rename-context CONTEXT_NAME NEW_NAME"
	desc.CommandPermission = "no-op"

	var examples []cmdutils.Example
	renameContext := cmdutils.Example{
		Desc:    "Rename the context 'old-name' to 'new-name' in your pulsarconfig file",
		Command: "pulsarctl context rename old-name new-name",
	}
	examples = append(examples, renameContext)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Context old_name renamed to new_name",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"rename",
		"Renames a context from the pulsarconfig file.",
		desc.ToString(),
		desc.ExampleToString(),
		"update")

	ops := new(renameContextOptions)
	ops.access = internal.NewDefaultPathOptions()

	// set the run function with name argument
	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doRunRenameContext(vc, ops)
	}, func(args []string) error {
		if len(args) != 2 {
			return errors.New("need two arguments apply to the command")
		}
		return nil
	})
}

type renameContextOptions struct {
	access internal.ConfigAccess
}

func doRunRenameContext(vc *cmdutils.VerbCmd, ops *renameContextOptions) error {
	oldName := vc.NameArgs[0]
	newName := vc.NameArgs[1]

	config, err := ops.access.GetStartingConfig()
	if err != nil {
		return err
	}

	configFile := ops.access.GetDefaultFilename()

	context, exists := config.Contexts[oldName]
	if !exists {
		return fmt.Errorf("cannot rename the context %q, it's not in %s", oldName, configFile)
	}

	_, newExists := config.Contexts[newName]
	if newExists {
		return fmt.Errorf("cannot rename the context %q, the context %q already exists in %s", oldName, newName, configFile)
	}

	config.Contexts[newName] = context
	delete(config.Contexts, oldName)

	if config.CurrentContext == oldName {
		config.CurrentContext = newName
	}

	if err := internal.ModifyConfig(ops.access, *config, true); err != nil {
		return err
	}

	vc.Command.Printf("Context %q renamed to %q.\n", oldName, newName)
	return nil
}
