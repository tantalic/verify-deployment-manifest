# Deployment Manifest

[![Version](https://badge.fury.io/gh/toolhouse%2Fverify-deployment-manifest.svg)](https://github.com/toolhouse/verify-deployment-manifest/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/toolhouse/verify-deployment-manifest)](https://goreportcard.com/report/github.com/toolhouse/verify-deployment-manifest) [![codebeat badge](https://codebeat.co/badges/4c4cc430-53ea-4022-a05a-dd9e34534940)](https://codebeat.co/projects/github-com-toolhouse-verify-deployment-manifest-master) [![](https://images.microbadger.com/badges/image/toolhouse/verify-deployment-manifest.svg)](https://microbadger.com/images/toolhouse/verify-deployment-manifest "Docker Image") [![license](https://img.shields.io/github/license/toolhouse/verify-deployment-manifest.svg)](https://github.com/toolhouse/verify-deployment-manifest/blob/master/LICENSE)

A simple method to query information about the deployed version of a site or
application. This is accomplished by providing a JSON deployment manifest that
includes the version control system (typically git) commit id and reference
(branch or tag) from which an application/site was built. (We also recommend
including additional fields which may be useful for inspection by additional
tools and by humans.) 

## Why?

This was developed with two primary use-cases in mind:

- As part of post-deployment tests within a CI/CD pipeline to verify the
  desired version of an application has been successfully deployed.
- Provide the ability to easily and quickly inspect (manually or through tools)
  the currently deployed version of an application in any environment.

## Deployment Manifest Specification

```json
{
    "commit": "052b2762382f956c378c99aa626ca8faf3d76562",
    "ref": "1.1.3",
    "date": "Tue May 9 22:01:14 GMT 2017",
    "pipelineId": "764"
}
```

## Tools

This repository includes the following tools for working with deployment
manifests:

- `verify-deployment-manifest` - Validate that the expected commit/ref is
  deployed to an environment by checking a deployment manifest for expected
  versions.

[verify]: ./cmd/verify-deployment-manifest/
