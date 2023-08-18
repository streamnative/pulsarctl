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

package oauth2

import (
	"errors"
	"strings"

	o "github.com/apache/pulsar-client-go/oauth2"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/auth"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func activateCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for activating a service account by supplying its credentials."
	desc.CommandPermission = "This command doesn't need pulsar permissions."

	var examples []cmdutils.Example
	activate := cmdutils.Example{
		Desc:    "Activate a service account by supplying its credentials.",
		Command: "pulsarctl oauth2 activate --key-file (key file path)",
	}
	examples = append(examples, activate)
	desc.CommandExamples = examples

	vc.SetDescription(
		"activate",
		"Activate a service account by supplying its credentials",
		desc.ToString(),
		desc.ExampleToString(),
		"activate")

	vc.SetRunFunc(func() error {
		return doActivate(vc, cmdutils.PulsarCtlConfig)
	})

	c := cmdutils.PulsarCtlConfig
	vc.FlagSetGroup.InFlagSet("OAuth 2.0", func(set *pflag.FlagSet) {
		set.StringVarP(&c.IssuerEndpoint, "issuer-endpoint", "i", c.IssuerEndpoint,
			"The OAuth 2.0 issuer endpoint")
		set.StringVarP(&c.Audience, "audience", "a", c.Audience,
			"The audience identifier for the Pulsar instance")
		set.StringVarP(&c.KeyFile, "key-file", "k", c.KeyFile,
			"The path to the private key file")
		set.StringVar(&c.Scope, "scope", c.Scope,
			"The OAuth 2.0 scope(s) to request")
		set.StringVar(
			&c.AuthParams,
			"auth-params",
			c.AuthParams,
			"Authentication parameters are used to configure the OAuth 2.0 provider.\n"+
				" OAuth2 example: \"{\"audience\":\"test\",\"issuerUrl\":\"https://sample\","+
				"\"privateKey\":\"/mnt/secrets/auth.json\",\"scope\":\"api://default/\"}\"\n")
	})
	vc.EnableOutputFlagSet()
}

func doActivate(vc *cmdutils.VerbCmd, config *cmdutils.ClusterConfig) error {
	config, err := applyClientCredentialsToConfig(config)
	if err != nil {
		return err
	}
	if config.KeyFile == "" {
		return errors.New("required: key-file")
	}
	if config.Audience == "" {
		return errors.New("required: audience")
	}

	flow, err := o.NewDefaultClientCredentialsFlow(o.ClientCredentialsFlowOptions{
		KeyFile:          config.KeyFile,
		AdditionalScopes: strings.Split(config.Scope, " "),
	})
	if err != nil {
		return err
	}

	grant, err := flow.Authorize(config.Audience)
	if err != nil {
		return err
	}

	store, err := auth.MakeKeyringStore()
	if err != nil {
		return err
	}

	err = store.SaveGrant(config.Audience, *grant)
	if err != nil {
		return err
	}

	userName, err := store.WhoAmI(config.Audience)
	if err != nil {
		return err
	}

	vc.Command.Printf(`Logged in as %s.
Welcome to Pulsar!
`, userName)
	return nil
}
