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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/magiconair/properties"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"gopkg.in/yaml.v2"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"

	"github.com/kris-nova/logger"
	"github.com/spf13/pflag"
)

var PulsarCtlConfig = mustLoadConfig()

// the configuration of the cluster that pulsarctl connects to
type ClusterConfig common.Config

func (c *ClusterConfig) FlagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet(
		"PulsarCtl Config",
		pflag.ContinueOnError)

	flags.StringVarP(
		&c.WebServiceURL,
		"admin-service-url",
		"s",
		c.WebServiceURL,
		"The admin web service url that pulsarctl connects to.")

	flags.StringVar(
		&c.AuthPlugin,
		"auth-plugin",
		c.AuthPlugin,
		"AuthPlugin is used to specify the plugin to use for authentication,\n"+
			" the supported values are \"org.apache.pulsar.client.impl.auth.AuthenticationTls\"\n"+
			" and \"org.apache.pulsar.client.impl.auth.AuthenticationToken\"")

	flags.StringVar(
		&c.AuthParams,
		"auth-params",
		c.AuthParams,
		"Authentication parameters are used to configure the authentication provider specified by"+
			" \"AuthPlugin\".\n"+
			" Tls example: \"tlsCertFile:val1,tlsKeyFile:val2\"\n"+
			" Token example: \"authParams=file:///path/to/token/file\" or \"authParams=token:tokenVal\"")

	flags.BoolVar(
		&c.TLSAllowInsecureConnection,
		"tls-allow-insecure",
		c.TLSAllowInsecureConnection,
		"Allow TLS insecure connection")

	flags.BoolVar(
		&c.TLSEnableHostnameVerification,
		"tls-enable-hostname-verification",
		c.TLSEnableHostnameVerification,
		"Enable TLS hostname verification")

	flags.StringVar(
		&c.TLSTrustCertsFilePath,
		"tls-trust-cert-path",
		c.TLSTrustCertsFilePath,
		"Allow TLS trust cert file path")

	flags.StringVar(
		&c.Token,
		"token",
		c.Token,
		"Using the token to authentication")

	flags.StringVar(
		&c.TokenFile,
		"token-file",
		c.TokenFile,
		"Using the token file to authentication")

	flags.StringVar(
		&c.TLSCertFile,
		"tls-cert-file",
		c.TLSCertFile,
		"File path for TLS cert used for authentication")

	flags.StringVar(
		&c.TLSKeyFile,
		"tls-key-file",
		c.TLSKeyFile,
		"File path for TLS key used for authentication")

	c.addBKFlags(flags)
	c.addOAuth2Flags(flags)

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

func (c *ClusterConfig) addOAuth2Flags(flags *pflag.FlagSet) {
	flags.StringVar(
		&c.IssuerEndpoint,
		"issuer-endpoint",
		"",
		"OAuth 2.0 issuer endpoint.")

	flags.StringVar(
		&c.ClientID,
		"client-id",
		"",
		"OAuth 2.0 client identifier.")

	flags.StringVar(
		&c.Audience,
		"audience",
		"", "OAuth 2.0 audience identifier.")

	flags.StringVar(
		&c.KeyFile,
		"key-file",
		"", "Path to the private key file.")
}

func (c *ClusterConfig) Client(version common.APIVersion) pulsar.Client {
	c.PulsarAPIVersion = version

	if len(c.WebServiceURL) == 0 {
		c.WebServiceURL = pulsar.DefaultWebServiceURL
	}

	if len(c.Token) > 0 && len(c.TokenFile) > 0 {
		logger.Critical("the token and token file can not be specified at the same time")
		os.Exit(1)
	}

	if len(c.TLSKeyFile) > 0 && len(c.TLSCertFile) == 0 {
		logger.Critical("tls-cert-file provided but tls-key-file missing. Both must be provided for TLS auth")
		os.Exit(1)
	}
	if len(c.TLSCertFile) > 0 && len(c.TLSKeyFile) == 0 {
		logger.Critical("tls-key-file provided but tls-cert-file missing. Both must be provided for TLS auth")
		os.Exit(1)
	}

	config := common.Config(*c)
	client, err := pulsar.New(&config)
	if err != nil {
		log.Fatalf("new pulsar client failed: " + err.Error())
	}
	return client
}

func (c *ClusterConfig) BookieClient() bookkeeper.Client {
	config := bookkeeper.DefaultConfig()

	if len(c.BKWebServiceURL) > 0 {
		config.WebServiceURL = c.BKWebServiceURL
	}

	bk, err := bookkeeper.New(config)
	if err != nil {
		log.Fatalf("create bookie client error: %s", err.Error())
	}

	return bk
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func DecodeContext() *Config {
	cfg := NewConfig()
	defaultPath := filepath.Join(utils.HomeDir(), ".config", "pulsar", "config")
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

// mustLoadConfig loads config from system.
// Priority:
// 1. HOME/.config/pulsar/config
// 2. PULSAR_CLIENT_CONF
func mustLoadConfig() *ClusterConfig {
	config := ClusterConfig{}

	ctxConf := DecodeContext()
	if ctxConf != nil {
		if ctxConf.CurrentContext != "" {
			ctx, exist := ctxConf.Contexts[ctxConf.CurrentContext]
			auth, existAuth := ctxConf.AuthInfos[ctxConf.CurrentContext]

			if !exist || !existAuth {
				log.Fatalf("wrong context:%s\n"+
					"auth-info and contexts must be specified at the same time\n", ctxConf.CurrentContext)
			}

			config.WebServiceURL = ctx.BrokerServiceURL
			config.TLSTrustCertsFilePath = auth.TLSTrustCertsFilePath
			config.TLSAllowInsecureConnection = auth.TLSAllowInsecureConnection
			config.Token = auth.Token
			config.TokenFile = auth.TokenFile
			config.IssuerEndpoint = auth.IssuerEndpoint
			config.ClientID = auth.ClientID
			config.Audience = auth.Audience
			config.KeyFile = auth.KeyFile

			config.BKWebServiceURL = ctx.BookieServiceURL
		}
	}

	if envConf, ok := os.LookupEnv("PULSAR_CLIENT_CONF"); ok {
		if props, err := properties.LoadFile(envConf, properties.UTF8); err == nil && props != nil {
			config.WebServiceURL = props.GetString("webServiceUrl", pulsar.DefaultWebServiceURL)
			config.TLSAllowInsecureConnection = props.GetBool("tlsAllowInsecureConnection", false)
			config.TLSTrustCertsFilePath = props.GetString("tlsTrustCertsFilePath", "")
			config.BKWebServiceURL = props.GetString("brokerServiceUrl", bookkeeper.DefaultWebServiceURL)
			config.AuthParams = props.GetString("authParams", "")
			config.AuthPlugin = props.GetString("authPlugin", "")
			config.TLSEnableHostnameVerification = props.GetBool("tlsEnableHostnameVerification", false)
		}
	}

	return &config
}
