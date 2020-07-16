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
	"fmt"

	o "github.com/apache/pulsar-client-go/oauth2"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/auth"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/oauth2/os"
)

func loginCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "This command is used for oauth2 user login."
	desc.CommandPermission = "This command doesn't need pulsar permissions."

	var examples []cmdutils.Example
	login := cmdutils.Example{
		Desc:    "Login as a oauth2 user",
		Command: "pulsarctl oauth2 login --issuer-endpoint (issuer) --audience (audience) --client-id (client-id)",
	}
	examples = append(examples, login)
	desc.CommandExamples = examples

	vc.SetDescription(
		"login",
		"Login to oauth2 server",
		desc.ToString(),
		desc.ExampleToString(),
		"login")

	var issuerEndpoint, clientID, audience, keyFile string

	vc.SetRunFunc(func() error {
		return doLogin(vc, issuerEndpoint, clientID, audience, false)
	})

	vc.FlagSetGroup.InFlagSet("Oauth2 login", func(set *pflag.FlagSet) {
		set.StringVarP(&issuerEndpoint, "issuer-endpoint", "i", "",
			"The OAuth 2.0 issuer endpoint")
		set.StringVarP(&audience, "audience", "a", "",
			"The audience identifier for the API server")
		set.StringVarP(&clientID, "client-id", "c", "",
			"The client ID to user for authorization grants")
		set.StringVarP(&keyFile, "key-file", "k", "",
			"Path to the private key file")
	})
}

func doLogin(vc *cmdutils.VerbCmd, issuerEndpoint, clientID, audience string, noRefresh bool) error {
	if issuerEndpoint == "" || clientID == "" || audience == "" {
		return errors.New("the arguments issuer-endpoint, client-id, audience can not be empty")
	}

	options := o.DeviceCodeFlowOptions{
		IssuerEndpoint:   issuerEndpoint,
		ClientID:         clientID,
		AdditionalScopes: nil,
		AllowRefresh:     noRefresh,
	}

	prompt := NewPrompt(false)
	flow, err := o.NewDefaultDeviceCodeFlow(options, prompt.Prompt)
	if err != nil {
		return errors.New("configuration error: unable to use device code flow: " + err.Error())
	}
	grant, err := flow.Authorize(audience)
	if err != nil {
		return errors.New("login failed: " + err.Error())
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

type PromptFunc struct {
	SkipOpen     bool
	osInteractor os.Interactor
}

func NewPrompt(skipOpen bool) *PromptFunc {
	return &PromptFunc{
		SkipOpen:     skipOpen,
		osInteractor: &os.DefaultInteractor{},
	}
}

func (p *PromptFunc) Prompt(code *o.DeviceCodeResult) error {
	if !p.SkipOpen {
		err := p.osInteractor.OpenURL(code.VerificationURIComplete)
		if err == nil {
			fmt.Printf(`We've launched your web browser to complete the login process.
Verification code: %s

Waiting for login to complete...
`, code.UserCode)
			return nil
		}
	}
	fmt.Printf(`Please follow these steps to complete the login procedure:
1. Using your web browser, go to: %s
2. Enter the following code: %s

Waiting for login to complete...
`, code.VerificationURI, code.UserCode)

	return nil
}
