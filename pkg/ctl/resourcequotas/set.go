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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func setResourceQuota(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Set the resource quota for specified namespace bundle, " +
		"or default quota if no namespace/bundle specified."
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the resource quota use default namespace/bundle",
		Command: "pulsarctl resource-quotas set",
	}
	setWithArgs := cmdutils.Example{
		Desc: "Set the resource quota for specified namespace bundle",
		Command: "pulsarctl resource-quotas set --namespace (namespace name) --bundle (bundle range)" +
			"--msgRateIn (msg rate in value)" +
			"--msgRateOut (msg rate out)" +
			"--bandwidthIn (bandwidth in)" +
			"--bandwidthOut (bandwidth out)" +
			"--memory (memory)" +
			"--dynamic",
	}
	examples = append(examples, set, setWithArgs)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set (default) resource quota successful",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"set",
		"Set the resource quota for specified namespace bundle, "+
			"or default quota if no namespace/bundle specified.",
		desc.ToString(),
		desc.ExampleToString(),
		"set")

	quotaData := &utils.ResourceQuotaData{}

	vc.SetRunFunc(func() error {
		return doSetResourceQuota(vc, quotaData)
	})

	vc.FlagSetGroup.InFlagSet("SchemaConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&quotaData.Names,
			"namespace",
			"n",
			"",
			"cluster/namespace, must be specified together with '--bundle'")
		flagSet.StringVarP(
			&quotaData.Bundle,
			"bundle",
			"b",
			"",
			"{start-boundary}_{end-boundary}, must be specified together with '--namespace'")
		flagSet.Int64Var(
			&quotaData.MsgRateIn,
			"msgRateIn",
			0,
			"expected incoming messages per second")
		flagSet.Int64Var(
			&quotaData.MsgRateOut,
			"msgRateOut",
			0,
			"expected outgoing messages per second")
		flagSet.Int64Var(
			&quotaData.BandwidthIn,
			"bandwidthIn",
			0,
			"expected inbound bandwidth (bytes/second)")
		flagSet.Int64Var(
			&quotaData.BandwidthOut,
			"bandwidthOut",
			0,
			"expected outbound bandwidth (bytes/second)")
		flagSet.Int64Var(
			&quotaData.Memory,
			"memory",
			0,
			"expected memory usage (Mbytes)")
		flagSet.BoolVar(
			&quotaData.Dynamic,
			"dynamic",
			false,
			"dynamic (allow to be dynamically re-calculated) or not")
		cobra.MarkFlagRequired(flagSet, "msgRateIn")
		cobra.MarkFlagRequired(flagSet, "MsgRateOut")
		cobra.MarkFlagRequired(flagSet, "bandwidthIn")
		cobra.MarkFlagRequired(flagSet, "bandwidthOut")
		cobra.MarkFlagRequired(flagSet, "memory")
	})
}

func doSetResourceQuota(vc *cmdutils.VerbCmd, quotaData *utils.ResourceQuotaData) error {
	var err error
	admin := cmdutils.NewPulsarClient()

	quota := utils.NewResourceQuota()
	quota.MsgRateIn = float64(quotaData.MsgRateIn)
	quota.MsgRateOut = float64(quotaData.MsgRateOut)
	quota.BandwidthIn = float64(quotaData.BandwidthIn)
	quota.BandwidthOut = float64(quotaData.BandwidthOut)
	quota.Memory = float64(quotaData.Memory)
	quota.Dynamic = quotaData.Dynamic

	switch {
	case quotaData.Bundle == "" && quotaData.Names == "":
		err = admin.ResourceQuotas().SetDefaultResourceQuota(*quota)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Println("Set default resource quota successful")
		}
	case quotaData.Bundle != "" && quotaData.Names != "":
		err = admin.ResourceQuotas().SetNamespaceBundleResourceQuota(quotaData.Names, quotaData.Bundle, *quota)
		if err != nil {
			cmdutils.PrintError(vc.Command.OutOrStderr(), err)
		} else {
			vc.Command.Println("Set resource quota successful")
		}
	default:
		return errors.New("Namespace and bundle must be provided together")
	}

	return err
}
