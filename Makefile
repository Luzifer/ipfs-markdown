VERSION = $(shell git describe --tags || git rev-parse --short HEAD)

default: container


build:
	docker run -v $(CURDIR):/src \
	-e LDFLAGS='-X main.version $(VERSION)' \
	centurylink/golang-builder:latest

container: build
	docker build -t registry.luzifer.io/ipfsmd .
	docker push registry.luzifer.io/ipfsmd
