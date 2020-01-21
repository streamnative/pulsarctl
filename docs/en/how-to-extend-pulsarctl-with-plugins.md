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

# Extent pulsarctl with plugins

This guide demonstrates how to install and write extensions for pulsarctl.

## Before you begin

You need to have a working `pulsarctl` binary installed. And the version of the pulsarctl need to be 0.4.0+.

## Install pulsarctl plugins

A plugin is nothing more than a standalone executable file, whose name begins with `pulsarctl-`. To install
a plugin, simply move its executable file to anywhere on your `PATH`.

## Writing pulsarctl plugins

You can write a plugin in any programming language or script that allows you to write command-line commands.

There is no plugin installation or pre-loading required. Plugin executables receive the inherited environment 
from the plsarctl binary. A plugin determines which command path it wishes to implement based on its name. For 
example, a plugin wanting to provide a new command pulsarctl foo, would simply be named pulsarctl-foo, and live 
somewhere in your PATH.

### Example plugin

```bash
#!/bin/bash

if [[ $1 == "args" ]]
then 
    echo "I am the args of the pulsarctl-foo"
    exit 0
fi

echo "I am a plugin named pulsarctl-foo"
```

### Use a plugin

To use the above plugin, simply make it executable:

```bash
chmod +x ./pulsarctl-foo
```

and place it anywhere in your `PATH`:

```bash
mv pulsarctl-foo /usr/local/bin
```

You may now invoke your plugin as a kubectl command:

```bash
pulsarctl foo
```

```
I am a plugin named pulsarctl-foo
```

All args and flags are passed as-is to the executable:

```bash
pulsarctl foo args
```

```
I am the args of the pulsarctl-foo
```
