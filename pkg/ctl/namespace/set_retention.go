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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func setRetention(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Set the retention policy for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	setRetentionWithTime := pulsar.Example{
		Desc:    "Set the retention policy for a namespace",
		Command: "pulsarctl namespaces set-retention tenant/namespace --time 100m",
	}

	setRetentionWithSize := pulsar.Example{
		Desc:    "Set the retention policy for a namespace",
		Command: "pulsarctl namespaces set-retention tenant/namespace --size 1G",
	}
	examples = append(examples, setRetentionWithTime, setRetentionWithSize)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Set retention successfully for [tenant/namespace]",
	}

	notTenantName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}

	notExistTenantName := pulsar.Output{
		Desc: "the tenant name not exist, please check the tenant name",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	notExistNsName := pulsar.Output{
		Desc: "the namespace not exist, please check namespace name",
		Out:  "[✖]  code: 404 reason: Namespace <tenant/namespace> does not exist",
	}

	notSetBacklog := pulsar.Output{
		Desc: "Retention Quota must exceed configured backlog quota for namespace",
		Out:  "Retention Quota must exceed configured backlog quota for namespace",
	}

	//Retention Quota must exceed configured backlog quota for namespace

	out = append(out, successOut, notTenantName, notExistTenantName, notExistNsName, notSetBacklog)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-retention",
		"Set the retention policy for a namespace",
		desc.ToString(),
		"set-retention",
	)

	var data pulsar.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSetRetention(vc, data)
	})

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&data.RetentionTimeStr,
			"time",
			"",
			"Retention time in minutes (or minutes, hours,days,weeks eg: 100m, 3h, 2d, 5w).\n"+
				"0 means no retention and -1 means infinite time retention")

		flagSet.StringVar(
			&data.LimitStr,
			"size",
			"",
			"Retention size limit (eg: 10M, 16G, 3T).\n"+
				"0 or less than 1MB means no retention and -1 means infinite size retention")

		cobra.MarkFlagRequired(flagSet, "time")
		cobra.MarkFlagRequired(flagSet, "size")
	})
}

func doSetRetention(vc *cmdutils.VerbCmd, data pulsar.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	sizeLimit, err := validateSizeString(data.LimitStr)
	if err != nil {
		return err
	}
	retentionTimeInSecond, err := parseRelativeTimeInSeconds(data.RetentionTimeStr)
	if err != nil {
		return err
	}

	var (
		retentionTimeInMin int
		retentionSizeInMB  int
	)

	if retentionTimeInSecond != -1 {
		fmt.Println("retentionTimeInSecond: ", retentionTimeInSecond)
		retentionTimeInMin = int(retentionTimeInSecond.Minutes())
	} else {
		retentionTimeInMin = -1
	}

	if sizeLimit != -1 {
		retentionSizeInMB = int(sizeLimit / (1024 * 1024))
	} else {
		retentionSizeInMB = -1
	}
	err = admin.Namespaces().SetRetention(ns, pulsar.NewRetentionPolicies(retentionTimeInMin, retentionSizeInMB))
	if err == nil {
		vc.Command.Printf("Set retention successfully for [%s]", ns)
	}

	return err
}
