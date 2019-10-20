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

## Requirement

- Go 1.11 +

## Usage

### Prerequisite

If you have not installed Go, install it according to the [installation instruction](http://golang.org/doc/install).

Since the `go mod` package management tool is used in this project, **Go 1.11 or higher** version is required.

### Install

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

- pulsar-admin broker-stats subcommand
- pulsar-admin brokers subcommand
- pulsar-admin ns-isolation-policy subcommand
- pulsar-admin resource-quotas subcommand
- pulsar-admin functions localrun options
- pulsar-admin sources localrun options
- pulsar-admin sinks localrun options
- pulsar-admin schemas extract options

## Contribute

Contributions are welcomed and greatly appreciated. 
For more information about how to submit a patch and the contribution workflow, see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Licensed under the Apache License Version 2.0: http://www.apache.org/licenses/LICENSE-2.0
