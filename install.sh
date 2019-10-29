#!/usr/bin/env bash
#
#/**
# * Licensed to the Apache Software Foundation (ASF) under one
# * or more contributor license agreements.  See the NOTICE file
# * distributed with this work for additional information
# * regarding copyright ownership.  The ASF licenses this file
# * to you under the Apache License, Version 2.0 (the
# * "License"); you may not use this file except in compliance
# * with the License.  You may obtain a copy of the License at
# *
# *     http://www.apache.org/licenses/LICENSE-2.0
# *
# * Unless required by applicable law or agreed to in writing, software
# * distributed under the License is distributed on an "AS IS" BASIS,
# * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# * See the License for the specific language governing permissions and
# * limitations under the License.
# */

set -e

version=v0.1.0

if [[ $(uname) == Darwin ]]; then
    curl -# -LO https://github.com/streamnative/pulsarctl/releases/download/$version/pulsarctl
    chmod +x pulsarctl
    mv pulsarctl /usr/local/bin
elif [[ $(expr substr $(uname -s) 1 5) == Linux ]]; then
    curl -# -LO https://github.com/streamnative/pulsarctl/releases/download/$version/pulsarctl-linux
    chmod +x pulsarctl-linux
    mv pulsarctl-linux /usr/local/bin/pulsarctl
fi

