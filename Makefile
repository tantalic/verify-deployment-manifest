DOCKER_IMAGE := toolhouse/verify-deployment-manifest
COMMIT := $(strip $(shell git rev-parse --short HEAD))
VERSION := $(strip $(shell git describe --always --dirty))

.PHONY: docker-build docker-push help
.DEFAULT_GOAL := help	

docker-image: ## Build a docker image
	docker build \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(COMMIT) \
		-t $(DOCKER_IMAGE):$(VERSION) .

docker-push: ## Push the docker image to DockerHub
	docker push $(DOCKER_IMAGE):$(VERSION)

help: ## Print available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
