ARG PULSAR_VERSION
FROM snstage/pulsar-all:$PULSAR_VERSION

# use root user
USER root

# install required packages
RUN apk update && apk add curl git build-base

# install golang
ENV GOLANG_VERSION 1.21.0
RUN curl -sSL https://storage.googleapis.com/golang/go$GOLANG_VERSION.linux-amd64.tar.gz \
		| tar -C /usr/local -xz
ENV PATH /usr/local/go/bin:$PATH

# copy the code into image
COPY . /pulsarctl

# Create file for active status.html endpoint
RUN mkdir -p /usr/local/apache && touch /usr/local/apache/htdocs

ENV PULSAR_HOME /pulsar
ENV PULSARCTL_HOME /pulsarctl

ENTRYPOINT /pulsarctl/scripts/entrypoint.sh
