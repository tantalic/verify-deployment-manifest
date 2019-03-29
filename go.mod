module github.com/toolhouse/deployment-manifest

go 1.12

replace (
	github.com/toolhouse/deployment-manifest/cmd/verify-deployment-manifest => ./cmd/verify-deployment-manifest
	github.com/toolhouse/deployment-manifest/pkg/deployment => ./pkg/deployment
)

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/hashicorp/errwrap v1.0.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/narqo/go-badge v0.0.0-20190124110329-d9415e4e1e9f
	github.com/pkg/errors v0.8.1
	github.com/the42/badge v0.0.0-20170523112329-0280203be5ca
	golang.org/x/image v0.0.0-20190321063152-3fc05d484e9f // indirect
)
