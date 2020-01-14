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

`pulsarctl` is a command-line tool developed in Go based on Pulsar REST API. 

It consists of the following two parts:

- Go admin API: users can directly call the interface provided by `pulsarctl` to perform operations related to Pulsar.
  
- Command-line tools: users can use `pulsarctl` to implement functions, which is similar to use pulsar-admin.

## Syntax

Use the following syntax to run `pulsarctl` command from your terminal window:

```
pulsarctl [resource] [command] [name] [flags]
```

The following are syntax descriptions.

### `resource`

Specifies the resources. This argument is required.

The following table includes the descriptions for `resource`.

`resource` | Syntax | Description
---|---|---
topics | `pulsarctl topics [sub-command] [name] [flags]` | Operations on persistent topics
tenants |  `pulsarctl tenants [sub-command] [name] [flags]` | Operations on tenants
subscriptions |  `pulsarctl subscriptions [sub-command] [name] [flags]` | Operations on subscriptions
sources |  `pulsarctl sources [sub-command] [name] [flags]` | Interface for managing Pulsar IO s ources (ingress data into Pulsar)
sinks |  `pulsarctl sinks [sub-command] [name] [flags]` | Interface for managing Pulsar IO sinks (egress data from Pulsar)
schemas |  `pulsarctl schemas [sub-command] [name] [flags]` | Operations on schemas
namespaces |  `pulsarctl namespaces [sub-command] [name] [flags]` | Operations on namespaces
functions |  `pulsarctl functions [sub-command] [name] [flags]` | Interface for managing Pulsar Functions (lightweight, Lambda-style compute processes that work with Pulsar)
clusters |  `pulsarctl clusters [sub-command] [name] [flags]` | Operations on clusters
brokers |  `pulsarctl brokers [sub-command] [name] [flags]` | Operations on brokers
broker-stats |  `pulsarctl broker-stats [sub-command] [name] [flags]` | Collect broker statistics
functions-worker |  `pulsarctl functions-worker [sub-command] [name] [flags]` | Collect function-worker statistics
ns-isolation-policy |  `pulsarctl ns-isolation-policy [sub-command] [name] [flags]` | Operations on namespace isolation policy
resource-quotas |  `pulsarctl resource-quotas [sub-command] [name] [flags]` | Operations on resource quotas

### `command`

Specifies the operation to be performed on one or more resources. This argument is required.

For example, create, get, delete, update, list, and so on.

### `name`

Specifies the name of the `resource`. This argument is required.

>#### Note
>
> `name` is case-sensitive. 

For example, `pulsarctl topics list public/default`, where `pulsar/default` is the namespace name.

### `flags`

Specifies the flags. This argument is optional. 

For example, you can use the `-s` or `--admin-service-url` flags to specify the address and port of the admin web service URL that pulsarctl connects to.

> #### NOTE
> 
> Flags that you specify from the command line override the default values and corresponding environment variables.

> #### Tip
> 
> * If you need help, run `pulsarctl help` from the terminal window.
> * For more information about pulsarctl, see [pulsarctl](link to pulsarctl website).

## Security

Currently, the encryption methods supported by pulsarctl are **TLS** and **JWT** (Java Web Token), and you can enable one of them.

### How to enable TLS

You need to configure a broker or a client.

#### Configure a broker

To enable TLS running in a broker, you need to set `tlsEnabled` to `true` in the `broker.conf` or `standalone.conf` file. 

**Example**

```
# Enable TLS

tlsEnabled=true
tlsCertificateFilePath=/test/auth/certs/broker-cert.pem
tlsKeyFilePath=/test/auth/certs/broker-key.pem
tlsTrustCertsFilePath=/test/auth/certs/cacert.pem
tlsAllowInsecureConnection=false
```

#### Configure a client

After running `pulsarctl help`, you can get the following information:

