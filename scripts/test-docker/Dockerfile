ARG PULSAR_VERSION
FROM snstage/pulsar-all:$PULSAR_VERSION

# use root user
USER root

# install required packages
RUN apk update && apk add curl git build-base

# install golang
COPY --from=golang:1.24.0-alpine /usr/local/go/ /usr/local/go/
ENV PATH /usr/local/go/bin:$PATH


# copy the code into image
COPY . /pulsarctl

# Create file for active status.html endpoint
RUN mkdir -p /usr/local/apache && touch /usr/local/apache/htdocs

ENV PULSAR_HOME /pulsar
ENV PULSARCTL_HOME /pulsarctl

ENTRYPOINT /pulsarctl/scripts/entrypoint.sh
