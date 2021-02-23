FROM envoyproxy/envoy-alpine-dev:latest AS base

# add envsubst
RUN apk add gettext

COPY service-envoy.yaml.template /etc/service-envoy.yaml.template
COPY start_service.sh /usr/local/bin/start_service.sh
COPY sub_upstream_port.sh /usr/local/bin/sub_upstream_port.sh 
COPY ssl/key.pem /etc/key.pem
COPY ssl/crt.pem /etc/crt.pem

FROM base as testing

# Validate Envoy config
COPY validate_envoy_config.sh /usr/local/bin/validate_envoy_config.sh

RUN chmod u+x /usr/local/bin/validate_envoy_config.sh
RUN ls -lh /usr/local/bin/
RUN /usr/local/bin/validate_envoy_config.sh

FROM base

RUN chmod u+x /usr/local/bin/start_service.sh
ENTRYPOINT /usr/local/bin/start_service.sh