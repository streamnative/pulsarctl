FROM alpine:3.20

ADD pulsarctl /usr/local/bin/pulsarctl

RUN apk add tzdata ca-certificates --no-cache \
    && chmod +x /usr/local/bin/pulsarctl
