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

package autorecovery

import (
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/pkg/errors"
)

func setLostBookieRecoveryDelayCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for setting the lost bookie recovery delay in second."
	desc.CommandPermission = "none"

	var examples []cmdutils.Example
	set := cmdutils.Example{
		Desc:    "Set the lost Bookie Recovery Delay",
		Command: "pulsarctl bookkeeper autorecovery setdelay (delay)",
	}
	examples = append(examples, set)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Successfully set the lost bookie recovery delay to (delay)(second)",
	}

	argError := cmdutils.Output{
		Desc: "the specified delay time is not specified or the delay time is specified more than one",
		Out:  "[✖]  the specified delay time is not specified or the delay time is specified more than one",
	}
	out = append(out, successOut, argError)
	desc.CommandOutput = out

	vc.SetDescription(
		"setdelay",
		"Set the lost bookie recovery delay",
		desc.ToString(),
		desc.ExampleToString())

	vc.SetRunFuncWithNameArg(func() error {
		return doLostBookieRecoveryDelay(vc)
	}, "the delay time is not specified or the delay time is specified more than one")
}

func doLostBookieRecoveryDelay(vc *cmdutils.VerbCmd) error {
	delay, err := strconv.Atoi(vc.NameArg)
	if err != nil {
		return errors.Errorf("invalid delay times %s", vc.NameArg)
	}

	admin := cmdutils.NewBookieClient()
	err = admin.AutoRecovery().SetLostBookieRecoveryDelay(delay)
	if err == nil {
		vc.Command.Printf("Successfully set the lost bookie recovery delay to %d(second)\n", delay)
	}

	return err
}
