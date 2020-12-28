#!/usr/bin/env bash
set -e

readonly PROJECT_ROOT=`cd $(dirname $0)/..; pwd`
readonly IMAGE_NAME=pulsarctl-test
readonly PULSAR_DEFAULT_VERSION="2.7.0.1"
readonly PULSAR_VERSION=${PULSAR_VERSION:-${PULSAR_DEFAULT_VERSION}}

docker build --build-arg PULSAR_VERSION=${PULSAR_VERSION} \
             -t ${IMAGE_NAME} \
             -f ${PROJECT_ROOT}/scripts/test-docker/Dockerfile ${PROJECT_ROOT}
case ${1} in
    token)
        env_file=${PROJECT_ROOT}/test/auth/token.env
        docker run --env-file ${env_file} -e TEST_ARGS=token ${IMAGE_NAME}
        ;;
    tls)
        env_file=${PROJECT_ROOT}/test/auth/tls.env
        docker run --env-file ${env_file} -e TEST_ARGS=tls ${IMAGE_NAME}
        ;;
    *)
        docker run ${IMAGE_NAME}
        ;;
esac
