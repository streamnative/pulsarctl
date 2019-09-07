package cmdutils

import (
	`log`
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

var PulsarCtlConfig = ClusterConfig{}

// the configuration of the cluster that pulsarctl connects to
type ClusterConfig struct {
	// the web service url that pulsarctl connects to. Default is http://localhost:8080
	WebServiceUrl string
	// Set the path to the trusted TLS certificate file
	TlsTrustCertsFilePath string
	// Configure whether the Pulsar client accept untrusted TLS certificate from broker (default: false)
	TlsAllowInsecureConnection bool

	AuthParams string
}

func (c *ClusterConfig) FlagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet(
		"PulsarCtl Config",
		pflag.ContinueOnError)

	flags.StringVarP(
		&c.WebServiceUrl,
		"admin-service-url",
		"s",
		pulsar.DefaultWebServiceURL,
		"The admin web service url that pulsarctl connects to.")

	flags.StringVar(
		&c.AuthParams,
		"auth-params",
		"",
		"Authentication parameters are used to configure the public and private key files required by tls\n" +
			" For example: \"tlsCertFile:val1,tlsKeyFile:val2\"")

	flags.BoolVar(
		&c.TlsAllowInsecureConnection,
		"tls-allow-insecure",
		false,
		"Allow TLS insecure connection")

	flags.StringVar(
		&c.TlsTrustCertsFilePath,
		"tls-trust-cert-pat",
		"",
		"Allow TLS trust cert file path")

	return flags
}

func (c *ClusterConfig) Client(version pulsar.ApiVersion) pulsar.Client {
	config := pulsar.DefaultConfig()

	if len(c.WebServiceUrl) > 0 && c.WebServiceUrl != config.WebServiceUrl {
		config.WebServiceUrl = c.WebServiceUrl
	}

	if len(c.TlsTrustCertsFilePath) > 0 && c.TlsTrustCertsFilePath != config.TlsOptions.TrustCertsFilePath {
		config.TlsOptions.TrustCertsFilePath = c.TlsTrustCertsFilePath
	}

	if c.TlsAllowInsecureConnection {
		config.TlsOptions.AllowInsecureConnection = true
	}

	if len(c.AuthParams) > 0 && c.AuthParams != config.AuthParams {
		config.AuthParams = c.AuthParams
	}

	config.ApiVersion = version

	client, err := pulsar.New(config)
	if err != nil {
		log.Fatalf("create pulsar client error: %s", err.Error())
	}
	return client
}
