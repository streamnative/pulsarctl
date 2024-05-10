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
	"log"
	"os"
	"strconv"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/auth"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/kris-nova/logger"
	"github.com/magiconair/properties"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper"
)

var PulsarCtlConfig = LoadFromEnv()

// the configuration of the cluster that pulsarctl connects to
type ClusterConfig config.Config

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
		"AuthPlugin is used to specify the plugin to use for authentication,"+
			" the supported values are \"org.apache.pulsar.client.impl.auth.AuthenticationTls\""+
			" and \"org.apache.pulsar.client.impl.auth.AuthenticationToken\"")

	flags.StringVar(
		&c.AuthParams,
		"auth-params",
		c.AuthParams,
		"Authentication parameters are used to configure the authentication provider specified by"+
			" \"AuthPlugin\"."+
			" Tls example: \"tlsCertFile:val1,tlsKeyFile:val2\""+
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

	return flags
}

func (c *ClusterConfig) addBKFlags(flags *pflag.FlagSet) {
	flags.StringVar(
		&c.BKWebServiceURL,
		"bookie-service-url",
		c.BKWebServiceURL,
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

func readConfigFile() (*Config, error) {
	cfg := NewConfig()

	defaultPath := fmt.Sprintf("%s/.config/pulsar/config", utils.HomeDir())
	if !Exists(defaultPath) {
		return nil, nil
	}

	content, err := os.ReadFile(defaultPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *ClusterConfig) ApplyContext(ctxConf *Config, contextName *string) {
	if ctxConf != nil {
		if contextName == nil {
			contextName = &ctxConf.CurrentContext
		}
		ctx, exist := ctxConf.Contexts[*contextName]
		if exist {
			c.WebServiceURL = ctx.BrokerServiceURL
			c.BKWebServiceURL = ctx.BookieServiceURL
		}
		authInfo, exist := ctxConf.AuthInfos[*contextName]
		if exist {
			c.TLSTrustCertsFilePath = authInfo.TLSTrustCertsFilePath
			c.TLSAllowInsecureConnection = authInfo.TLSAllowInsecureConnection
			c.Token = authInfo.Token
			c.TokenFile = authInfo.TokenFile
			c.IssuerEndpoint = authInfo.IssuerEndpoint
			c.ClientID = authInfo.ClientID
			c.Audience = authInfo.Audience
			c.KeyFile = authInfo.KeyFile
			c.Scope = authInfo.Scope
			if len(authInfo.KeyFile) != 0 {
				// auto active the OAuth2 if the key file is set
				c.AuthPlugin = auth.OAuth2PluginName
			}
		}
	}
}

func (c *ClusterConfig) Client(version config.APIVersion) Client {
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

	config := config.Config(*c)
	config.PulsarAPIVersion = version

	adminClient, err := admin.New(&config)
	if err != nil {
		logger.Critical("client error: %s", err.Error())
		os.Exit(1)
	}
	return &client{admin: adminClient}
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

func LoadFromEnv() *ClusterConfig {
	config := ClusterConfig{}
	if len(config.WebServiceURL) == 0 {
		config.WebServiceURL = admin.DefaultWebServiceURL
	}

	if envConf, ok := os.LookupEnv("PULSAR_CLIENT_CONF"); ok {
		if props, err := properties.LoadFile(envConf, properties.UTF8); err == nil && props != nil {
			config.WebServiceURL = props.GetString("webServiceUrl", admin.DefaultWebServiceURL)
			config.TLSAllowInsecureConnection = props.GetBool("tlsAllowInsecureConnection", false)
			config.TLSTrustCertsFilePath = props.GetString("tlsTrustCertsFilePath", "")
			config.BKWebServiceURL = props.GetString("brokerServiceUrl", bookkeeper.DefaultWebServiceURL)
			config.AuthParams = props.GetString("authParams", "")
			config.AuthPlugin = props.GetString("authPlugin", "")
			config.TLSEnableHostnameVerification = props.GetBool("tlsEnableHostnameVerification", false)
		}
	} else if clientFromEnv, ok := os.LookupEnv("PULSAR_CLIENT_FROM_ENV"); ok && clientFromEnv == "true" {
		if webServiceURL, ok := os.LookupEnv("webServiceUrl"); ok {
			config.WebServiceURL = webServiceURL
		}
		if tlsAllowInsecureConnection, ok := os.LookupEnv("tlsAllowInsecureConnection"); ok {
			config.TLSAllowInsecureConnection, _ = strconv.ParseBool(tlsAllowInsecureConnection)
		}
		if tlsTrustCertsFilePath, ok := os.LookupEnv("tlsTrustCertsFilePath"); ok {
			config.TLSTrustCertsFilePath = tlsTrustCertsFilePath
		}
		if brokerServiceURL, ok := os.LookupEnv("brokerServiceUrl"); ok {
			config.BKWebServiceURL = brokerServiceURL
		}
		if authParams, ok := os.LookupEnv("authParams"); ok {
			config.AuthParams = authParams
		}
		if authPlugin, ok := os.LookupEnv("authPlugin"); ok {
			config.AuthPlugin = authPlugin
		}
		if tlsEnableHostnameVerification, ok := os.LookupEnv("tlsEnableHostnameVerification"); ok {
			config.TLSEnableHostnameVerification, _ = strconv.ParseBool(tlsEnableHostnameVerification)
		}
	} else {
		ctxConf, err := readConfigFile()
		if err != nil {
			logger.Critical("configuration error: %s", err.Error())
			os.Exit(1)
		}
		config.ApplyContext(ctxConf, nil)
	}

	return &config
}
