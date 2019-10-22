<!--

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.

-->

# Overview of pulsarctl

Pulsarctl is a command-line tool developed using the Go language based on the Pulsar REST API. It consists of the following two parts:

- Go admin API, the user can directly call the interface provided by Pulsarctl to perform operations related to pulsar.
- Command line tools, users can use Pulsarctl to implement functions similar to pulsar-admin.


This overview covers `pulsarctl` syntax, describes the command operations, and provides common examples.For details about each command, including all the supported flags and subcommands, see the [pulsarctl](link to pulsarctl website) reference documentation.


## Syntax

Use the following syntax to run pulsarctl commands from your terminal window:

```
pulsarctl [command] [sub-command] [NAME] [flags]
```

where command, sub-command and options are:

- **command:** Specifies the resources of pulsar, for example clusters, topics, namespaces, functions and so on.
- **sub-command:** Specifies the operation that you want to perform on one or more resources, for example create, get, delete, update, list.
- **NAME:** Specifies the name of the resource. Names are case-sensitive. For example:

```
pulsarctl topics list public/default
```

the `pulsar/default` means namespace name.

- **flags:** Specifies optional flags. For example, you can use the `-s` or `--admin-service-url` flags to specify the address and port of the admin web service url that pulsarctl connects to.

> NOTE: Flags that you specify from the command line override default values and any corresponding environment variables.

If you need help, just run `pulsarctl help` from the terminal window.


## Commands

The following table includes short descriptions and the general syntax for all of the pulsarctl commands:

Ccommand | Syntax | Description
---|---|---
topics | pulsarctl topics [sub-command] [NAME] [flags] | Operations on persistent topics
tenants |  pulsarctl tenants [sub-command] [NAME] [flags] | Operations about tenants
subscriptions |  pulsarctl subscriptions [sub-command] [NAME] [flags] | Operations about subscriptions
sources |  pulsarctl sources [sub-command] [NAME] [flags] | Interface for managing Pulsar IO Sources (ingress data into Pulsar)
sinks |  pulsarctl sinks [sub-command] [NAME] [flags] | Interface for managing Pulsar IO sinks (egress data from Pulsar)
schemas |  pulsarctl schemas [sub-command] [NAME] [flags] | Operations about schemas
namespaces |  pulsarctl namespaces [sub-command] [NAME] [flags] | Operations about namespaces
functions |  pulsarctl functions [sub-command] [NAME] [flags] | Interface for managing Pulsar Functions (lightweight, Lambda-style compute processes that work with Pulsar)
clusters |  pulsarctl clusters [sub-command] [NAME] [flags] | Operations about clusters
brokers |  pulsarctl brokers [sub-command] [NAME] [flags] | Operations about brokers
broker-stats |  pulsarctl broker-stats [sub-command] [NAME] [flags] | Operations to collect broker statistics
functions-worker |  pulsarctl functions-worker [sub-command] [NAME] [flags] | Operations to collect function-worker statistics
ns-isolation-policy |  pulsarctl ns-isolation-policy [sub-command] [NAME] [flags] | Operations about namespace isolation policy
resource-quotas |  pulsarctl resource-quotas [sub-command] [NAME] [flags] | Operations about resource quotas

> NOTE: For more about command operations, see the [pulsarctl](link to pulsarctl website) reference documentation.



## Security

Currently, the encryption methods supported by pulsarctl include the following two methods:

- JWT(Java Web Token)
- TLS

When you run `pulsarctl help` from the terminal window, you can get the following flags information:

```
Common flags:
  -s, --admin-service-url string    The admin web service url that pulsarctl connects to. (default "http://localhost:8080")
      --auth-params string          Authentication parameters are used to configure the public and private key files required by tls
                                     For example: "tlsCertFile:val1,tlsKeyFile:val2"
  -C, --color string                toggle colorized logs (true,false,fabulous) (default "true")
  -h, --help                        help for this command
      --tls-allow-insecure          Allow TLS insecure connection
      --tls-trust-cert-pat string   Allow TLS trust cert file path
      --token string                Using the token to authentication
      --token-file string           Using the token file to authentication
  -v, --verbose int                 set log level, use 0 to silence, 4 for debugging (default 3)
```

