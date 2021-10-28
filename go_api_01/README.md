```
oapi-codegen -generate "types" -package openapi api/definition.json > internal/pkg/oas-api/types.gen.go
oapi-codegen -generate "server" -package openapi api/definition.json > internal/pkg/oas-api/server.gen.go
```