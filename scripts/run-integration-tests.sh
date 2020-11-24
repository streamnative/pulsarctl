#!/usr/bin/env bash
set -e

readonly PROJECT_ROOT=`cd $(dirname $0)/..; pwd`
readonly IMAGE_NAME=pulsarctl-test

docker build --build-arg PULSAR_VERSION=latest \
             -t ${IMAGE_NAME} \
             -f ${PROJECT_ROOT}/scripts/test-docker/Dockerfile ${PROJECT_ROOT}

docker run ${IMAGE_NAME}
