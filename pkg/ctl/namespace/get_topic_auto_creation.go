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

func getTopicAutoCreation(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Get topic auto-creation config for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	desc.CommandExamples = []cmdutils.Example{
		{
			Desc:    "Get topic auto-creation config for a namespace",
			Command: "pulsarctl namespaces get-auto-topic-creation tenant/namespace",
		},
	}

	vc.SetDescription(
		"get-auto-topic-creation",
		"Get topic auto-creation config for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"get-auto-topic-creation",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doGetTopicAutoCreation(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
	vc.EnableOutputFlagSet()
}

func doGetTopicAutoCreation(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	config, err := admin.Namespaces().GetTopicAutoCreation(*ns)
	if err == nil {
		err = vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), cmdutils.NewOutputContent().WithObject(config))
	}
	return err
}