### How to enable TLS

#### Config client

As shown above, you can use the following four flags to use TLS:

Flags | Description
---|---
--admin-service-url | The admin web service url that pulsarctl connects to
--auth-params | Authentication parameters are used to configure the public and private key files required by tls
--tls-allow-insecur | Allow TLS insecure connection
--tls-trust-cert-pat | Allow TLS trust cert file path

For example:

```
pulsarctl \
    --admin-service-url https://localhost:8443 \
    --auth-params "{\"tlsCertFile\":\"/test/auth/certs/client-cert.pem\",\"tlsKeyFile\":\"/test/auth/certs/client-key.pem\"}" \
	--tls-allow-insecure \
	--tls-trust-cert-pat /test/auth/certs/cacert.pem \
	topics list public/default
```

#### Config broker

To enable TLS running in broker, you need to set `tlsEnabled` to `true` in the `broker.conf` or `standalone.conf` file. For examples:

```
# Enable TLS

tlsEnabled=true
tlsCertificateFilePath=/test/auth/certs/broker-cert.pem
tlsKeyFilePath=/test/auth/certs/broker-key.pem
tlsTrustCertsFilePath=/test/auth/certs/cacert.pem
tlsAllowInsecureConnection=false
```

For more detailed about TLS, see the [TLS](https://pulsar.apache.org/docs/en/security-tls-transport/) reference documentation.

### How to enable JWT

#### Config client

As shown above, you can use the following two flags to use JWT:


Flags | Description
---|---
--token | Using the token to authentication
--token-file | Using the token file to authentication

During use, `token` and `token-file` only need to specify one of them, `token` means the string value, `token-file` means the path of token file.

For example:

```
pulsarctl \
    --token "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXVzZXIifQ.Yb52IE0B5wzooAdSlIlskEgb6_HBXST8k3lINZS5wwg" \
    topics list public/default
```

Or

```
pulsarctl \
    --ttoken-file file:///test/auth/tokens/token
    topics list public/default
```

#### Config broker

To enable JWT running in broker, you need to set `authenticationEnabled` to `true` in the `broker.conf` or `standalone.conf` file. For examples:

```
# Enable JWT

# Enable authentication
authenticationEnabled=true

# Autentication provider name list, which is comma separated list of class names
authenticationProviders=org.apache.pulsar.broker.authentication.AuthenticationProviderToken

# Configure the secret key to be used to validate auth tokens
tokenSecretKey=file:///test/auth/token/secret.key
```

For more detailed about JWT, see the [JWT](https://pulsar.apache.org/docs/en/security-token-admin/) reference documentation.

## Admin API

The user can directly call the interface provided by Pulsarctl to perform operations related to pulsar.

### How to use admin api

`pulsarctl clusters list standalone` as an example, use way as follows:

#### Interface definition

```
// List returns the list of clusters
List() ([]string, error)
```

#### Source code example

```
package main

import (
	"fmt"
	"net/http"

	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func main() {
	config := &pulsar.Config{
		WebServiceURL: "http://localhost:8080",
		HTTPClient:    http.DefaultClient,

		// If the server enable the TLSAuth
		// Auth: auth.NewAuthenticationTLS()

		// If the server enable the TokenAuth
		// TokenAuth: auth.NewAuthenticationToken()
	}

	// the default NewPulsarClient will use v2 APIs. If you need to request other version APIs,
	// you can specified the API version like this:
	// admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V2)
	admin, err := pulsar.New(config)
	if err != nil {
		// handle the err
		return
	}

	// more APIs, you can find them in the pkg/pulsar/admin.go
	// You can find all the method in the pkg/pulsar
	clusters, err := admin.Clusters().List()
	if err != nil {
		// handle the error
		return
	}

	// handle the result
	fmt.Println(clusters)
}
```

For more detailed about admin api, see the [Admin API](https://godoc.org/github.com/streamnative/pulsarctl) reference documentation.


