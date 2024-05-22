FROM alpine:3.19.1

ADD pulsarctl /usr/local/bin/pulsarctl

RUN apk update \
    && apk upgrade --no-cache \
    && apk add tzdata ca-certificates --no-cache \
    && chmod +x /usr/local/bin/pulsarctl
