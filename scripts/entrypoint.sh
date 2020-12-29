#!/usr/bin/env bash

set -e

readonly PULSARCTL_HOME=${PULSARCTL_HOME:-"/pulsarctl"}
readonly TEST_ARGS=${TEST_ARGS:-""}

pushd ${PULSARCTL_HOME}
# startup pulsar service
scripts/pulsar-service-startup.sh
# run tests
case ${TEST_ARGS} in
    token)
        echo "running token tests"
        go test -v ./pkg/auth/token.go ./pkg/auth/token_test.go
        ;;
    tls)
        echo "running tls tests"
        go test -v ./pkg/ctl/cluster -run TestTLS
        ;;
    *)
        echo "running normal unit tests"
        go test -v $(go list ./... | grep -v bookkeeper | grep -v bkctl | grep -v functions | grep -v sources | grep -v sinks | grep -v test)
        ;;
esac
# stop pulsar service
scripts/pulsar-service-shutdown.sh
popd
