#!/usr/bin/env bash

set -ex

echo "Releasing pulsarctl"

version=${1#v}
if [[ "x$version" == "x" ]]; then
  echo "You need give a version number of the pulsarctl"
  exit 1
fi

# Create a direcotry to save assets
ASSETSDIR=release
mkdir $ASSETSDIR

function build_amd64_linux() {
  DIR=pulsarctl-amd64-linux
  mkdir ${DIR}
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pulsarctl -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-$version" .
  mv pulsarctl ${DIR}
  cp -r plugins ${DIR}
  tar -czf ${DIR}.tar.gz ${DIR}
  mv ${DIR}.tar.gz $ASSETSDIR
}

function build_386_linux() {
  DIR=pulsarctl-386-linux
  mkdir ${DIR}
  CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o pulsarctl -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-$version" .
  mv pulsarctl ${DIR}
  cp -r plugins ${DIR}
  tar -czf ${DIR}.tar.gz ${DIR}
  mv ${DIR}.tar.gz $ASSETSDIR
}

function build_amd64_darwin() {
  DIR=pulsarctl-amd64-darwin
  mkdir ${DIR}
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o pulsarctl -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-$version" .
  mv pulsarctl ${DIR}
  cp -r plugins ${DIR}
  tar -czf ${DIR}.tar.gz ${DIR}
  mv ${DIR}.tar.gz $ASSETSDIR
}

function build_arm64_darwin() {
  DIR=pulsarctl-arm64-darwin
  mkdir ${DIR}
  CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o pulsarctl -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-$version" .
  mv pulsarctl ${DIR}
  cp -r plugins ${DIR}
  tar -czf ${DIR}.tar.gz ${DIR}
  mv ${DIR}.tar.gz $ASSETSDIR
}

function build_386_darwin() {
  DIR=pulsarctl-386-darwin
  mkdir ${DIR}
  CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o pulsarctl -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-$version" .
  mv pulsarctl ${DIR}
  cp -r plugins ${DIR}
  tar -czf ${DIR}.tar.gz ${DIR}
  mv ${DIR}.tar.gz $ASSETSDIR
}

function build_amd64_windows() {
  DIR=pulsarctl-amd64-windows
  mkdir ${DIR}
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o  pulsarctl.exe -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-$version" .
  mv pulsarctl.exe ${DIR}
  cp -r plugins ${DIR}
  tar -czf ${DIR}.tar.gz ${DIR}
  mv ${DIR}.tar.gz $ASSETSDIR
}

function build_386_windows() {
  DIR=pulsarctl-386-windows
  mkdir ${DIR}
  CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o  pulsarctl.exe -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-$version" .
  mv pulsarctl.exe ${DIR}
  cp -r plugins ${DIR}
  tar -czf ${DIR}.tar.gz ${DIR}
  mv ${DIR}.tar.gz $ASSETSDIR
}

function build_doc() {
  echo ${version} > VERSION
  make cli
  mv pulsarctl-site-${version}.tar.gz ${ASSETSDIR}
}

build_amd64_linux
build_386_linux
build_amd64_darwin
build_arm64_darwin
build_amd64_windows
build_386_windows
build_doc
