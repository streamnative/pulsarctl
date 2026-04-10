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

package topicpolicies

import (
	"fmt"
	"io"

	adminpkg "github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func topicName(vc *cmdutils.VerbCmd) (*utils.TopicName, error) {
	if vc.NameError != nil {
		return nil, vc.NameError
	}
	return utils.GetTopicName(vc.NameArg)
}

func topicPolicies(global bool) (adminpkg.TopicPolicies, error) {
	return adminpkg.TopicPoliciesOf(cmdutils.NewPulsarClient(), global)
}

func addScopeFlags(vc *cmdutils.VerbCmd, global *bool, applied *bool) {
	vc.FlagSetGroup.InFlagSet("TopicPolicyScope", func(set *pflag.FlagSet) {
		if global != nil {
			set.BoolVarP(global, "global", "g", false, "use global topic policies")
		}
		if applied != nil {
			set.BoolVarP(applied, "applied", "a", false, "get the applied policy for the topic")
		}
	})
}

func writePolicyOutput(vc *cmdutils.VerbCmd, obj interface{}, text string, args ...interface{}) error {
	if vc.OutputConfig == nil {
		vc.EnableOutputFlagSet()
	}
	return vc.OutputConfig.WriteOutput(
		vc.Command.OutOrStdout(),
		cmdutils.NewOutputContent().
			WithObject(obj).
			WithTextFunc(func(w io.Writer) error {
				if text == "" {
					_, err := fmt.Fprintln(w, obj)
					return err
				}
				_, err := fmt.Fprintf(w, text, args...)
				return err
			}),
	)
}
