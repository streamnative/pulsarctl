package cmdutils


import (
	"github.com/spf13/pflag"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

var PulsarCtlConfig = ClusterConfig{}

// the configuration of the cluster that pulsarctl connects to
type ClusterConfig struct {
	// the web service url that pulsarctl connects to. Default is http://localhost:8080
	WebServiceUrl string
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
	return flags
}

func (c *ClusterConfig) Client(version pulsar.ApiVersion) pulsar.Client {
	config := pulsar.DefaultConfig()

	if len(c.WebServiceUrl) > 0 && c.WebServiceUrl != config.WebServiceUrl {
		config.WebServiceUrl = c.WebServiceUrl
	}

	config.ApiVersion = version

	return pulsar.New(config)
}

