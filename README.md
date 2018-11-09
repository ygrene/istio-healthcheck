# Ygrene Istio Healthcheck

## Why?
Because: https://istio.io/help/ops/security/health-checks-and-mtls/
* The Kubernetes health check comes from the Kubelet, which isn't (nor will be) a participant in the service mesh

## Ok, how do I use it?
In your Dockerfile run the following:

```dockerfile
RUN 
```