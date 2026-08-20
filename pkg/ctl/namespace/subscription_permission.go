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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func GetSubPermissionsCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get permissions to access subscription admin API"
	desc.CommandPermission = "This command requires tenant admin permissions."

	desc.CommandExamples = []cmdutils.Example{
		{
			Desc:    "Get permissions to access subscription admin API",
			Command: "pulsarctl namespaces subscription-permission tenant/namespace",
		},
	}

	vc.SetDescription(
		"subscription-permission",
		"Get permissions to access subscription admin API",
		desc.ToString(),
		desc.ExampleToString(),
		"subscription-permission",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetSubPermissions(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
	vc.EnableOutputFlagSet()
}

func doGetSubPermissions(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	permissions, err := admin.Namespaces().GetSubPermissions(*ns)
	if err == nil {
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), cmdutils.NewOutputContent().WithObject(permissions))
	}
	return err
}