```
Common flags:
  -s, --admin-service-url string    The admin web service url that pulsarctl connects to. (default "http://localhost:8080")
      --auth-params string          Authentication parameters are used to configure the public and private key files required by tls
                                     For example: "tlsCertFile:val1,tlsKeyFile:val2"
  -C, --color string                toggle colorized logs (true,false,fabulous) (default "true")
  -h, --help                        help for this command
      --tls-allow-insecure          Allow TLS insecure connection
      --tls-trust-cert-path string   Allow TLS trust cert file path
      --token string                Using the token to authentication
      --token-file string           Using the token file to authentication
  -v, --verbose int                 set log level, use 0 to silence, 4 for debugging (default 3)
```

As shown above, you can use the following four flags to use TLS:

Flags | Description
---|---
`--admin-service-url` | The admin web service URL that pulsarctl connects to
`--auth-params` | Authentication parameters are used to configure the public and private key files required by TLS
`--tls-allow-insecur` | Allow TLS insecure connection
`--tls-trust-cert-path` | Allow TLS trust cert file path

**Example**

```
pulsarctl \
    --admin-service-url https://localhost:8443 \
    --auth-params "{\"tlsCertFile\":\"/test/auth/certs/client-cert.pem\",\"tlsKeyFile\":\"/test/auth/certs/client-key.pem\"}" \
	--tls-allow-insecure \
	--tls-trust-cert-path /test/auth/certs/cacert.pem \
	topics list public/default
```

For more information about TLS, see [Transport Encryption Using TLS](https://pulsar.apache.org/docs/en/security-tls-transport/).

### How to enable JWT

You need to configure a broker or a client.

#### Config broker

To enable JWT running in a broker, you need to set `authenticationEnabled` to `true` in the `broker.conf` or `standalone.conf` file. 

**Example**

```
# Enable JWT

# Enable authentication
authenticationEnabled=true

# Autentication provider name list, which is comma separated list of class names
authenticationProviders=org.apache.pulsar.broker.authentication.AuthenticationProviderToken

# Configure the secret key to be used to validate auth tokens
tokenSecretKey=file:///test/auth/token/secret.key
```

#### Configure a client

After running `pulsarctl help`, you can get the following information:

```
Common flags:
  -s, --admin-service-url string    The admin web service url that pulsarctl connects to. (default "http://localhost:8080")
      --auth-params string          Authentication parameters are used to configure the public and private key files required by tls
                                     For example: "tlsCertFile:val1,tlsKeyFile:val2"
  -C, --color string                toggle colorized logs (true,false,fabulous) (default "true")
  -h, --help                        help for this command
      --tls-allow-insecure          Allow TLS insecure connection
      --tls-trust-cert-path string   Allow TLS trust cert file path
      --token string                Using the token to authentication
      --token-file string           Using the token file to authentication
  -v, --verbose int                 set log level, use 0 to silence, 4 for debugging (default 3)
``` 

As shown above, you can use the following two flags to use JWT, and you only need to specify one of them:

Flags | Description
---|---
`--token string` | String value
`--token-file string` | Path of token file

**Example**

* Use `--token string`

    ```
    pulsarctl \
        --token "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXVzZXIifQ.Yb52IE0B5wzooAdSlIlskEgb6_HBXST8k3lINZS5wwg" \
        topics list public/default
    ```

* Use `--token-file string`

    ```
    pulsarctl \
        --token-file file:///test/auth/tokens/token
        topics list public/default
    ```

For more information about JWT, see the [Token Authentication Admin](https://pulsar.apache.org/docs/en/security-token-admin/).

## Admin API

To perform operations against Pulsar, you can use admin API as well.

**Example**

Here takes Go admin API as an example.

If you want to get a list of standalone cluster (the command `pulsarctl clusters list standalone` provides the same result as well) , you can definite and implement an Go interface as below.

  
*  Interface definition

    ```
    // List returns the list of clusters
    List() ([]string, error)
    ```

* Code example

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

For more information about admin API, see [Admin API](https://godoc.org/github.com/streamnative/pulsarctl).
