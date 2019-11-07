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
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func GrantPermissionCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for granting permissions to a client role on a topic."
	desc.CommandPermission = "This command requires namespace admin permissions."
	desc.CommandScope = "non-partitioned topic, a partition of a partitioned topic, partitioned topic"

	var examples []cmdutils.Example
	grant := cmdutils.Example{
		Desc:    "Grant permissions to a client on a single topic (topic-name)",
		Command: "pulsarctl topic grant-permissions --role (role) --actions (action-1) --actions (action-2) (topic-name)",
	}
	examples = append(examples, grant)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Grant role %s and actions %v to the topic %s successfully",
	}

	flagError := cmdutils.Output{
		Desc: "the specified role is empty",
		Out:  "Invalid role name",
	}

	actionsError := cmdutils.Output{
		Desc: "the specified actions is not allowed.",
		Out: "The auth action  only can be specified as 'produce', " +
			"'consume', or 'functions'. Invalid auth action '(actions)'",
	}
	out = append(out, successOut, ArgError, flagError, actionsError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"grant-permissions",
		"Grant permissions to a client on a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"grant")

	var role string
	var actions []string

	vc.SetRunFuncWithNameArg(func() error {
		return doGrantPermission(vc, role, actions)
	}, "the topic name is not specified or the topic name is specified more than one")

	vc.FlagSetGroup.InFlagSet("GrantPermissions", func(set *pflag.FlagSet) {
		set.StringVar(&role, "role", "",
			"Client role to which grant permissions")
		set.StringSliceVar(&actions, "actions", []string{},
			"Actions to be granted (produce,consume,functions)")
		cobra.MarkFlagRequired(set, "role")
		cobra.MarkFlagRequired(set, "actions")
	})
}

func doGrantPermission(vc *cmdutils.VerbCmd, role string, actions []string) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	if role == "" {
		return errors.New("Invalid role name")
	}

	authActions, err := getAuthActions(actions)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().GrantPermission(*topic, role, authActions)
	if err == nil {
		vc.Command.Printf(
			"Grant permissions for the role %s and the actions %v to "+
				"the topic %s successfully\n", role, actions, topic.String())
	}

	return err
}

func getAuthActions(actions []string) ([]common.AuthAction, error) {
	authActions := make([]common.AuthAction, 0)
	for _, v := range actions {
		a, err := common.ParseAuthAction(v)
		if err != nil {
			return nil, err
		}
		authActions = append(authActions, a)
	}
	return authActions, nil
}
