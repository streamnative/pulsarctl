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

package context

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/context/internal"
)

func SetContextCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Sets a context entry in pulsarconfig, " +
		"Specifying a name that already exists will merge new fields " +
		"on top of existing values for those fields."
	desc.CommandPermission = "no-op"

	var examples []cmdutils.Example
	setContext := cmdutils.Example{
		Desc:    "Sets the user field on the gce context entry without touching other values",
		Command: "pulsarctl context set [options]",
	}
	examples = append(examples, setContext)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Set context successful",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	// update the description
	vc.SetDescription(
		"set",
		"Sets a context entry in pulsarconfig",
		desc.ToString(),
		desc.ExampleToString(),
		"create")

	ops := new(createContextOptions)

	// set the run function with name argument
	vc.SetRunFuncWithNameArg(func() error {
		return doRunSetContext(vc, ops)
	}, "the context name is not specified or the context name is specified more than one")
}

func doRunSetContext(vc *cmdutils.VerbCmd, o *createContextOptions) error {
	name := vc.NameArg
	o.access = internal.NewDefaultPathOptions()

	config, err := o.access.GetStartingConfig()
	if err != nil {
		return err
	}

	startingStanza, exists := config.Contexts[name]
	if !exists {
		startingStanza = new(cmdutils.Context)
	}

	startingAuth, exists := config.AuthInfos[name]
	if !exists {
		startingAuth = new(cmdutils.AuthInfo)
	}

	context, authInfo := o.modifyContextConf(*startingStanza, *startingAuth)
	config.Contexts[name] = &context
	config.AuthInfos[name] = &authInfo

	if err := internal.ModifyConfig(o.access, *config, true); err != nil {
		return err
	}

	if exists {
		vc.Command.Printf("Context %q modified.\n", name)
	} else {
		vc.Command.Printf("Context %q created.\n", name)
	}

	return nil
}

type createContextOptions struct {
	access           internal.ConfigAccess
	authInfo         *cmdutils.AuthInfo
	brokerServiceURL string
	bookieServiceURL string
}

func (o *createContextOptions) modifyContextConf(existingContext cmdutils.Context,
	existingAuth cmdutils.AuthInfo) (cmdutils.Context, cmdutils.AuthInfo) {

	modifiedContext := existingContext
	modifiedAuth := existingAuth

	o.brokerServiceURL = cmdutils.PulsarCtlConfig.WebServiceURL
	o.bookieServiceURL = cmdutils.PulsarCtlConfig.BKWebServiceURL

	o.authInfo = new(cmdutils.AuthInfo)
	o.authInfo.TokenFile = cmdutils.PulsarCtlConfig.TokenFile
	o.authInfo.Token = cmdutils.PulsarCtlConfig.Token
	o.authInfo.TLSTrustCertsFilePath = cmdutils.PulsarCtlConfig.TLSTrustCertsFilePath
	o.authInfo.TLSAllowInsecureConnection = cmdutils.PulsarCtlConfig.TLSAllowInsecureConnection

	if o.authInfo != nil {
		modifiedAuth.TokenFile = o.authInfo.TokenFile
		modifiedAuth.Token = o.authInfo.Token
		modifiedAuth.TLSAllowInsecureConnection = o.authInfo.TLSAllowInsecureConnection
		modifiedAuth.TLSTrustCertsFilePath = o.authInfo.TLSTrustCertsFilePath
	}

	if o.brokerServiceURL != "" {
		modifiedContext.BrokerServiceURL = o.brokerServiceURL
	}

	if o.bookieServiceURL != "" {
		modifiedContext.BookieServiceURL = o.bookieServiceURL
	}

	return modifiedContext, modifiedAuth
}
