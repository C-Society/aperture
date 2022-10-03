# Istio Configuration

This Chart inserts Envoy filters that integrate with Aperture Agent.

## Parameters

### Envoy Filter Parameters

| Name                           | Description                                                                                                                                                      | Value            |
| ------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------- |
| `envoyFilter.name`             | Name of service running aperture-agent                                                                                                                           | `aperture-agent` |
| `envoyFilter.namespace`        | Namespace where aperture-agent is running                                                                                                                        | `aperture-agent` |
| `envoyFilter.port`             | Port serving ext authz API and for streaming access logs                                                                                                         | `8080`           |
| `envoyFilter.authzGrpcTimeout` | Timeout in seconds to authz requests made to aperture-agent. Note: aperture-agent scheduler has max_timeout parameter that must tuned to match the setting here. | `0.5s`           |
| `envoyFilter.maxRequestBytes`  | Maximum size of request that is sent over ext authz API                                                                                                          | `8192`           |

