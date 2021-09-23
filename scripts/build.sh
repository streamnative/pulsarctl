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

LDFLAGS+="-X \"github.com/streamnative/pulsarctl/pkg/cmdutils.ReleaseVersion=${version}\" "
LDFLAGS+="-X \"github.com/streamnative/pulsarctl/pkg/cmdutils.BuildTS=$(date -u '+%Y-%m-%d %H:%M:%S')\" "
LDFLAGS+="-X \"github.com/streamnative/pulsarctl/pkg/cmdutils.GitHash=$(git rev-parse HEAD)\" "
LDFLAGS+="-X \"github.com/streamnative/pulsarctl/pkg/cmdutils.GitBranch=$(git rev-parse --abbrev-ref HEAD)\" "
LDFLAGS+="-X \"github.com/streamnative/pulsarctl/pkg/cmdutils.GoVersion=$(go version)\" "

echo $LDFLAGS
# Create a direcotry to save assets
ASSETS_DIR=${ROOT_DIR}/release
mkdir $ASSETS_DIR

build() {
    local arch=${1}
    local os=${2}
    local base_dir=dist
    local dirname=pulsarctl-${arch}-${os}
    local dir=${base_dir}/${dirname}
    mkdir -p ${dir}
    CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build \
        -ldflags "${LDFLAGS}" -o pulsarctl
    mv pulsarctl ${dir}
    cp -r plugins ${dir}
    pushd $base_dir
    tar -czf ${dirname}.tar.gz ${dirname}
    mv ${dirname}.tar.gz ${ASSETS_DIR}
    popd
}

function build_doc() {
  echo ${version} > VERSION
  make cli
  mv pulsarctl-site-${version}.tar.gz ${ASSETSDIR}
}

build amd64 linux
build 386 linux
build amd64 darwin
build arm64 darwin
build amd64 windows
build 386 windows
build_doc

popd
