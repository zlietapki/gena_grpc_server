[//]: # (start name:common)
microboiler
===========

Project boilerplate

[//]: # (start name:project type:add)

grpc_server
-----------

Boilerplate gRPC microservice template. Clean architecture, standard Go project layout.

### Structure

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
```

### Run

```bash
task run
# or
docker compose up
```

### Build

```bash
task build
# binary: bin/server
```

### Test / Lint

```bash
task test
task lint
```

### Health check

```bash
grpcurl -plaintext localhost:8888 grpc.health.v1.Health/Check
```

### Example requests

```shell
grpcurl -plaintext localhost:8888 list
grpcurl -plaintext localhost:8888 list api.proto.Example
grpcurl -plaintext localhost:8888 describe api.proto.Example.Hello
grpcurl -plaintext localhost:8888 describe .api.proto.HelloReq

grpcurl -plaintext -d @ localhost:8888 api.proto.Example.Hello <<EOM
{
  "name": "ZlieTapki"
}
EOM
```
