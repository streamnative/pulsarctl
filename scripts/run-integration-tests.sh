#!/usr/bin/env bash
set -e

readonly PROJECT_ROOT=`cd $(dirname $0)/..; pwd`
readonly IMAGE_NAME=pulsarctl-test
readonly PULSAR_DEFAULT_VERSION="latest"
readonly PULSAR_VERSION=${PULSAR_VERSION:-${PULSAR_DEFAULT_VERSION}}
# Sink/source tests copy connector files from this image into the broker image.
readonly PULSAR_IO_IMAGE=${PULSAR_IO_IMAGE:-"apachepulsar/pulsar-all"}
readonly PULSAR_IO_IMAGE_VERSION=${PULSAR_IO_IMAGE_VERSION:-${PULSAR_VERSION}}
readonly PULSAR_IO_CONNECTORS_DIR=${PULSAR_IO_CONNECTORS_DIR:-"/pulsar/connectors"}
readonly PULSAR_IMAGE=${PULSAR_IMAGE:-"apachepulsar/pulsar"}

build_target=tests
case ${1} in
    sink|source)
        build_target=tests-with-connectors
        ;;
esac

docker build --target ${build_target} \
             --build-arg PULSAR_VERSION=${PULSAR_VERSION} \
             --build-arg PULSAR_IMAGE=${PULSAR_IMAGE} \
             --build-arg PULSAR_IO_IMAGE=${PULSAR_IO_IMAGE} \
             --build-arg PULSAR_IO_IMAGE_VERSION=${PULSAR_IO_IMAGE_VERSION} \
             --build-arg PULSAR_IO_CONNECTORS_DIR=${PULSAR_IO_CONNECTORS_DIR} \
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
        docker run --name function --rm -e TEST_ARGS=function -e PULSAR_STANDALONE_USE_ZOOKEEPER=true -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    sink)
        docker run --name sink --rm -e TEST_ARGS=sink -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    source)
        docker run --name sink --rm -e TEST_ARGS=source -e FUNCTION_ENABLE=true ${IMAGE_NAME}
        ;;
    packages)
        docker run --name packages --rm -e TEST_ARGS=packages -e PULSAR_STANDALONE_USE_ZOOKEEPER=true -e PULSAR_PREFIX_enablePackagesManagement=true -e PULSAR_PREFIX_zookeeperServers=127.0.0.1:2181 ${IMAGE_NAME}
        ;;
    *)
        env_file=${PROJECT_ROOT}/test/policies/policies.env
        docker run --env-file ${env_file} ${IMAGE_NAME}
        ;;
esac
