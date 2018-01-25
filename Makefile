PKG := github.com/toolhouse/verify-deployment-manifest
DOCKER_IMAGE := toolhouse/verify-deployment-manifest
GOVERSION := 1.9.2

COMMIT := $(strip $(shell git rev-parse --short HEAD))
VERSION := $(strip $(shell git describe --always --dirty))

.PHONY: linux-amd64 docker-build docker-push update-ca help
.DEFAULT_GOAL := help	

linux-amd64:
	docker run --env GOOS=linux --env GOARCH=amd64 --env CGO_ENABLED=0 --rm -v "`pwd`":"/go/src/$(PKG)" -w /go/src/$(PKG) golang:$(GOVERSION) go build -a -tags netgo -ldflags '-w' -o verify-deployment-manifest-linux_amd64

docker-image: linux-amd64 ## Build a docker image
	docker build \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(COMMIT) \
		-t $(DOCKER_IMAGE):$(VERSION) .

docker-push: ## Push the docker image to DockerHub
	docker push $(DOCKER_IMAGE):$(VERSION)

help: ## Print available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
