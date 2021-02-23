# opa-custom-plugin-test

# How to run
1. Build an OPA executable with the plugin
```
go build -o opa++ .
```
2. Run the executable with the provided config file
```
./opa++ run --server --config-file sample-config/sample-config.yaml
```
3. Exercise the plugin via the OPA API
```
curl localhost:8181/v1/data
```
4. If everything worked you will see the string `"You did it!"` written to stdout.