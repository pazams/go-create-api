.PHONY: build 
build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/server/main.go

.PHONY: deploy
deploy:
	./__scripts/deploy.sh

.PHONY: deploy-sec-conf
deploy-sec-conf:
	./__scripts/deploy-sec-conf.sh

.PHONY: serve
serve:
	./__scripts/run.sh

.PHONY: test
test:
	./__scripts/test.sh
