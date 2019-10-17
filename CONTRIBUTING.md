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

# How to contribute

If you would like to contribute code to this project, fork the repository and send a pull request.

## Prerequisite

In this project, use `go mod` as the package management tool and make sure your Go version is higher then `Go 1.11`.

## Fork

Before contributing, you need to fork [pulsarctl](https://github.com/streamnative/pulsarcli) to your GitHub repository.

## Contribution flow

```bash
$ git remote add streamnative https://github.com/streamnative/pulsarctl.git
# sync with the remote master
$ git checkout master
$ git fetch streamnative
$ git rebase streamnative/master
$ git push origin master
# create a PR branch
$ git checkout -b your_branch   
# do something
$ git add [your change files]
$ git commit -sm "xxx"
$ git push origin your_branch
```

## Configure GoLand

The `pulsarctl` uses `go mod` to manage dependencies, so make sure your IDE enables `Go Modules(vgo)`.

To configure annotation processing in GoLand, follow the steps below.

1. To open the **Go Modules Settings** window, in GoLand, click **Preferences** > **Go** > **Go Modules(vgo)**.

2. Select the **Enable Go Modules(vgo) integration** checkbox.

3. Click **Apply** and **OK**.

## Code style

The code style suggested by the Golang community is used in `pulsarctl`. 
For more information, see [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).

To make your pull request easy to review, maintain and develop, follow this style.

## Create a new file

The project uses the open source protocol of Apache License 2.0. If you need to create a new file when developing new features, 
add the license at the beginning of each file. The location of the header file: [header file](.header).

## Update dependencies

The `pulsarctl` uses [Go 1.11 module](https://github.com/golang/go/wiki/Modules) to manage dependencies. To add or update a dependency, use the `go mod edit` command to change the dependency.
