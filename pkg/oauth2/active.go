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

	o "github.com/apache/pulsar-client-go/oauth2"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/auth"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func activateCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for activating a service account by supplying its credentials."
	desc.CommandPermission = "This command doesn't need pulsar permissions."

	var examples []cmdutils.Example
	activate := cmdutils.Example{
		Desc:    "Activate a service account by supplying its credentials.",
		Command: "pulsarctl oauth2 activate --issuer-endpoint (issuer) --audience (audience) --key-file (key file path)",
	}
	examples = append(examples, activate)
	desc.CommandExamples = examples

	vc.SetDescription(
		"activate",
		"Activate a service account by supplying its credentials",
		desc.ToString(),
		desc.ExampleToString(),
		"activate")

	var issuerEndpoint, audience, keyFile string
	var scopes []string

	vc.SetRunFunc(func() error {
		return doActivate(vc, issuerEndpoint, audience, keyFile, scopes)
	})

	vc.FlagSetGroup.InFlagSet("Oauth2 login", func(set *pflag.FlagSet) {
		set.StringVarP(&issuerEndpoint, "issuer-endpoint", "i", "",
			"The OAuth 2.0 issuer endpoint")
		_ = set.MarkHidden("issuer-endpoint")
		set.StringVarP(&audience, "audience", "a", "",
			"The audience identifier for the Pulsar instance")
		set.StringVarP(&keyFile, "key-file", "k", "",
			"Path to the private key file")
		set.StringSliceVar(&scopes, "scope", []string{},
			"OAuth 2.0 scopes to request")
	})
}

func doActivate(vc *cmdutils.VerbCmd, issuerEndpoint, audience, keyFile string, scopes []string) error {
	if audience == "" || keyFile == "" {
		return errors.New("the arguments issuer-endpoint, audience, key-file can not be empty")
	}

	flow, err := o.NewDefaultClientCredentialsFlow(o.ClientCredentialsFlowOptions{
		KeyFile:          keyFile,
		AdditionalScopes: scopes,
	})
	if err != nil {
		return err
	}

	grant, err := flow.Authorize(audience)
	if err != nil {
		return err
	}

	store, err := auth.MakeKeyringStore()
	if err != nil {
		return err
	}

	err = store.SaveGrant(audience, *grant)
	if err != nil {
		return err
	}

	userName, err := store.WhoAmI(audience)
	if err != nil {
		return err
	}

	vc.Command.Printf(`Logged in as %s.
Welcome to Pulsar!
`, userName)
	return nil
}
