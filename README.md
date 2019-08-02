# knative-gloo-grpc

## Requirements

- gloo gateway-proxy: `glooctl install gateway`
- VirtualService: `virtualservice.yaml` (HTTP-to-GRPC transformation)
	- `namespace` MUST BE `gloo-system`
	- **DON'T FORGET** to change the name of the upstream according to your upstream assigned name
- knative (`glooctl install knative -k` [-k -e if you also want eventing])