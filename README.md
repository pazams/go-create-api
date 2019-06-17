# Create Go API

## What
An template for a go JSON api with Google App Engine ready deployment.

## Features
- DI with google/wire
- Integration tests are written in a different language (JS) - makes it harder to cheat the block box.
- Tests side effects needed for tests (e.g. migrate DB down and up), are decoupled into a separate binary.

## How 

### Add dependencies to DI(wire)
Install wire: `$ GO111MODULE=on go get github.com/google/wire/cmd/wire`
Generate `wire_gen.go`: `$ GO111MODULE=on wire ./pkg/api/`

If you get "pkg/api/wire_gen.go: not enough arguments in call to ..." it means we need to regenerate wiring

### Build (updates go.mod)
$ GO111MODULE=on go build ./cmd/server/

### Run
```bash 
$ ./run.sh
```

### Test
```bash 
$ ./test.sh
```

### Test iteration:
$ docker-compose up --build --no-deps test

### Deploy
Expand app.yaml for GAE
```base
$ POSTGRES_PASSWORD=<the-password> envsubst < "app-subst.yaml" > "app.yaml"
```

$ GO111MODULE=on gcloud app deploy --project=<your GCP project>
