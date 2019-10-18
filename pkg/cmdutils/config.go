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

package cmdutils

import (
	"log"
	"os"

	"github.com/streamnative/pulsarctl/pkg/auth"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/kris-nova/logger"
	"github.com/spf13/pflag"
)

var PulsarCtlConfig = ClusterConfig{}

// the configuration of the cluster that pulsarctl connects to
type ClusterConfig struct {
	// the web service url that pulsarctl connects to. Default is http://localhost:8080
	WebServiceURL string
	// Set the path to the trusted TLS certificate file
	TLSTrustCertsFilePath string
	// Configure whether the Pulsar client accept untrusted TLS certificate from broker (default: false)
	TLSAllowInsecureConnection bool

	AuthParams string

	// Token and TokenFile is used to config the pulsarctl using token to authentication
	Token     string
	TokenFile string
}

func (c *ClusterConfig) FlagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet(
		"PulsarCtl Config",
		pflag.ContinueOnError)

	flags.StringVarP(
		&c.WebServiceURL,
		"admin-service-url",
		"s",
		pulsar.DefaultWebServiceURL,
		"The admin web service url that pulsarctl connects to.")

	flags.StringVar(
		&c.AuthParams,
		"auth-params",
		"",
		"Authentication parameters are used to configure the public and private key files required by tls\n"+
			" For example: \"tlsCertFile:val1,tlsKeyFile:val2\"")

	flags.BoolVar(
		&c.TLSAllowInsecureConnection,
		"tls-allow-insecure",
		false,
		"Allow TLS insecure connection")

	flags.StringVar(
		&c.TLSTrustCertsFilePath,
		"tls-trust-cert-pat",
		"",
		"Allow TLS trust cert file path")

	flags.StringVar(
		&c.Token,
		"token",
		"",
		"Using the token to authentication")

	flags.StringVar(
		&c.TokenFile,
		"token-file",
		"",
		"Using the token file to authentication")

	return flags
}

func (c *ClusterConfig) Client(version pulsar.APIVersion) pulsar.Client {
	config := pulsar.DefaultConfig()

	if len(c.WebServiceURL) > 0 && c.WebServiceURL != config.WebServiceURL {
		config.WebServiceURL = c.WebServiceURL
	}

	if len(c.TLSTrustCertsFilePath) > 0 && c.TLSTrustCertsFilePath != config.TLSOptions.TrustCertsFilePath {
		config.TLSOptions.TrustCertsFilePath = c.TLSTrustCertsFilePath
	}

	if c.TLSAllowInsecureConnection {
		config.TLSOptions.AllowInsecureConnection = true
	}

	if len(c.AuthParams) > 0 && c.AuthParams != config.AuthParams {
		config.AuthParams = c.AuthParams
	}

	if len(c.Token) > 0 && len(c.TokenFile) > 0 {
		logger.Critical("the token and token file can not be specified at the same time")
		os.Exit(1)
	}

	if len(c.Token) > 0 || len(c.TokenFile) > 0 {
		if len(c.TLSTrustCertsFilePath) > 0 {
			logger.Critical("the token and tls can not be specified at the same time")
			os.Exit(1)
		}

		tokenParams := make(map[string]string)
		if len(c.Token) > 0 {
			tokenParams["token"] = c.Token
		}

		if len(c.TokenFile) > 0 {
			tokenParams["file"] = c.TokenFile
		}

		tokenAuth, err := auth.NewAuthenticationTokenWithParams(tokenParams)
		if err != nil {
			logger.Critical("%s\n", err.Error())
			os.Exit(1)
		}

		config.TokenAuth = tokenAuth
	}

	config.APIVersion = version

	client, err := pulsar.New(config)
	if err != nil {
		log.Fatalf("create pulsar client error: %s", err.Error())
	}
	return client
}
