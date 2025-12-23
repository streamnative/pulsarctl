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
	"strings"

	o "github.com/apache/pulsar-client-go/oauth2"
	"github.com/spf13/pflag"

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
		Command: "pulsarctl oauth2 login",
	}
	examples = append(examples, login)
	desc.CommandExamples = examples

	vc.SetDescription(
		"login",
		"Login to oauth2 server",
		desc.ToString(),
		desc.ExampleToString(),
		"login")

	vc.SetRunFunc(func() error {
		return doLogin(vc, cmdutils.PulsarCtlConfig, false)
	})

	c := cmdutils.PulsarCtlConfig
	vc.FlagSetGroup.InFlagSet("OAuth 2.0", func(set *pflag.FlagSet) {
		set.StringVarP(&c.IssuerEndpoint, "issuer-endpoint", "i", c.IssuerEndpoint,
			"The OAuth 2.0 issuer endpoint")
		set.StringVarP(&c.Audience, "audience", "a", c.Audience,
			"The audience identifier for the Pulsar instance")
		set.StringVarP(&c.ClientID, "client-id", "c", c.ClientID,
			"The OAuth 2.0 client identifier for pulsarctl")
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

func whoAmI(grant *o.AuthorizationGrant) (string, error) {
	switch grant.Type {
	case o.GrantTypeClientCredentials:
		if grant.ClientCredentials == nil {
			return "", errors.New("authentication data is not usable")
		}
		return grant.ClientCredentials.ClientEmail, nil
	case o.GrantTypeDeviceCode:
		if grant.Token == nil {
			return "", errors.New("authentication data is not available")
		}
		return o.ExtractUserName(*grant.Token)
	default:
		return "", errors.New("authentication type is not supported")
	}
}

func doLogin(vc *cmdutils.VerbCmd, config *cmdutils.ClusterConfig, noRefresh bool) error {
	config, err := applyClientCredentialsToConfig(config)
	if err != nil {
		return err
	}
	if config.IssuerEndpoint == "" {
		return errors.New("required: issuer-endpoint")
	}
	if config.ClientID == "" {
		return errors.New("required: client-id")
	}
	if config.Audience == "" {
		return errors.New("required: audience")
	}

	options := o.DeviceCodeFlowOptions{
		IssuerEndpoint:   config.IssuerEndpoint,
		ClientID:         config.ClientID,
		AdditionalScopes: strings.Split(config.Scope, " "),
		AllowRefresh:     !noRefresh,
	}

	prompt := NewPrompt(false)
	flow, err := o.NewDefaultDeviceCodeFlow(options, prompt.Prompt)
	if err != nil {
		return errors.New("configuration error: unable to use device code flow: " + err.Error())
	}
	grant, err := flow.Authorize(config.Audience)
	if err != nil {
		return errors.New("login failed: " + err.Error())
	}

	userName, err := whoAmI(grant)
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
		var err error
		if code.VerificationURIComplete != "" {
			err = p.osInteractor.OpenURL(code.VerificationURIComplete)
		} else {
			err = p.osInteractor.OpenURL(code.VerificationURI)
		}
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
