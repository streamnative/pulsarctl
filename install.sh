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

if [[ "x${version}" == "x" ]]; then
    version=$(curl -s https://raw.githubusercontent.com/streamnative/pulsarctl/master/stable.txt)
fi

discoverArch() {
  ARCH=$(uname -m)
  case $ARCH in
    x86) ARCH="386";;
    x86_64) ARCH="amd64";;
    i686) ARCH="386";;
    i386) ARCH="386";;
  esac
}

usage() {
    cat <<EOF
This script is used to install pulsarctl
Options:
       -h,--help               Prints the usage message
       -u,--user               Install to the user install directory for your platform. Typically '$HOME/.pulsarctl/pulsarctl'.
       -v,--version            Install a specific version of pulsarctl
EOF
}

userOnly=false

while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -u|--user)
    userOnly=true
    shift
    ;;
    -v|--version)
    version="$2"
    shift
    shift
    ;;
    -h|--help)
    usage
    exit 0
    ;;
    *)
    echo "unknown option: $key"
    usage
    exit 1
    ;;
esac
done


discoverArch
OS=$(echo `uname`|tr '[:upper:]' '[:lower:]')

copyBinary() {
  target_dir=/usr/local/bin
  if [[ "${userOnly}" == "true" ]]; then
      target_dir=${HOME}/.pulsarctl
      mkdir -p ${target_dir}
  fi
    
  target_binary_file=pulsarctl${version}
  target_binary_file_path=${target_dir}/${target_binary_file}
  if [[ -f ${target_binary_file_path} ]]; then
    rm ${target_binary_file_path}
  fi
  mv pulsarctl ${target_binary_file_path}
  chmod +x ${target_binary_file_path}
  if [[ -f ${target_dir}/pulsarctl ]];then
    rm ${target_dir}/pulsarctl
  fi
  ln -s ${target_binary_file} ${target_dir}/pulsarctl
}

installNew() {
  TARFILE=pulsarctl-${ARCH}-${OS}.tar.gz
  UNTARFILE=pulsarctl-${ARCH}-${OS}

  if [[ -f ${TARFILE} ]]; then
    rm -f ${TARFILE} 
  fi

  if [[ -d ${UNTARFILE} ]]; then
    rm -rf ${UNTARFILE} 
  fi

  if [[ -f ${UNTARFILE} ]]; then
    rm -f ${UNTARFILE} 
  fi

  curl --retry 10 -L -o ${TARFILE} https://github.com/streamnative/pulsarctl/releases/download/${version}/${TARFILE}
  tar -xzf ${TARFILE}

  pushd ${UNTARFILE}

  copyBinary

  local plugins_dir=${HOME}/.pulsarctl/plugins
  
  mkdir -p ${plugins_dir}
  cp -r plugins/* ${plugins_dir}
  rm -rf ${TARFILE}
  rm -rf ${UNTARFILE}

  echo "The plugins of pulsarctl ${version} are successfully installed under directory '${plugins_dir}'."
  echo
  echo "In order to use this plugins, please add the plugin directory '${plugins_dir}' to the system PATH. You can do so by adding the following line to your bash profile."
  echo
  echo 'export PATH=${PATH}:${HOME}/.pulsarctl/plugins'
  echo
  echo "Happy Pulsaring!"

  export PATH=${HOME}/.pulsarctl:${HOME}/.pulsarctl/plugins:${PATH}
  popd
}

installOld() {
  binary=pulsarctl-${ARCH}-${OS}

  if [[ -d ${binary} ]]; then
    rm -rf ${binary} 
  fi

  if [[ -f ${binary} ]]; then
    rm -f ${binary} 
  fi

  curl --retry 10 -L -o ${binary} https://github.com/streamnative/pulsarctl/releases/download/${version}/${binary}
  mv ${binary} pulsarctl

  copyBinary

  echo "Happy Pulsaring!"
}

case $version in
  v0.1.0)
    installOld
    ;;
  v0.2.0)
    installOld
    ;;
  v0.3.0)
    installOld
    ;;
  *)
    installNew
esac
