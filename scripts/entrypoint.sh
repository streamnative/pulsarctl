#!/usr/bin/env bash

set -e

readonly PULSARCTL_HOME=${PULSARCTL_HOME:-"/pulsarctl"}
pushd ${PULSARCTL_HOME}
# startup pulsar service
scripts/pulsar-service-startup.sh
# run tests
go test -v $(go list ./... | grep -v bookkeeper | grep -v bkctl | grep -v functions | grep -v sources | grep -v sinks | grep -v test)
# stop pulsar service
scripts/pulsar-service-shutdown.sh
popd
