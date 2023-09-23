FROM envoyproxy/envoy:v1.26-latest
COPY ./envoy-k8s.yaml /etc/envoy/envoy.yaml
