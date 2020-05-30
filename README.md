# Endpoint Readiness Check Action

Github Action to run docker compose if required and check the readiness of an endpoint of the service that one wants to run integration test against.

# Usage

This is a very simple action which performs readniness check of an endpoint of a sevice that one wishes to run integration test against. The `timeout` and the `docker-compose` parameters are optional see below the comment.

```
  - name: Check the readiness of the endpoint
    uses: akhettar/readiness-check@master
    with:
      readiness-endpoint: 'http://localhost:8080/v1' # the readiness endpoint
      timeout: '2' # timeout in seconds how long to wait for the readiness endpoint to become reacheable
      docker-compose: 'true' # Run docker compose command as part of checking for the readiness endpoint 
```        
