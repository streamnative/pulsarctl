#!/usr/bin/env bash
set -e

readonly PULSAR_HOME=${PULSAR_HOME:-"/pulsar"}

echo "--- Run pulsar service at the directory ${PULSAR_HOME} ---"
pushd ${PULSAR_HOME}
bin/apply-config-from-env.py conf/standalone.conf
bin/pulsar-daemon start standalone -nss -nfw
until curl http://localhost:8080/admin/v2/tenants > /dev/null 2>&1
do
    sleep 1
    echo "Wait for pulsar service to be ready...$(date +%H:%M:%S)"
done
echo "--- Pulsar service is ready ---"
popd