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
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func setPersistence(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Set the persistence policy for a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example
	setPersistence := pulsar.Example{
		Desc: "Set the persistence policy for a namespace",
		Command: "pulsarctl namespaces set-persistence tenant/namespace \n" +
			"\t--ensemble-size 2 \n" +
			"\t--write-quorum-size 2 \n" +
			"\t--ack-quorum-size 2 \n" +
			"\t--ml-mark-delete-max-rate 2.0",
	}

	examples = append(examples, setPersistence)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Set the persistence policies successfully for [tenant/namespace]",
	}

	noNamespaceName := pulsar.Output{
		Desc: "you must specify a tenant/namespace name, please check if the tenant/namespace name is provided",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}

	tenantNotExistError := pulsar.Output{
		Desc: "the tenant does not exist",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	nsNotExistError := pulsar.Output{
		Desc: "the namespace does not exist",
		Out:  "[✖]  code: 404 reason: Namespace <tenant/namespace> does not exist",
	}

	errArgsForBk := pulsar.Output{
		Desc: "Bookkeeper Ensemble >= WriteQuorum >= AckQuoru, please c ",
		Out:  "code: 412 reason: Bookkeeper Ensemble >= WriteQuorum >= AckQuoru",
	}

	out = append(out, successOut, noNamespaceName, tenantNotExistError, nsNotExistError, errArgsForBk)
	desc.CommandOutput = out

	vc.SetDescription(
		"set-persistence",
		"Set the persistence policy for a namespace",
		desc.ToString(),
		"set-persistence",
	)

	var data pulsar.NamespacesData

	vc.SetRunFuncWithNameArg(func() error {
		return doSetPersistence(vc, data)
	})

	vc.FlagSetGroup.InFlagSet("Namespaces", func(flagSet *pflag.FlagSet) {
		flagSet.IntVarP(
			&data.BookkeeperEnsemble,
			"ensemble-size",
			"e",
			0,
			"Number of bookies to use for a topic")

		flagSet.IntVarP(
			&data.BookkeeperWriteQuorum,
			"write-quorum-size",
			"w",
			0,
			"How many writes to make of each entry")

		flagSet.IntVarP(
			&data.BookkeeperAckQuorum,
			"ack-quorum-size",
			"a",
			0,
			"Number of acks (garanteed copies) to wait for each entry")

		flagSet.Float64VarP(
			&data.ManagedLedgerMaxMarkDeleteRate,
			"ml-mark-delete-max-rate",
			"r",
			0,
			"Throttling rate of mark-delete operation (0 means no throttle)")

		cobra.MarkFlagRequired(flagSet, "ensemble-size")
		cobra.MarkFlagRequired(flagSet, "write-quorum-size")
		cobra.MarkFlagRequired(flagSet, "ack-quorum-size")
		cobra.MarkFlagRequired(flagSet, "ml-mark-delete-max-rate")
	})
}

func doSetPersistence(vc *cmdutils.VerbCmd, data pulsar.NamespacesData) error {
	ns := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	persistencePolicies := pulsar.NewPersistencePolicies(data.BookkeeperEnsemble, data.BookkeeperWriteQuorum,
		data.BookkeeperAckQuorum, data.ManagedLedgerMaxMarkDeleteRate)
	err := admin.Namespaces().SetPersistence(ns, persistencePolicies)
	if err == nil {
		vc.Command.Printf("Set the persistence policies successfully for [%s]", ns)
	}
	return err
}
