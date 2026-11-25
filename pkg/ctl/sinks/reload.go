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

package sinks

import (
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func reloadSinksCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Reload built-in Pulsar IO sinks"
	desc.CommandPermission = "This command requires tenant admin permissions."

	desc.CommandExamples = []cmdutils.Example{
		{
			Desc:    "Reload built-in Pulsar IO sinks",
			Command: "pulsarctl sinks reload",
		},
	}

	vc.SetDescription("reload", "Reload built-in Pulsar IO sinks", desc.ToString(), desc.ExampleToString(), "reload")
	vc.SetRunFunc(func() error {
		return doReloadSinks(vc)
	})
}

func doReloadSinks(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	err := admin.Sinks().ReloadBuiltInSinks()
	if err == nil {
		vc.Command.Println("Reloaded built-in sinks successfully")
	}
	return err
}
