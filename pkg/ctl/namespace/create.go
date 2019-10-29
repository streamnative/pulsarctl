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
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

const MaxBundles = int64(1) << 32

func createNs(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Creates a new namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	create := pulsar.Example{
		Desc:    "creates a namespace named (namespace-name)",
		Command: "pulsarctl namespaces create (namespace-name)",
	}
	examples = append(examples, create)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Created (namespace-name) successfully",
	}

	noNamespaceName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  the namespace name is not specified or the namespace name is specified more than one",
	}

	tenantNotExistError := pulsar.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := pulsar.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace (tenant/namespace) does not exist",
	}

	positiveBundleErr := pulsar.Output{
		Desc: "Invalid number of bundles, please check --bundles value",
		Out:  "Invalid number of bundles. Number of numBundles has to be in the range of (0, 2^32].",
	}

	out = append(out, successOut, tenantNotExistError, noNamespaceName, nsNotExistError, positiveBundleErr)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Create a new namespace",
		desc.ToString(),
		desc.ExampleToString(),
		"create",
	)

	var namespaceData pulsar.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doCreate(vc, namespaceData)
	}, "the namespace name is not specified or the namespace name is specified more than one")

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.IntVarP(
			&namespaceData.NumBundles,
			"bundles",
			"b",
			0,
			"number of bundles to activate")

		flagSet.StringSliceVarP(
			&namespaceData.Clusters,
			"clusters",
			"c",
			nil,
			"List of clusters this namespace will be assigned")
	})
}

func doCreate(vc *cmdutils.VerbCmd, data pulsar.NamespacesData) error {
	tenantAndNamespace := vc.NameArg
	admin := cmdutils.NewPulsarClient()

	if data.NumBundles < 0 || data.NumBundles > int(MaxBundles) {
		return errors.New("invalid number of bundles. Number of numBundles has to be in the range of (0, 2^32]")
	}

	ns, err := pulsar.GetNamespaceName(tenantAndNamespace)
	if err != nil {
		return err
	}
	policies := pulsar.NewDefaultPolicies()
	if data.NumBundles > 0 {
		policies.Bundles = pulsar.NewBundlesDataWithNumBundles(data.NumBundles)
	}

	if data.Clusters != nil {
		policies.ReplicationClusters = data.Clusters
	}

	err = admin.Namespaces().CreateNsWithPolices(ns.String(), *policies)
	if err == nil {
		vc.Command.Printf("Created %s successfully\n", ns.String())
	}
	return err
}
