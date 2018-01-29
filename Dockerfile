FROM golang:1.9.3 as builder

# Install dep to use for dependency management
ENV DEP_VERSION 0.3.2
RUN curl -o /usr/local/bin/dep -L https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 && \
      chmod a+x /usr/local/bin/dep

WORKDIR /go/src/github.com/toolhouse/verify-deployment-manifest

# Install dependencies
COPY Gopkg.* ./
RUN dep ensure -vendor-only

# Build go binary
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'


FROM alpine:3.5

# SSL CA Root Certs
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/toolhouse/verify-deployment-manifest/verify-deployment-manifest /verify-deployment-manifest
CMD ["/verify-deployment-manifest"]
WORKDIR /

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


# The environment variables are used to configure the container at runtime:
# ENV REF 1.2.1
# ENV COMMIT 650bc8377cc3f7fc668b56915f5696ceb79cd36f
# ENV URL http://www.example.com/deployment.json
