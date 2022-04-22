#!/usr/bin/env bash
set -e

readonly PROJECT_ROOT=`cd $(dirname $0)/..; pwd`
readonly IMAGE_NAME=pulsarctl-test
readonly PULSAR_DEFAULT_VERSION="2.10.0.2"
readonly PULSAR_VERSION=${PULSAR_VERSION:-${PULSAR_DEFAULT_VERSION}}

docker build --build-arg PULSAR_VERSION=${PULSAR_VERSION} \
             -t ${IMAGE_NAME} \
             -f ${PROJECT_ROOT}/scripts/test-docker/Dockerfile ${PROJECT_ROOT}
case ${1} in
    token)
        env_file=${PROJECT_ROOT}/test/auth/token.env
        docker run --rm --env-file ${env_file} -e TEST_ARGS=token ${IMAGE_NAME}
        ;;
    tls)
        env_file=${PROJECT_ROOT}/test/auth/tls.env
        docker run --rm --env-file ${env_file} -e TEST_ARGS=tls ${IMAGE_NAME}
        ;;
    function)
        docker run --name function --rm -e TEST_ARGS=function -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    sink)
        docker run --name sink --rm -e TEST_ARGS=sink -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    source)
        docker run --name sink --rm -e TEST_ARGS=source -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    packages)
        docker run --name packages --rm -e TEST_ARGS=packages -e PULSAR_PREFIX_enablePackagesManagement=true -e PULSAR_PREFIX_zookeeperServers=127.0.0.1:2181 ${IMAGE_NAME}
        ;;
    *)
        env_file=${PROJECT_ROOT}/test/policies/policies.env
        docker run --env-file ${env_file} ${IMAGE_NAME}
        ;;
esac
