MINOR_VERSION=1
VERSION=$(shell cat VERSION)
export GO_VERSION=$(shell go version | awk '{print $$3}')

clean-site:
	rm -rf site/gen-pulsarctldocs/generators/pulsarctl-site-${VERSION}.tar.gz
	rm -rf site/gen-pulsarctldocs/generators/includes
	rm -rf site/gen-pulsarctldocs/generators/build
	rm -rf site/gen-pulsarctldocs/generators/manifest.json

build-site: clean-site
	go run site/gen-pulsarctldocs/main.go --pulsar-version v1_$(MINOR_VERSION)
	docker run -v ${PWD}/site/gen-pulsarctldocs/generators/includes:/source -v ${PWD}/site/gen-pulsarctldocs/generators/build:/build -v ${PWD}/site/gen-pulsarctldocs/generators/:/manifest pwittrock/brodocs
	mkdir -p dist
	tar -czvf site/gen-pulsarctldocs/generators/pulsarctl-site-${VERSION}.tar.gz -C site/gen-pulsarctldocs/generators/build/ .
	mv site/gen-pulsarctldocs/generators/pulsarctl-site-${VERSION}.tar.gz dist/pulsarctl-site-${VERSION}.tar.gz

build:
	CGO_ENABLED=0 go build -o bin/pulsarctl

goreleaser-release-snapshot:
	goreleaser release --rm-dist --snapshot

goreleaser-release:
	goreleaser release --rm-dist

install-goreleaser:
	go install github.com/goreleaser/goreleaser@v1.4.1
	goreleaser --version

.PHONY: install
install:
	go install github.com/streamnative/pulsarctl
