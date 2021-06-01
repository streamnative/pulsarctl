ARG PULSAR_VERSION
FROM streamnative/pulsar-all:$PULSAR_VERSION

# install golang
ENV GOLANG_VERSION 1.13.3
RUN curl -sSL https://storage.googleapis.com/golang/go$GOLANG_VERSION.linux-amd64.tar.gz \
		| tar -v -C /usr/local -xz
ENV PATH /usr/local/go/bin:$PATH

# copy the code into image
COPY . /pulsarctl
RUN cp /pulsarctl/test/standalone.conf /pulsar/conf/standalone.conf
ENV PULSAR_HOME /pulsar
ENV PULSARCTL_HOME /pulsarctl

ENTRYPOINT /pulsarctl/scripts/entrypoint.sh
