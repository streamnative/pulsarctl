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

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![LICENSE](https://img.shields.io/hexpm/l/pulsar.svg)](https://github.com/streamnative/pulsarctl/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/Godoc-reference-blue.svg)](https://godoc.org/github.com/streamnative/pulsarctl)

# Pulsarctl

A CLI tool for the [Apache Pulsar](https://pulsar.incubator.apache.org/) project.

## Usage

### Install

#### Install pulsarctl binary with curl on Linux

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/streamnative/pulsarctl/master/install.sh)"
```

#### Install pulsarctl binary with curl on macOS

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/streamnative/pulsarctl/master/install.sh)"
```

#### Install pulsarctl binary on Windows

If you have command-line shell tools, such as PowerShell. You can download pulsarctl by this way:

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/streamnative/pulsarctl/master/install.sh)"
```

Or you can download the executable directly.

``` bash
curl -LO https://github.com/streamnative/pulsarctl/releases/download/v0.1.0/pulsarctl.exe
```

To find out the latest stable version, take a look at https://raw.githubusercontent.com/streamnative/pulsarctl/master/stable.txt

## Install by code

### Prerequisite

- Go 1.11 +

If you have not installed Go, install it according to the [installation instruction](http://golang.org/doc/install).

Since the `go mod` package management tool is used in this project, **Go 1.11 or higher** version is required.

### Download Code

1. Clone the project from GitHub to your local.

```bash
git clone https://github.com/streamnative/pulsarctl.git
```

2. Use `go mod` to get the dependencies needed for the project.

```bash
go mod download
```

After entering the `go mod download` command, if some libs can not be downloaded, then you can download them by referring to the proxy provided by [GOPROXY.io](https://goproxy.io/).

### Build

```bash
export GO111MODULE=on

go build -o pulsarctl main.go
```

## Project Status

The following is an incomplete list of features that are not yet implemented:
 
#### Functions
- localrun

#### Sources
- localrun
- available-sources
- reload

#### Sinks
- localrun
- available-sources
- reload

#### Schemas
- extract

#### Namespaces
- delete-bookie-affinity-group
- get-bookie-affinity-group
- set-bookie-affinity-group 

#### Bookies
- racks-placement
- get-bookie-rack
- delete-bookie-rack
- set-bookie-rack

## Different With Java Pulsar Admin

We move the subscription commands from the Topics to the Subscriptions in pulsarctl.
 
| pulsar-admin | pulsarctl |
| ------------ | --------- |
| bin/pulsar-admin topics create-subscription | pulsarctl subscription create |
| bin/pulsar-admin topics unsubscribe | pulsarctl subscription delete |
| bin/pulsar-admin topics skip | pulsarctl subscription skip |
| bin/pulsar-admin topics expire-messages | pulsarctl subscription expire |
| bin/pulsar-admin topics peek-messages | pulsarctl subscription peek |
| bin/pulsar-admin topics reset-cursor | pulsarctl subscription seek |
| bin/pulsar-admin topics subscriptions | pulsarctl subscription list |

## Contribute

Contributions are welcomed and greatly appreciated. 
For more information about how to submit a patch and the contribution workflow, see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Licensed under the Apache License Version 2.0: http://www.apache.org/licenses/LICENSE-2.0
