#!/usr/bin/env bash

# Waiting bookie service start up, if the bookie service does not
# start up in 120 seconds, that means the bookie service is not start
# up successfully.
function checkBookie() {
    failed=0
    until curl localhost:8080; do
        echo waiting service start...
        failed=`expr ${failed} + 1`
        if [[ ${failed} == 120 ]]; then
            echo service start up was failed
            exit 1
        fi
        sleep 1
    done
}

case $1 in
    bookie) checkBookie
    ;;
    *) echo Which service you would like to check?
       echo Available service: bookie
    ;;
esac