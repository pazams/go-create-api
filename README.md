# Create Go API

## What
A template for a go JSON api with Google App Engine ready deployment.

## Features
- DI with google/wire
- Integration tests are written in a different language (JS) - makes it harder to cheat the black box.
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
$ ./run.sh
```

### Test
```bash 
$ ./test.sh
```
Note: while iterating on the tests, after running the test command, if the app code didn't change, we can just build and run the tests with while everything else is still running: `$ docker-compose up --build --no-deps test`

### Deploy
Expand app.yaml for GAE with secrets (avoids secrets commited to source control)
```bash
$ API_TOKEN=<the-api-token> POSTGRES_GCP_CONNECTION_NAME=<project:zone:instance> POSTGRES_PASSWORD=<the-password> envsubst < "app-subst.yaml" > "app.yaml"
```

Deploy on GCP:
```bash
$ GO111MODULE=on gcloud app deploy --project=<your-GCP-project>
```
