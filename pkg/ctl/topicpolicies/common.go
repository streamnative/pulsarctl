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
	"context"
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

func topicPolicyResources(vc *cmdutils.VerbCmd, global bool) (adminpkg.TopicPolicies, *utils.TopicName, error) {
	topic, err := topicName(vc)
	if err != nil {
		return nil, nil, err
	}
	policies, err := topicPolicies(global)
	if err != nil {
		return nil, nil, err
	}
	return policies, topic, nil
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
	oc := cmdutils.NewOutputContent().WithObject(obj)
	if text == "" {
		return vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
	}
	return vc.OutputConfig.WriteOutput(
		vc.Command.OutOrStdout(),
		oc.
			WithTextFunc(func(w io.Writer) error {
				_, err := fmt.Fprintf(w, text, args...)
				return err
			}),
	)
}

func getOptionalIntPolicyCmd(
	vc *cmdutils.VerbCmd,
	use string,
	short string,
	getter func(context.Context, adminpkg.TopicPolicies, utils.TopicName, bool) (*int, error),
) {
	var global bool
	var applied bool
	vc.SetDescription(use, short, short, "", use)
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		value, err := getter(vc.Command.Context(), policies, *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, value, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func getOptionalInt64PolicyCmd(
	vc *cmdutils.VerbCmd,
	use string,
	short string,
	getter func(context.Context, adminpkg.TopicPolicies, utils.TopicName, bool) (*int64, error),
) {
	var global bool
	var applied bool
	vc.SetDescription(use, short, short, "", use)
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		value, err := getter(vc.Command.Context(), policies, *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, value, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func getOptionalBoolPolicyCmd(
	vc *cmdutils.VerbCmd,
	use string,
	short string,
	getter func(context.Context, adminpkg.TopicPolicies, utils.TopicName, bool) (*bool, error),
) {
	var global bool
	var applied bool
	vc.SetDescription(use, short, short, "", use)
	addScopeFlags(vc, &global, &applied)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		value, err := getter(vc.Command.Context(), policies, *topic, applied)
		if err != nil {
			return err
		}
		return writePolicyOutput(vc, value, "")
	}, "the topic name is not specified or the topic name is specified more than one")
}

func setEnableDisablePolicyCmd(
	vc *cmdutils.VerbCmd,
	use string,
	short string,
	setter func(context.Context, adminpkg.TopicPolicies, utils.TopicName, bool) error,
) {
	var global bool
	var enable bool
	var disable bool
	vc.SetDescription(use, short, short, "", use)
	addScopeFlags(vc, &global, nil)
	vc.FlagSetGroup.InFlagSet("Policy", func(set *pflag.FlagSet) {
		set.BoolVarP(&enable, "enable", "e", false, "enable policy")
		set.BoolVarP(&disable, "disable", "d", false, "disable policy")
	})
	vc.SetRunFuncWithNameArg(func() error {
		if enable == disable {
			return fmt.Errorf("need to specify either --enable or --disable")
		}
		policies, topic, err := topicPolicyResources(vc, global)
		if err != nil {
			return err
		}
		err = setter(vc.Command.Context(), policies, *topic, enable)
		if err == nil {
			vc.Command.Printf("%s successfully for [%s]\n", short, topic.String())
		}
		return err
	}, "the topic name is not specified or the topic name is specified more than one")
}
