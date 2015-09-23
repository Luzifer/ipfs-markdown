VERSION = $(shell git describe --tags || git rev-parse --short HEAD)

default: container


build:
	docker run \
			--rm \
			-v $(CURDIR):/usr/code \
			-e GOPATH=/usr/code/.gobuild:/usr/code/.gobuild/src/$(PROJECT_NAMESPACE)/$(PROJECT)/Godeps/_workspace \
			-w /usr/code \
			golang:1.5 \
			go build -a -ldflags "-X main.version=$(VERSION)" -o $(BIN)

container: build
	docker build -t registry.luzifer.io/ipfsmd .
