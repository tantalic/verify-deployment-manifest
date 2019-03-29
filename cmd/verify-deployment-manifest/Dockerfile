ARG GO_VERSION=1.12.1

FROM golang:${GO_VERSION}-alpine as builder

# Install git (required to fetch dependencies)
RUN apk update && apk add git && rm -rf /var/cache/apk/*

# Build Settings
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /src

# Install dependencies
COPY ./go.mod ./go.sum          ./
RUN go mod download

# Build go binary
COPY . .
RUN go build -installsuffix 'static' ./cmd/verify-deployment-manifest/

FROM alpine:3.5

# SSL CA Root Certs
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /src/verify-deployment-manifest /verify-deployment-manifest
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
