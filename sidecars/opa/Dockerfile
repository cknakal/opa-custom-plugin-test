# syntax=docker/dockerfile:experimental
FROM gcr.io/cloud-builders/go@sha256:8939daf2defe9c3f344f95887b349628712a8124504cd10949b18a8e409600fa AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache \
    git \
    gcc \
    zip \
    rsync


#ca certs
RUN apk add -U --no-cache ca-certificates

ENV GOPATH=/go

WORKDIR /app
COPY main.go .
COPY plugin.go .
COPY go.mod .
COPY go.sum .
COPY sample-config sample-config
RUN go mod download

RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o opa_envoy_linux_amd64 .

FROM gcr.io/distroless/base

COPY --from=builder /app/opa_envoy_linux_amd64 /app/opa_envoy_linux_amd64
COPY --from=builder /app/sample-config /app/sample-config

WORKDIR /app

ENTRYPOINT [ \
    "./opa_envoy_linux_amd64", \
    "run", \
    "--server", \
    "--log-level=debug", \
    "--config-file=sample-config/sample-config.yaml", \
    "--diagnostic-addr=0.0.0.0:8282" \
]