#!/bin/sh
source sub_upstream_port.sh

envoy -c /etc/service-envoy.yaml --service-cluster service --log-level error
