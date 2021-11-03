# go_api_01

## Description

- golang server project
  - provide api
    - openapi

## TODO

- openapi
  - refactor test
  - error handling
  - add test pattern
  - modify path name
- graphql
- grpc

## Command

### Generate 

https://github.com/deepmap/oapi-codegen

```
oapi-codegen -generate "types" -package openapi api/definition.json > internal/pkg/oas-api/types.gen.go
oapi-codegen -generate "server" -package openapi api/definition.json > internal/pkg/oas-api/server.gen.go
```