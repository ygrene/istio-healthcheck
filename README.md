# Ygrene Istio Healthcheck

## Why?
Because: https://istio.io/help/ops/security/health-checks-and-mtls/
* The Kubernetes health check comes from the Kubelet, which isn't (nor will be) a participant in the service mesh

## Ok, how do I use it?
In your Dockerfile run the following:

```dockerfile
RUN curl -L -o istiohealthcheck.tar.gz https://github.com/ygrene/istio-healthcheck/releases/download/v1.0.0/istio-healthcheck_1.0.0_Linux_x86_64.tar.gz
RUN tar -vxf istiohealthcheck.tar.gz
```

Then somewhere in your Docker Entrypoint:

`nohup ./istio-healthcheck &`

## Contributing
* You can contribute via PR and then you'll need `goreleaser` to actually run a release