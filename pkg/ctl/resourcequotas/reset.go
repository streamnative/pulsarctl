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

package resourcequotas

import (
	"errors"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func resetNamespaceBundleResourceQuota(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Reset the specified namespace bundle's resource quota to default value."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	reset := cmdutils.Example{
		Desc:    "Reset the specified namespace bundle's resource quota to default value",
		Command: "pulsarctl resource-quotas reset (namespace name) (bundle range)",
	}

	examples = append(examples, reset)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Reset resource quota successful",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"reset",
		"Reset the specified namespace bundle's resource quota to default value.",
		desc.ToString(),
		desc.ExampleToString(),
		"clear")

	vc.SetRunFuncWithMultiNameArgs(func() error {
		return doResetNamespaceBundleResourceQuota(vc)
	}, func(args []string) error {
		if len(args) != 2 {
			return errors.New("need two arguments apply to the command")
		}
		return nil
	})
}

func doResetNamespaceBundleResourceQuota(vc *cmdutils.VerbCmd) error {
	namespace := vc.NameArgs[0]
	bundle := vc.NameArgs[1]
	admin := cmdutils.NewPulsarClient()

	nsName, err := utils.GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	err = admin.ResourceQuotas().ResetNamespaceBundleResourceQuota(nsName.String(), bundle)
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		vc.Command.Println("Reset resource quota successful")
	}

	return err
}
