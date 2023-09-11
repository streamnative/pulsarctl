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

func removeTopicAutoCreation(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Remove topic auto-creation config for a namespace, defaulting to broker settings"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	topicAutoCreation := cmdutils.Example{
		Desc:    "Remove topic auto-creation config for a namespace, defaulting to broker settings",
		Command: "pulsarctl namespaces remove-topic-auto-creation tenant/namespace",
	}
	examples = append(examples, topicAutoCreation)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Remove topic auto-creation config successfully for [tenant/namespace]",
	}

	noNamespaceName := cmdutils.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}

	tenantNotExistError := cmdutils.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := cmdutils.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError)
	desc.CommandOutput = out

	vc.SetDescription(
		"remove-topic-auto-creation",
		"Remove topic auto-creation for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"remove-topic-auto-creation",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveTopicAutoCreation(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doRemoveTopicAutoCreation(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}
	err = admin.Namespaces().RemoveTopicAutoCreation(*ns)
	if err == nil {
		vc.Command.Printf("Removed topic auto-creation config successfully for [%s]\n", ns)
	}
	return err
}
