module github.com/toolhouse/deployment-manifest

go 1.12

replace (
	github.com/toolhouse/deployment-manifest/cmd/verify-deployment-manifest => ./cmd/verify-deployment-manifest
	github.com/toolhouse/deployment-manifest/pkg/deployment => ./pkg/deployment
)

require (
	github.com/hashicorp/errwrap v1.0.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/pkg/errors v0.8.1
)
