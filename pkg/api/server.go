package api

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/pazams/go-create-api/pkg/api/config"
)

// Server ..
type Server struct {
	srv *http.Server
	c   *config.Config
}

// NewServer ..
func NewServer(
	router http.Handler,
	c *config.Config,
) (*Server, error) {

	address := fmt.Sprintf(":%s", c.AppPort)

	srv := &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}
	return &Server{
		srv: srv,
		c:   c,
	}, nil
}

// Start starts the server
func (s *Server) Start() {
	log.Infof("Running on %s", s.c.AppPort)
	log.Fatal(s.srv.ListenAndServe())
}
