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

package nsisolationpolicy

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

var checkNsIsolationPolicyArgs = func(args []string) error {
	if len(args) != 2 {
		return errors.New("need to specified the cluster name and the policy name")
	}
	return nil
}

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"ns-isolation-policy",
		"Operations about namespace isolation policy",
		"",
		"ns-isolation-policy")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getNsIsolationPolicy)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getNsIsolationPolicies)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteNsIsolationPolicy)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getBrokerWithPolicies)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getAllBrokersWithPolicies)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, setPolicy)

	return resourceCmd
}
