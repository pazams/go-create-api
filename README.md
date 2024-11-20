# How to run

## Run with docker
```bash 
$ make serve
```

while it's running you can test the api with:
```bash 
$ curl -i -s -X GET localhost:3001/v1/ping
```

See the `docker-compose` file to edit any env variables
Here are the env variables definitions:
```
PORT        // any port number
APP_ENV     // "test" or other string
HAS_PROXY   // "true" | "false" true should be set if deployed behind a reverse proxy. default to "false"
```

## Test
```bash 
$ make test
```

## Build
```bash
$ make build
```

### Add dependencies to DI(wire)
- `go install github.com/google/wire/cmd/wire@latest`
- Generate `wire_gen.go`: `$ wire ./pkg/api/`
- When adding code, if you get "pkg/api/wire_gen.go: not enough arguments in call to ..." it means we need to regenerate wiring.