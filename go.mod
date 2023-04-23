module github.com/streamnative/pulsarctl

go 1.18

require (
	github.com/99designs/keyring v1.2.1
	github.com/apache/pulsar-client-go v0.9.0
	github.com/docker/go-connections v0.4.0
	github.com/fatih/color v1.7.0
	github.com/ghodss/yaml v1.0.0
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/golang/protobuf v1.5.2
	github.com/imdario/mergo v0.3.8
	github.com/kris-nova/logger v0.0.0-20181127235838-fd0d87064b06
	github.com/kris-nova/lolgopher v0.0.0-20180921204813-313b3abb0d9b
	github.com/magiconair/properties v1.8.5
	github.com/olekukonko/tablewriter v0.0.1
	github.com/onsi/gomega v1.19.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.0
	github.com/testcontainers/testcontainers-go v0.0.10
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-runewidth v0.0.4 // indirect
	golang.org/x/net v0.7.0 // indirect
)

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20220422013727-9388b58f7150
