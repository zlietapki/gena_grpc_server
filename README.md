# microboiler_grpc_server

Boilerplate gRPC microservice template. Clean architecture, standard Go project layout.

## Structure

```
cmd/server/          — entrypoint
internal/
  config/            — env-based config
  domain/            — entities, interfaces (Repository)
  repository/        — infrastructure stub impl
  usecase/           — business logic
  grpc_handler/      — gRPC transport handlers
pkg/
  grpcserver/        — gRPC server + interceptors
api/                 — proto/contract references
```

## Run

```bash
cp .env.example .env
task run
```

## Build

```bash
task build
# binary: bin/server
```

## Test / Lint

```bash
task test
task lint
```

## Health check

```bash
grpcurl -plaintext localhost:50051 grpc.health.v1.Health/Check
```
