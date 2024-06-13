#!/usr/bin/env bash
set -e

readonly PULSAR_HOME=${PULSAR_HOME:-"/pulsar"}

echo "--- Run pulsar service at the directory ${PULSAR_HOME} ---"
pushd ${PULSAR_HOME}
bin/apply-config-from-env.py conf/standalone.conf
if [ ${FUNCTION_ENABLE} ]; then
    bin/pulsar-daemon start standalone
else
    bin/pulsar-daemon start standalone -nss -nfw
fi
until curl http://127.0.0.1:8080/admin/v2/tenants > /dev/null
do
    sleep 1
    echo "Wait for pulsar service to be ready...$(date +%H:%M:%S)"
done
echo "--- Pulsar service is ready ---"
popd