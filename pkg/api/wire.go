//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"

	"github.com/pazams/go-create-api/pkg/api/config"
	"github.com/pazams/go-create-api/pkg/api/controllers"
)

// InitializeServer resolves all dependencies for dependency injection and returns the server object
func InitializeServer() (*Server, error) {
	wire.Build(
		NewServer,
		NewRouter,
		config.New,
		controllers.NewPingController
	)
	return &Server{}, nil
}
