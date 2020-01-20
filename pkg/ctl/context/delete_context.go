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
	"fmt"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/context/internal"
)

func deleteContextCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Delete the specified context from the pulsarconfig"
	desc.CommandPermission = "This command does not need any permission"

	var examples []cmdutils.Example
	deleteContext := cmdutils.Example{
		Desc:    "Delete the context for the `test` cluster",
		Command: "pulsarctl context delete test",
	}
	examples = append(examples, deleteContext)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "deleted context (context-name) from (pulsarconfig)",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"delete",
		"delete context NAME",
		desc.ToString(),
		desc.ExampleToString(),
		"del")

	ops := new(deleteContextOptions)
	ops.access = internal.NewDefaultPathOptions()

	// set the run function with name argument
	vc.SetRunFuncWithNameArg(func() error {
		return doRunDeleteContext(vc, ops)
	}, "the context name is not specified or the context name is specified more than one")
}

type deleteContextOptions struct {
	access internal.ConfigAccess
}

func doRunDeleteContext(vc *cmdutils.VerbCmd, ops *deleteContextOptions) error {
	config, err := ops.access.GetStartingConfig()
	if err != nil {
		return err
	}

	name := vc.NameArg
	if len(name) == 0 {
		vc.Command.Help()
		return nil
	}

	configFile := ops.access.GetDefaultFilename()

	_, ok := config.Contexts[name]
	if !ok {
		return fmt.Errorf("cannot delete context %s, not in %s", name, configFile)
	}

	_, ok = config.AuthInfos[name]
	if !ok {
		return fmt.Errorf("cannot delete auth info %s, not in %s", name, configFile)
	}

	if config.CurrentContext == name {
		vc.Command.Printf("warning: this removed your active context, " +
			"use \"pulsarctl context use\" to select a different one\n")
	}

	delete(config.Contexts, name)
	delete(config.AuthInfos, name)

	if err := internal.ModifyConfig(ops.access, *config, true); err != nil {
		return err
	}

	vc.Command.Printf("deleted context %s from %s\n", name, configFile)

	return nil
}
