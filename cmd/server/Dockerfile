# NOTE: this docker file needs to be run from the context of the top level repo folder
FROM golang:1.23 as build-env
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./cmd/server/

FROM alpine
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=build-env /go/bin/app /go/bin/app
COPY migrations/ /go/bin/migrations
WORKDIR /go/bin
ENTRYPOINT ["./app"]
