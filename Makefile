MINOR_VERSION=1

# Build pulsarctl docs

cleancli:
	rm -f main
	rm -rf $(shell pwd)/site/gen-pulsarctldocs/generators/includes
	rm -rf $(shell pwd)/site/gen-pulsarctldocs/generators/build
	rm -rf $(shell pwd)/site/gen-pulsarctldocs/generators/manifest.json

cli: cleancli
	go run site/gen-pulsarctldocs/main.go --pulsar-version v1_$(MINOR_VERSION)
	docker run -v ${PWD}/site/gen-pulsarctldocs/generators/includes:/source -v ${PWD}/site/gen-pulsarctldocs/generators/build:/build -v ${PWD}/site/gen-pulsarctldocs/generators/:/manifest pwittrock/brodocs
