# `verify-deployment-manifest`

[![](https://images.microbadger.com/badges/image/toolhouse/verify-deployment-manifest.svg)](https://microbadger.com/images/toolhouse/verify-deployment-manifest "Docker Image") 

A tool for verifying a site/application is running the desired version by
checking the deployment manifest. 

## How to use

The application is primarily designed to be run inside a Docker container
(although it can also be run as a standalone binary). The application is
configured through the following environment variables:

| Environment Variable | Description                                                                                                                          | Example                                |
|----------------------|--------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------|
| `URL`                | The URL of the deployment manifest file for the deployment to check.                                                                 | http://www.example.com/deployment.json |
| `COMMIT`             | The commit that should be deployed to check for. This can be either the full commit SHA or an abbreviated (short) prefix to the SHA. | b252eb498a07                           |
| `REF`                | The reference (branch or tag) name that should be deployed to check for.                                                             | master                                 |

`URL` is required. You also must define `REF` or `COMMIT` (or both).

If the desired commit and/or ref values match those in the manifest the
application will exit with code 0. If values do not match the application will
exit with a code greater then 0.

### Example

For example, to check for a deployment for example.com which should be on tag
`1.3.1` and commit `b252eb498a0791db07496601ebc7a059dd55cfe9`.

```shell
docker run \
   --env URL=https://www.example.com/deployment.json \
   --env REF=1.3.1 \
   --env COMMIT=b252eb498a0791db07496601ebc7a059dd55cfe9 \
   toolhouse/verify-deployment-manifest:latest
```
