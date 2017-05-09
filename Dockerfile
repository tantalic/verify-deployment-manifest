FROM alpine:3.5
MAINTAINER Kevin Stock <kevin@toolhouse.com>

# SSL CA Root Certs
RUN apk --no-cache add ca-certificates

# Labels: http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="verify-deployment-manifest" \
      org.label-schema.description="A tool for verifying a deployment manifest JSON file from a server" \
      org.label-schema.url="https://github.com/toolhouse/verify-deployment-manifest" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/toolhouse/verify-deployment-manifest" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"

WORKDIR /
ADD ./verify-deployment-manifest-linux_amd64 /verify-deployment-manifest
EXPOSE 80

# The environment variables are used to configure the container at runtime:
# ENV REF 1.2.1
# ENV COMMIT 650bc8377cc3f7fc668b56915f5696ceb79cd36f
# ENV URL http://www.example.com/deployment.json

CMD ["/verify-deployment-manifest"]
