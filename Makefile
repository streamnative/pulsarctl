MINOR_VERSION=1
VERSION=$(shell cat VERSION)

LDFLAGS += -X "github.com/streamnative/pulsarctl/pkg/cmdutils.ReleaseVersion=$(shell git describe --tags --always)"
LDFLAGS += -X "github.com/streamnative/pulsarctl/pkg/cmdutils.BuildTS=$(shell date -u '+%Y-%m-%d %H:%M:%S')"
LDFLAGS += -X "github.com/streamnative/pulsarctl/pkg/cmdutils.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "github.com/streamnative/pulsarctl/pkg/cmdutils.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "github.com/streamnative/pulsarctl/pkg/cmdutils.GoVersion=$(shell go version)"

GO := GO111MODULE=on go
GOBUILD := $(GO) build

# Build pulsarctl docs

cleancli:
	rm -f main
	rm -rf $(shell pwd)/site/gen-pulsarctldocs/generators/pulsarctl-site-${VERSION}.tar.gz
	rm -rf $(shell pwd)/site/gen-pulsarctldocs/generators/includes
	rm -rf $(shell pwd)/site/gen-pulsarctldocs/generators/build
	rm -rf $(shell pwd)/site/gen-pulsarctldocs/generators/manifest.json

cli: cleancli
	export GO111MODULE=on 
	go run site/gen-pulsarctldocs/main.go --pulsar-version v1_$(MINOR_VERSION)
	docker run -v ${PWD}/site/gen-pulsarctldocs/generators/includes:/source -v ${PWD}/site/gen-pulsarctldocs/generators/build:/build -v ${PWD}/site/gen-pulsarctldocs/generators/:/manifest pwittrock/brodocs
	tar -czvf ${PWD}/site/gen-pulsarctldocs/generators/pulsarctl-site-${VERSION}.tar.gz -C ${PWD}/site/gen-pulsarctldocs/generators/build/ .
	mv ${PWD}/site/gen-pulsarctldocs/generators/pulsarctl-site-${VERSION}.tar.gz ${PWD}/pulsarctl-site-${VERSION}.tar.gz

pulsarctl: 
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o pulsarctl