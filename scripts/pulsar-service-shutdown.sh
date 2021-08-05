#!/usr/bin/env bash
set -e

readonly PULSAR_HOME=${PULSAR_HOME:-"/pulsar"}
pushd ${PULSAR_HOME}
echo "--- Stop the pulsar service ---"
bin/pulsar-daemon stop standalone -force
echo "--- Pulsar service is stopped ---"
popd