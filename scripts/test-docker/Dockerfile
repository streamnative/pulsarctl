ARG PULSAR_IMAGE=apachepulsar/pulsar
ARG PULSAR_VERSION=latest
ARG PULSAR_IO_IMAGE=apachepulsar/pulsar-all
ARG PULSAR_IO_IMAGE_VERSION=${PULSAR_VERSION}

FROM ${PULSAR_IO_IMAGE}:${PULSAR_IO_IMAGE_VERSION} AS connectors

FROM ${PULSAR_IMAGE}:${PULSAR_VERSION} AS tests

# use root user
USER root

# install required packages
RUN apk update && apk add curl git build-base

# install golang
COPY --from=golang:1.25.12-alpine /usr/local/go/ /usr/local/go/
ENV PATH /usr/local/go/bin:$PATH


# copy the code into image
COPY . /pulsarctl

# Create file for active status.html endpoint
RUN mkdir -p /usr/local/apache && touch /usr/local/apache/htdocs

ENV PULSAR_HOME /pulsar
ENV PULSARCTL_HOME /pulsarctl

ENTRYPOINT /pulsarctl/scripts/entrypoint.sh

FROM tests AS tests-with-connectors
ARG PULSAR_IO_CONNECTORS_DIR=/pulsar/connectors
COPY --from=connectors ${PULSAR_IO_CONNECTORS_DIR}/ /pulsar/connectors/
