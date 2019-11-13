[![Build Status](https://travis-ci.com/pazams/go-create-api.svg?branch=master)](https://travis-ci.com/pazams/go-create-api)

# Create Go API

## What
A template for a go JSON api with Google App Engine ready deployment.

## Why
As of Go version 1.11, GAE offers a standard environment that lifts all previous restrictions (See https://blog.golang.org/appengine-go111 and https://cloud.google.com/blog/products/application-development/go-1-11-is-now-available-on-app-engine).

## Features
- Dependency injection with `google/wire`
- Integrates with postgres on [Cloud SQL](https://cloud.google.com/sql/), as well as local dockerised postgres for integration testing.
- Integration tests are totally decoupled for app code. By using a different language (JS), it prevents reuse of struct definitions and enforces a black box approach.
- System side effects needed only for tests (e.g. migrate DB down and up), are decoupled into a separate binary.

## How to

### Make it your own (fork)
- Fork this repository
- Rewrite all the imports and other file to point to your fork with:
```bash
$ perl -i -p -e 's/github.com\/pazams\/go-create-api/github.com\/<your-gh-handle>\/<your-repo-name>/g' **/*.*
```

### Add dependencies to DI(wire)
- Install wire: `$ GO111MODULE=on go get github.com/google/wire/cmd/wire`
- Generate `wire_gen.go`: `$ GO111MODULE=on wire ./pkg/api/`
- When adding code, if you get "pkg/api/wire_gen.go: not enough arguments in call to ..." it means we need to regenerate wiring.

### Build (updates go.mod)
`$ GO111MODULE=on go build ./cmd/server/`

### Run
```bash 
$ make serve
```

### Test
```bash 
$ make test
```
Note: while iterating on the tests, after running the test command, if the app code didn't change, we can just build and run the tests with while everything else is still running: `$ docker-compose up --build --no-deps test`

### Deploy
have these env vars set:
```
API_TOKEN=<the-api-token>
POSTGRES_GCP_CONNECTION_NAME=<project_id:zone:instance>
POSTGRES_PASSWORD=<the-password>
GCP_PROJECT=<project_id>
```

Then, deploy with:
```bash
$ make deploy
```

Note: 
- `app-subst.yaml` is committed to source and does NOT have secrets expended
- `app.yaml` is NOT committed to source control and does have secrets expended

## TODO
- [ ] micro services?
