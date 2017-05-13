# Verify Deployment Manifest

[![Version](https://badge.fury.io/gh/toolhouse%2Fverify-deployment-manifest.svg)](https://github.com/toolhouse/verify-deployment-manifest/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/toolhouse/verify-deployment-manifest)](https://goreportcard.com/report/github.com/toolhouse/verify-deployment-manifest) [![codebeat badge](https://codebeat.co/badges/4c4cc430-53ea-4022-a05a-dd9e34534940)](https://codebeat.co/projects/github-com-toolhouse-verify-deployment-manifest-master) [![](https://images.microbadger.com/badges/image/toolhouse/verify-deployment-manifest.svg)](https://microbadger.com/images/toolhouse/verify-deployment-manifest "Docker Image") [![license](https://img.shields.io/github/license/toolhouse/verify-deployment-manifest.svg)](https://github.com/toolhouse/verify-deployment-manifest/blob/master/LICENSE)

A simple tool to verify the deployed site/application is running the desired version. The application fetches a deployment manifest over HTTP and checks the values in the manifest against the desired values. The deployment manifest is a JSON document that includes the current commit and ref (branch/tag). Additional fields are ignored by this tool. It is recommended to include additional fields with other information about the deployment that may be used by other tools or inspection by humans.

Example deployment manifest:

```json
{
    "commit": "052b2762382f956c378c99aa626ca8faf3d76562",
    "ref": "v1.1.3",
    "date": "Tue May  9 22:01:14 GMT 2017",
    "pipelineId": "764"
}
```

If the desired commit and/or ref values match those in the manifest the application will exit with code 0. If values do not match the application will exit with a code greater then 0.

## Why?

This was developed with two primary use-cases in mind:

- As part of post-deployment tests within a CI/CD pipeline
- To verify/detect the completion of a deployment that includes manual steps


## How?

The application is primarily designed to be run inside a Docker container (although it can also be run as a standalone binary). The application is configured through the following environment variables:

| Environment Variable | Description                                                                                                                          | Example                                |
|----------------------|--------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| `URL`                | The URL of the deployment manifest file for the deployment to check.                                                                 | http://www.example.com/deployment.json |
| `COMMIT`             | The commit that should be deployed to check for. This can be either the full commit SHA or an abbreviated (short) prefix to the SHA. | b252eb498a07                           |
| `REF`                | The reference (branch or tag) name that should be deployed to check for.                                                             | master                                 |

`URL` is required. You also must define `REF` or `COMMIT` (or both).

### Example

For example, to check for a deployment for example.com which should be on tag `v1.3.1` and commit `b252eb498a0791db07496601ebc7a059dd55cfe9`

```shell
docker run --env URL=https://www.example.com/deployment.json --env REF=v1.3.1 --env COMMIT=b252eb498a07 toolhouse/verify-deployment-manifest:v0.2.1
```
