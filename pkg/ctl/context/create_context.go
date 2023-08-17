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
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/bookkeeper"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/context/internal"
)

func setContextCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Sets a context entry in pulsarconfig, " +
		"Specifying a name that already exists will merge new fields " +
		"on top of existing values for those fields."
	desc.CommandPermission = "This command does not need any permission"

	var examples []cmdutils.Example
	setContext := cmdutils.Example{
		Desc:    "Sets the user field on the gce context entry without touching other values",
		Command: "pulsarctl context set [options]",
	}

	setClusterContext := cmdutils.Example{
		Desc: "Use set of context to define your cluster",
		Command: "pulsarctl context set development --admin-service-url=\"http://{host}:8080\"" +
			" --bookie-service-url=\"http://{host}:8083\"",
	}

	examples = append(examples, setContext, setClusterContext)
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
	ops.vc = vc
	ops.flags = &cmdutils.ClusterConfig{}

	// set the run function with name argument
	vc.SetRunFuncWithNameArg(func() error {
		return doRunSetContext(vc, ops)
	}, "the context name is not specified or the context name is specified more than one")

	vc.FlagSetGroup.InFlagSet("OAuth 2.0", func(set *pflag.FlagSet) {
		set.StringVarP(&ops.flags.IssuerEndpoint, "issuer-endpoint", "i", ops.flags.IssuerEndpoint,
			"The OAuth 2.0 issuer endpoint")
		set.StringVarP(&ops.flags.Audience, "audience", "a", ops.flags.Audience,
			"The audience identifier for the Pulsar instance")
		set.StringVarP(&ops.flags.ClientID, "client-id", "c", ops.flags.ClientID,
			"The OAuth 2.0 client identifier for pulsarctl")
		set.StringVarP(&ops.flags.KeyFile, "key-file", "k", ops.flags.KeyFile,
			"The path to the private key file")
		set.StringVar(&ops.flags.Scope, "scope", ops.flags.Scope,
			"The OAuth 2.0 scope(s) to request")
	})
	vc.ClusterConfigOverride = ops.flags
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
		startingStanza.BrokerServiceURL = admin.DefaultWebServiceURL
		startingStanza.BookieServiceURL = bookkeeper.DefaultWebServiceURL
	}

	startingAuth, exists := config.AuthInfos[name]
	if !exists {
		startingAuth = new(cmdutils.AuthInfo)
	}

	context, authInfo := o.modifyContextConf(*startingStanza, *startingAuth)
	config.Contexts[name] = &context
	config.AuthInfos[name] = &authInfo
	config.CurrentContext = name

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
	vc     *cmdutils.VerbCmd
	flags  *cmdutils.ClusterConfig
	access internal.ConfigAccess
}

func (o *createContextOptions) modifyContextConf(existingContext cmdutils.Context,
	existingAuth cmdutils.AuthInfo) (cmdutils.Context, cmdutils.AuthInfo) {

	f := o.vc.Command.Flags()
	modifiedContext := existingContext
	modifiedAuth := existingAuth

	if f.Changed("admin-service-url") {
		modifiedContext.BrokerServiceURL = o.flags.WebServiceURL
	}
	if f.Changed("bookie-service-url") {
		modifiedContext.BookieServiceURL = o.flags.BKWebServiceURL
	}
	if f.Changed("token-file") {
		modifiedAuth.TokenFile = o.flags.TokenFile
	}
	if f.Changed("token") {
		modifiedAuth.Token = o.flags.Token
	}
	if f.Changed("tls-trust-cert-path") {
		modifiedAuth.TLSTrustCertsFilePath = o.flags.TLSTrustCertsFilePath
	}
	if f.Changed("tls-allow-insecure") {
		modifiedAuth.TLSAllowInsecureConnection = o.flags.TLSAllowInsecureConnection
	}
	if f.Changed("issuer-endpoint") {
		modifiedAuth.IssuerEndpoint = o.flags.IssuerEndpoint
	}
	if f.Changed("client-id") {
		modifiedAuth.ClientID = o.flags.ClientID
	}
	if f.Changed("audience") {
		modifiedAuth.Audience = o.flags.Audience
	}
	if f.Changed("key-file") {
		modifiedAuth.KeyFile = o.flags.KeyFile
	}
	if f.Changed("scope") {
		modifiedAuth.Scope = o.flags.Scope
	}

	return modifiedContext, modifiedAuth
}
