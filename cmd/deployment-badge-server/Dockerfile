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
RUN go build -installsuffix 'static' ./cmd/deployment-badge-server/

FROM alpine:3.9.2

# SSL CA Root Certs
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /src/deployment-badge-server /deployment-badge-server

# Run as non-root user on port 8000
EXPOSE 8000
USER nobody
CMD ["/deployment-badge-server"]

# Labels: http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="deployment-badge-server" \
      org.label-schema.description="A server to display image badges with deployment information" \
      org.label-schema.url="https://github.com/toolhouse/deployment-manifest/tree/master/cmd/deployment-badge-server" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/toolhouse/deployment-manifest" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"
