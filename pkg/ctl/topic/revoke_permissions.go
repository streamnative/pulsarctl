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

package topic

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func RevokePermissions(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for revoking a client role permissions on a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic, partitioned topic"

	var examples []cmdutils.Example
	revoke := cmdutils.Example{
		Desc:    "Revoke permissions of a topic (topic-name)",
		Command: "pulsarctl topic revoke-permissions --role (role) (topic-name)",
	}
	examples = append(examples, revoke)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Revoke permissions for the role (role) of the topic (topic-name) successfully\n",
	}

	flagError := cmdutils.Output{
		Desc: "the specified role is empty",
		Out:  "Invalid role name",
	}
	out = append(out, successOut, flagError, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"revoke-permissions",
		"Revoke a client role permissions on a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"revoke")

	var role string

	vc.SetRunFuncWithNameArg(func() error {
		return doRevokePermissions(vc, role)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("RevokePermissions", func(set *pflag.FlagSet) {
		set.StringVar(&role, "role", "", "Client role to which revoke permissions")
		cobra.MarkFlagRequired(set, "role")
	})
}

func doRevokePermissions(vc *cmdutils.VerbCmd, role string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	if role == "" {
		return errors.New("Invalid role name")
	}
	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().RevokePermission(*topic, role)
	if err == nil {
		vc.Command.Printf("Revoke permissions for the role %s of "+
			"the topic %s successfully\n", role, topic.String())
	}

	return err
}
