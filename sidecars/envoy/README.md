# Envoy Sidecar
This is a very thin wrapper around the standard Envoy container.

Where this differs:
* [Self signed certs](ssl) are added to support TLS termination.
* Configuration template is added to the container to configure Envoy in the way we need.
  * Currently the only part that is templatized is the [UPSTREAM_PORT](service-envoy.yaml.template#L164)

## Notable configuration elements

* UPSTREAM_PORT: This allows us to override the port Envoy will send traffic to. The default value for this is `8080`
* Standard Endpoints
  * /health, /env, /docs. /swagger are all defined within the config to not [require authorization](service-envoy.yaml.template#L21-L48).
* All other requests are configured to be sent to the [OPA sidecar](service-envoy.yaml.template#L49-L64)
* TLS termination [is configured](service-envoy.yaml.template#L65-L74) using the SSL certs described above
* with_request_body.with_request_body.max_request_bytes is set to 10 KB which means post body can have a max size of 
10KB