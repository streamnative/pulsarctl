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

package namespace

import (
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func SetIsAllowAutoUpdateSchemaCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Set the whether to allow auto update schema on a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions"

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Enable automatically update schema on a namespace",
		Command: "pulsarctl namespaces set-is-allow-auto-update-schema --enable (namespace-name)",
	})
	examples = append(examples, cmdutils.Example{
		Desc:    "Disable automatically update update schema on a namespace",
		Command: "pulsarctl namespaces set-is-allow-auto-update-schema --disable (namespace-name)",
	})
	desc.CommandExamples = examples

	vc.SetDescription(
		"set-is-allow-auto-update-schema",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString())

	var (
		enable  bool
		disable bool
	)
	vc.FlagSetGroup.InFlagSet("IsAllowAutoUpdateSchema", func(set *pflag.FlagSet) {
		set.BoolVar(&enable, "enable", false, "enable automatically update schema")
		set.BoolVar(&disable, "disable", false, "disable automatically update schema")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		if enable == disable {
			return errors.New("specify only one of --enable or --disable")
		}
		return doSetIsAllowAutoUpdateSchema(vc, enable)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doSetIsAllowAutoUpdateSchema(vc *cmdutils.VerbCmd, isAllowUpdateSchema bool) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Namespaces().SetIsAllowAutoUpdateSchema(*ns, isAllowUpdateSchema)
	if err == nil {
		action := "enable"
		if !isAllowUpdateSchema {
			action = "disable"
		}
		vc.Command.Printf("Successfully %s auto update schema on a namespace %s\n", action, ns.String())
	}
	return err
}
