//+build wireinject

package api

import (
	"github.com/google/wire"

	"github.com/pazams/go-create-api/pkg/api/config"
	"github.com/pazams/go-create-api/pkg/api/controllers"
	"github.com/pazams/go-create-api/pkg/api/data"
	"github.com/pazams/go-create-api/pkg/api/middlewares"
)

// InitializeServer resolves all dependencies for dependency injection and returns the server object
func InitializeServer() (*Server, error) {
	wire.Build(
		NewServer,
		NewRouter,
		config.New,
		data.New,
		middlewares.NewAPIAuthMiddleware,
		middlewares.NewCORSMiddleware,
		controllers.NewBookController,
		controllers.NewPongController,
	)
	return &Server{}, nil
}
