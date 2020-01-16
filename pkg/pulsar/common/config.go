package common

type Config struct {
	// the web service url that pulsarctl connects to. Default is http://localhost:8080
	WebServiceURL string

	// the bookkeeper service url that pulsarctl connects to.
	BKWebServiceURL string
	// Set the path to the trusted TLS certificate file
	TLSTrustCertsFilePath string
	// Configure whether the Pulsar client accept untrusted TLS certificate from broker (default: false)
	TLSAllowInsecureConnection bool

	TLSEnableHostnameVerification bool

	AuthPlugin string

	AuthParams string

	// TLS Cert and Key Files for authentication
	TLSCertFile string
	TLSKeyFile string

	// Token and TokenFile is used to config the pulsarctl using token to authentication
	Token     string
	TokenFile string
	PulsarApiVersion APIVersion
}
