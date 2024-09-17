ARG VERSION=0.5.0

FROM golang:1.18-alpine as builder

ARG VERSION
ARG TARGETOS
ARG TARGETARCH


WORKDIR /workspace

COPY ./ .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o pulsarctl -ldflags "-X github.com/streamnative/pulsarctl/pkg/pulsar.ReleaseVersion=Pulsarctl-Go-${VERSION}" .

FROM alpine:3.16.0

RUN apk add --no-cache \
  tzdata=2022c-r0 \
  ca-certificates=20220614-r0 \
  bash=5.1.16-r2 \
  openssl=1.1.1q-r0

COPY ./plugins /root/.pulsarctl/plugins

ENV PATH=${PATH}:/root/.pulsarctl/plugins

COPY --from=builder /workspace/pulsarctl /usr/local/bin/pulsarctl
RUN chmod +x /usr/local/bin/pulsarctl
