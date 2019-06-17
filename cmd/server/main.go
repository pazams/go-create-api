package main

import (
	"github.com/pazams/go-create-api/pkg/api"
)

func main() {
	server, err := api.InitializeServer()
	if err != nil {
		panic(err)
	}
	server.Start()
}
