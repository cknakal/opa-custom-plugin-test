# opa-custom-plugin-test

## Quick Start
Run the OPA sidecar and Envoy proxy
```
$ make run-dev
```
Query the Envoy proxy in hopes of Envoy forwarding the request to OPA
```
$ curl localhost:8080/v1/data
```