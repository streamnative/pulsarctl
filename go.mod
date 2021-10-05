module github.com/streamnative/pulsarctl

go 1.13

require (
	github.com/99designs/keyring v1.1.6
	github.com/apache/pulsar-client-go v0.6.1-0.20211005052936-bfbb2a2eea0b // indirect
	github.com/apache/pulsar-client-go/oauth2 v0.0.0-20211005052936-bfbb2a2eea0b
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/fatih/color v1.7.0
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.4.3
	github.com/imdario/mergo v0.3.8
	github.com/kris-nova/logger v0.0.0-20181127235838-fd0d87064b06
	github.com/kris-nova/lolgopher v0.0.0-20180921204813-313b3abb0d9b
	github.com/magiconair/properties v1.8.0
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/olekukonko/tablewriter v0.0.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.5.1
	github.com/testcontainers/testcontainers-go v0.0.10
	golang.org/x/net v0.0.0-20210220033124-5f55cee0dc0d // indirect
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v2 v2.3.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/apache/pulsar-client-go => github.com/apache/pulsar-client-go v0.6.1-0.20211005052936-bfbb2a2eea0b
replace github.com/apache/pulsar-client-go/oauth2 => github.com/apache/pulsar-client-go/oauth2 v0.0.0-20211005052936-bfbb2a2eea0b
