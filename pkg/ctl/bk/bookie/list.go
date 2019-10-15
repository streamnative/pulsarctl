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

package bookie

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func ListCmd(vc *cmdutils.VerbCmd) {
	var desc pulsar.LongDescription
	desc.CommandUsedFor = "This command is used for listing all the available bookies that type is the specified type."
	desc.CommandPermission = "none"

	var exampels []pulsar.Example
	list := pulsar.Example{
		Desc:    "List all the available bookies that type is the specified type",
		Command: "pulsarctl bk bookies list (type)",
	}

	showHostname := pulsar.Example{
		Desc:    "List all the available bookies that type is the specified type and print the hostname of the bookies",
		Command: "pulsarctl bk bookies list --show-hostname (type)",
	}
	exampels = append(exampels, list, showHostname)
	desc.CommandExamples = exampels

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: `{
	"bookieSocketAddress": "hostname",
}`,
	}

	argError := pulsar.Output{
		Desc: "the type is not specified or the type is specified more than one",
		Out:  "[✖]  the type is not specified or the type is specified more than one",
	}

	typeError := pulsar.Output{
		Desc: "the type is invalid",
		Out: "[✖]  invalid bookie type. the bookie type only can " +
			"be specified as 'rw' or 'ro'",
	}

	out = append(out, successOut, argError, typeError)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"List all the available bookies",
		desc.ToString(),
		desc.ExampleToString())

	var show bool

	vc.SetRunFuncWithNameArg(func() error {
		return doList(vc, show)
	}, "the type is not specified or the type is specified more than one")

	vc.FlagSetGroup.InFlagSet("List Bookie", func(set *pflag.FlagSet) {
		set.BoolVarP(&show, "show-hostname", "p", false,
			"show the hostname of the bookies")
	})
}

func doList(vc *cmdutils.VerbCmd, show bool) error {
	t, err := pulsar.ParseBookieType(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewBookieClient()
	bookies, err := admin.Bookie().List(t, show)
	if err == nil {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), bookies)
	}

	return err
}
