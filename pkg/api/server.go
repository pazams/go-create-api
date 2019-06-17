package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pazams/go-create-api/pkg/api/config"
	"github.com/pazams/go-create-api/pkg/api/data"
)

// Server ..
type Server struct {
	srv *http.Server
	c   *config.Config
}

// NewServer ..
func NewServer(
	router http.Handler,
	d *data.DAL,
	c *config.Config,
) (*Server, error) {

	err := d.MigrateUp()
	if err != nil {
		return nil, err
	}
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
	log.Println(fmt.Sprintf("Running on %s", s.c.AppPort))
	log.Fatal(s.srv.ListenAndServe())
}
