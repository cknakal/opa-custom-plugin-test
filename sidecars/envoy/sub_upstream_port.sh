#!/bin/sh

export UPSTREAM_PORT="${UPSTREAM_PORT:-8080}"
envsubst < /etc/service-envoy.yaml.template > /etc/service-envoy.yaml