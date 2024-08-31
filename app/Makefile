go-lint:
ifeq (, $(shell which golangci-lint))
	$(error golangci-lint binary is not found in path; install it from https://golangci-lint.run/usage/install/)
endif
	golangci-lint run --timeout=10m

run:
	go run main.go

mod:
	go mod tidy
	go mod verify

mod-upgrade:
	go get -u

publish:
ifeq (, $(shell which ko))
	$(error ko binary is not found in path; install it from https://ko.build/install/)
endif
	KO_DOCKER_REPO='ghcr.io/slamdev/sample-cli' ko build --bare \
	 --tags=$${GIT_TAG:-develop} \
	 --image-label org.opencontainers.image.source=$(shell git remote get-url origin) \
	 --image-label org.opencontainers.image.revision=$(shell git rev-parse HEAD)
