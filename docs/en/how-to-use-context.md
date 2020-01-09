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

# How to use pulsarctl context

In multi-cluster, **context** is a very interesting and useful function. It can help users cache information of multiple clusters, and can switch between multiple clusters.

The feature will be support in v0.3.0 of pulsarctl.

## Define contexts and auth info

Suppose you have two clusters, one for development work and one for scratch work. In the `development` cluster, the broker service URL is `http://1.2.3.4:8080`, the bookie service URL is `http://1.2.3.4:8083`. In your `scratch` cluster, the broker service URL is `http://5.6.7.8:8080`, the bookie service URL is `http://5.6.7.8:8083`.

Now, you can use `set` of context to define your cluster, the command as follows:

```bash
$ pulsarctl context set development --admin-service-url="http://1.2.3.4:8080" --bookie-service-url="http://1.2.3.4:8083"
$ pulsarctl context set scratch --admin-service-url="http://5.6.7.8:8080" --bookie-service-url="http://5.6.7.8:8083"
```

Then you can use the following command check content of `pulsar`:

```bash
$ cat $HOME/.config/pulsar
```

The content of pulsar as follows:

```text
auth-info:
  development:
    locationoforigin: $HOME/.config/pulsar
    tls_trust_certs_file_path: ""
    tls_allow_insecure_connection: false
    token: ""
    tokenFile: ""
  scratch:
    locationoforigin: $HOME/.config/pulsar
    tls_trust_certs_file_path: ""
    tls_allow_insecure_connection: false
    token: ""
    tokenFile: ""
contexts:
  development:
    admin-service-url: http://1.2.3.4:8080
    bookie-service-url: http://1.2.3.4:8083
  scratch:
    admin-service-url: http://5.6.7.8:8080
    bookie-service-url: http://5.6.7.8:8083
current-context: ""
```

If you want to `set` the auth info, you can specify the following flags:

- `--token`
- `--token-file`
- `--tls-trust-cert-path`
- `--tls-allow-insecure`

## Set the current context

When you define the info of cluster in `$HOME/.config/pulsar`, you can quickly switch between clusters by using the following command(suppose you want to use scratch cluster):

```bash
$ pulsarctl context use scratch
```

Now, you are using the context and auth info of `scratch`. And you can validate the current context value by using `pulsarctl context current`.

If you don't know the current names of cluster, you can use the following command:

```bash
$ pulsarctl context get
```

The output as follows:

```text
+---------+--------+--------------------------+-----------------------+
| CURRENT |  NAME  |    BROKER SERVICE URL    |  BOOKIE SERVICE URL   |
+---------+--------+--------------------------+-----------------------+
| *       | test-2 | http://159.65.2.188:8080 | http://localhost:8080 |
|         | test-1 | http://159.65.9.22:8080  | http://localhost:8080 |
+---------+--------+--------------------------+-----------------------+
```

## Rename the context

For some reason, we defined the name of the current context incorrectly. If you want to modify the name, you can use the following command:

```bash
$ pulsarctl context rename old_name new_name
```

## Delete the context

If the current cluster information is invalid, you want to delete it(suppose the cluster name is `scratch`), you can use the following command:

```bash
$ pulsarctl context delete scratch
```


