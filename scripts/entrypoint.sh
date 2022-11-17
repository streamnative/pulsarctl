#!/usr/bin/env bash

set -e

readonly PULSARCTL_HOME=${PULSARCTL_HOME:-"/pulsarctl"}
readonly TEST_ARGS=${TEST_ARGS:-""}

function checkFunctionWorker() {
    failed=0
    until curl --silent localhost:8080/admin/v2/persistent/public/functions/coordinate/stats; do
        echo "waiting function worker service start..."
        failed=`expr ${failed} + 1`
        if [[ ${failed} == 30 ]]; then
            echo "function worker service start up was failed"
            exit 1
        fi
        sleep 5
    done
    sleep 30
}

pushd ${PULSARCTL_HOME}
# startup pulsar service
scripts/pulsar-service-startup.sh

# disable cgo
export CGO_ENABLED=0
# run tests
case ${TEST_ARGS} in
    token)
        echo "running token tests"
        CGO_ENABLED=1 go test -v -race ./pkg/auth/token.go ./pkg/auth/token_test.go
        ;;
    tls)
        echo "running tls tests"
        CGO_ENABLED=1 go test -v -race ./pkg/ctl/cluster -run TestTLS
        ;;
    function)
        echo "running function tests"
        checkFunctionWorker
        cp /pulsar/examples/api-examples.jar test/functions/
        cp /pulsar/examples/python-examples/logging_function.py test/functions/
        CGO_ENABLED=1 go test -v -race -test.timeout 10m $(go list ./... | grep function)
        ;;
    sink)
        echo "running sink tests"
        checkFunctionWorker
        mkdir -p test/sinks
        cp /pulsar/connectors/pulsar-io-data-generator-*.nar test/sinks/data-generator.nar
        CGO_ENABLED=1 go test -v -race -test.timeout 10m ./pkg/ctl/sinks
        ;;
    source)
        echo "running source tests"
        checkFunctionWorker
        mkdir -p test/sources
        cp /pulsar/connectors/pulsar-io-data-generator-*.nar test/sources/data-generator.nar
        CGO_ENABLED=1 go test -v -race -test.timeout 10m ./pkg/ctl/sources
        ;;
    packages)
        echo "running packages tests"
        checkFunctionWorker
        cp /pulsar/examples/api-examples.jar test/functions/
        cp /pulsar/examples/python-examples/logging_function.py test/functions/
        CGO_ENABLED=1 go test -v -race -test.timeout 10m $(go list ./... | grep packages)
        ;;
    *)
        echo "running normal unit tests"
        CGO_ENABLED=1 go test -v -race $(go list ./... | grep -v bookkeeper | grep -v bkctl | grep -v functions | grep -v sources | grep -v sinks | grep -v packages | grep -v test)
        ;;
esac
# stop pulsar service
scripts/pulsar-service-shutdown.sh
popd
