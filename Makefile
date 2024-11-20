.PHONY: build
build:
	go build -race cmd/server/main.go

.PHONY: test
test:
	PORT=0 APP_ENV=test go test -v ./pkg/api/

.PHONY: serve
serve:
	./__scripts/run.sh