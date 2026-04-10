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
	"errors"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getSchemaValidationEnforceCmd(vc *cmdutils.VerbCmd) {
	vc.SetDescription("get-schema-validation-enforce", "Get schema validation enforce flag for a topic",
		"Get schema validation enforce flag for a topic", "")
	vc.SetRunFuncWithNameArg(func() error {
		topic, err := utils.GetTopicName(vc.NameArg)
		if err != nil {
			return err
		}
		admin := cmdutils.NewPulsarClient()
		enabled, err := admin.Topics().GetSchemaValidationEnforced(*topic)
		if err == nil {
			vc.Command.Printf("%t\n", enabled)
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}

func setSchemaValidationEnforceCmd(vc *cmdutils.VerbCmd) {
	var enable bool
	var disable bool
	vc.SetDescription("set-schema-validation-enforce", "Set schema validation enforce flag for a topic",
		"Set schema validation enforce flag for a topic", "")
	vc.FlagSetGroup.InFlagSet("SchemaValidationEnforce", func(set *pflag.FlagSet) {
		set.BoolVarP(&enable, "enable", "e", false, "enable schema validation enforce")
		set.BoolVarP(&disable, "disable", "d", false, "disable schema validation enforce")
	})
	vc.SetRunFuncWithNameArg(func() error {
		topic, err := utils.GetTopicName(vc.NameArg)
		if err != nil {
			return err
		}
		if enable == disable {
			return errors.New("need to specify either --enable or --disable")
		}
		admin := cmdutils.NewPulsarClient()
		err = admin.Topics().SetSchemaValidationEnforced(*topic, enable)
		if err == nil {
			vc.Command.Printf("Set schema validation enforce successfully for [%s]\n", topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}
