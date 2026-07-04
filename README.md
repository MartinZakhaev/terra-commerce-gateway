# Terra Commerce Gateway

Go API gateway and backend-for-frontend for Terra Commerce.

## Responsibilities

- Resolve tenant context from the approved child-subdomain URL model.
- Own opaque browser sessions and the hosted ZITADEL OIDC callback flow.
- Enforce the machine-readable operation-permission registry.
- Validate idempotency and optimistic-version headers.
- Proxy trusted internal commands and queries to the commerce runtime.
- Serve audience-filtered Server-Sent Events.
- Expose liveness and readiness endpoints.

The normative specifications live in `MartinZakhaev/terra-commerce-master-spec`.

## Local development

```bash
go run ./cmd/gateway
```

The default listener is `:8080`. Copy `.env.example` into your preferred local environment loader before adding external integrations.
