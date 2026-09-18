#!/usr/bin/env bash
set -e

readonly PROJECT_ROOT=`cd $(dirname $0)/..; pwd`
readonly IMAGE_NAME=pulsarctl-test
readonly PULSAR_DEFAULT_VERSION="5.0.0-SNAPSHOT"
readonly PULSAR_VERSION=${PULSAR_VERSION:-${PULSAR_DEFAULT_VERSION}}
readonly PULSAR_IMAGE="docker-proxy.streamnative.io/snstage/pulsar-cloud"

docker build --build-arg PULSAR_VERSION=${PULSAR_VERSION} \
             --build-arg PULSAR_IMAGE=${PULSAR_IMAGE} \
             -t ${IMAGE_NAME} \
             -f ${PROJECT_ROOT}/scripts/test-docker/Dockerfile ${PROJECT_ROOT}
case ${1} in
    token)
        env_file=${PROJECT_ROOT}/test/auth/token.env
        docker run -e PULSAR_PREFIX_defaultNumNamespaceBundles=4 --rm --env-file ${env_file} -e TEST_ARGS=token ${IMAGE_NAME}
        ;;
    tls)
        env_file=${PROJECT_ROOT}/test/auth/tls.env
        docker run -e PULSAR_PREFIX_defaultNumNamespaceBundles=4 --rm --env-file ${env_file} -e TEST_ARGS=tls ${IMAGE_NAME}
        ;;
    function)
        docker run -e PULSAR_PREFIX_defaultNumNamespaceBundles=4 --name function --rm -e TEST_ARGS=function -e PULSAR_STANDALONE_USE_ZOOKEEPER=true -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    sink)
        docker run -e PULSAR_PREFIX_defaultNumNamespaceBundles=4 --name sink --rm -e TEST_ARGS=sink -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    source)
        docker run -e PULSAR_PREFIX_defaultNumNamespaceBundles=4 --name sink --rm -e TEST_ARGS=source -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    packages)
        docker run -e PULSAR_PREFIX_defaultNumNamespaceBundles=4 --name packages --rm -e TEST_ARGS=packages -e PULSAR_STANDALONE_USE_ZOOKEEPER=true -e PULSAR_PREFIX_enablePackagesManagement=true -e PULSAR_PREFIX_zookeeperServers=127.0.0.1:2181 ${IMAGE_NAME}
        ;;
    *)
        env_file=${PROJECT_ROOT}/test/policies/policies.env
        docker run -e PULSAR_PREFIX_defaultNumNamespaceBundles=4 --env-file ${env_file} ${IMAGE_NAME}
        ;;
esac

