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
	"strconv"

	"github.com/streamnative/pulsar-admin-go/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetIsAllowAutoUpdateSchemaCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Get the whether to allow auto update schema on a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions"

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    desc.CommandUsedFor,
		Command: "pulsarctl namespaces get-is-allow-auto-update-schema (namespace-name)",
	})
	desc.CommandExamples = examples

	vc.SetDescription(
		"get-is-allow-auto-update-schema",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doGetIsAllowAutoUpdateSchema(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetIsAllowAutoUpdateSchema(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	result, err := admin.Namespaces().GetIsAllowAutoUpdateSchema(*ns)
	if err == nil {
		vc.Command.Println(strconv.FormatBool(result))
	}
	return err
}
