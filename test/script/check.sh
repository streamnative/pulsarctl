#!/usr/bin/env bash

# Waiting bookie HTTP service start up, if the bookie HTTP service does not
# start up in 30 seconds, that means the bookie HTTP service is not start
# up successfully.
function checkBookieHTTP() {
    failed=0
    until curl localhost:8080; do
        echo waiting service start...
        failed=`expr ${failed} + 1`
        if [[ ${failed} == 30 ]]; then
            echo service start up was failed
            exit 1
        fi
        sleep 1
    done
}

function checkFunctionWorker() {
    failed=0
    until curl localhost:8080/admin/v2/persistent/public/functions/coordinate/stats; do
        echo "waiting function worker service start..."
        failed=`expr ${failed} + 1`
        if [[ ${failed} == 30 ]]; then
            echo "function worker service start up was failed"
            exit 1
        fi
        sleep 1
    done
}

case $1 in
    bookieHTTP) checkBookieHTTP
    ;;
    functionWorker) checkFunctionWorker
    ;;
    *) echo Which service you would like to check?
       echo Available service: bookieHTTP, functionWorker
    ;;
esac