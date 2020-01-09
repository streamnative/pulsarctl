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
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/kris-nova/logger"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

var PulsarCtlConfig = ClusterConfig{}

// the configuration of the cluster that pulsarctl connects to
type ClusterConfig struct {
	// the web service url that pulsarctl connects to. Default is http://localhost:8080
	WebServiceURL string

	// the bookkeeper service url that pulsarctl connects to.
	BKWebServiceURL string
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

	c.addBKFlags(flags)

	return flags
}

func (c *ClusterConfig) addBKFlags(flags *pflag.FlagSet) {
	flags.StringVar(
		&c.BKWebServiceURL,
		"bookie-service-url",
		bookkeeper.DefaultWebServiceURL,
		"The bookie web service url that pulsarctl connects to.",
	)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func (c *ClusterConfig) DecodeContext() *Config {
	cfg := NewConfig()

	defaultPath := fmt.Sprintf("%s/.pulsar/config", utils.HomeDir())
	if !Exists(defaultPath) {
		return nil
	}

	content, err := ioutil.ReadFile(defaultPath)
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil
	}

	return cfg
}

func (c *ClusterConfig) Client(version common.APIVersion) pulsar.Client {
	config := pulsar.DefaultConfig()

	ctxConf := c.DecodeContext()
	if ctxConf.CurrentContext != "" {
		ctx := ctxConf.Contexts[ctxConf.CurrentContext]
		auth := ctxConf.AuthInfos[ctxConf.CurrentContext]

		c.WebServiceURL = ctx.BrokerServiceURL

		c.TLSTrustCertsFilePath = auth.TLSTrustCertsFilePath
		c.TLSAllowInsecureConnection = auth.TLSAllowInsecureConnection
		c.Token = auth.Token
		c.TokenFile = auth.TokenFile
	}

	if len(c.WebServiceURL) > 0 && c.WebServiceURL != config.WebServiceURL {
		config.WebServiceURL = c.WebServiceURL
	}

	if len(c.TLSTrustCertsFilePath) > 0 && c.TLSTrustCertsFilePath != config.TLSCertFile {
		config.TLSCertFile = c.TLSTrustCertsFilePath
	}

	if c.TLSAllowInsecureConnection {
		config.TLSAllowInsecureConnection = true
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
		config.TokenFile = c.TokenFile
		config.Token = c.Token
	}

	config.APIVersion = version

	client, err := pulsar.New(config)
	if err != nil {
		log.Fatalf("create pulsar client error: %s", err.Error())
	}
	return client
}

func (c *ClusterConfig) BookieClient() bookkeeper.Client {
	config := bookkeeper.DefaultConfig()
	ctxConf := c.DecodeContext()
	if ctxConf.CurrentContext != "" {
		ctx := ctxConf.Contexts[ctxConf.CurrentContext]
		c.BKWebServiceURL = ctx.BookieServiceURL
	}

	if len(c.BKWebServiceURL) > 0 {
		config.WebServiceURL = c.BKWebServiceURL
	}

	bk, err := bookkeeper.New(config)
	if err != nil {
		log.Fatalf("create bookie client error: %s", err.Error())
	}

	return bk
}
