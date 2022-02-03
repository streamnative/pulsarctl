#!/usr/bin/env bash

set -ex

echo "Releasing pulsarctl"

version=${1#v}
if [[ "x$version" == "x" ]]; then
  echo "You need give a version number of the pulsarctl"
  exit 1
fi

ROOT_DIR=`cd $(dirname $0)/../; pwd`
pushd $ROOT_DIR

echo "${version}" > VERSION

make goreleaser-release-snapshot
make build-doc

popd
