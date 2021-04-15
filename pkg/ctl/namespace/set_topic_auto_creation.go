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
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func setTopicAutoCreation(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set topic auto-creation config for a namespace, overriding broker settings"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	topicAutoCreation := cmdutils.Example{
		Desc: "Set topic auto-creation config for a namespace, overriding broker settings",
		Command: "pulsarctl namespaces set-topic-auto-creation tenant/namespace \n" +
			"\t--type partitioned \n" +
			"\t--partitions 2",
	}
	examples = append(examples, topicAutoCreation)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set topic auto-creation config successfully for [tenant/namespace]",
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
		"set-topic-auto-creation",
		"Set topic auto-creation for a namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"set-topic-auto-creation",
	)

	var disable bool
	var topicType string
	var partitions int

	vc.SetRunFuncWithNameArg(func() error {
		return doSetTopicAutoCreation(vc, disable, topicType, partitions)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("TopicAutoCreation", func(set *pflag.FlagSet) {
		set.BoolVar(&disable, "disable", false, "disable topic auto-creation")
		set.StringVar(&topicType, "type", "", "topic type to auto-create")
		set.IntVar(&partitions, "partitions", 0, "number of partitions on auto-created partitioned topics")
	})
}

func doSetTopicAutoCreation(vc *cmdutils.VerbCmd, disable bool, topicType string, partitions int) error {
	admin := cmdutils.NewPulsarClient()
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}

	config := utils.TopicAutoCreationConfig{
		Allow: !disable,
	}
	if !disable {
		parsedTopicType, err := utils.ParseTopicType(topicType)
		if err != nil {
			return err
		}
		config.Type = parsedTopicType
		config.Partitions = partitions
	}

	err = admin.Namespaces().SetTopicAutoCreation(*ns, config)
	if err == nil {
		vc.Command.Printf("Set topic auto-creation config successfully for [%s]\n", ns)
	}
	return err
}
