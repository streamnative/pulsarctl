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

version=`curl -s https://raw.githubusercontent.com/streamnative/pulsarctl/master/stable.txt`

discoverArch() {
  ARCH=$(uname -m)
  case $ARCH in
    x86) ARCH="386";;
    x86_64) ARCH="amd64";;
    i686) ARCH="386";;
    i386) ARCH="386";;
  esac
}

# discovers the operating system for this system.
discoverOS() {
  OS=$(echo `uname`|tr '[:upper:]' '[:lower:]')
  SUFFIX=''

  case "$OS" in
    # Minimalist GNU for Windows
    mingw*)
        OS='windows'
        SUFFIX='.exe'
        ;;
  esac
}

discoverArch
discoverOS

curl -# -LO https://github.com/streamnative/pulsarctl/releases/download/$version/pulsarctl-${ARCH}-${OS}${SUFFIX}

if [[ ${OS} != windows ]]; then
    mv pulsarctl-${ARCH}-${OS} pulsarctl
    chmod +x pulsarctl
    mv pulsarctl /usr/local/bin
fi
