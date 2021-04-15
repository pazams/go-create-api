package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

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

	// setup logging
	if c.AppEnv == "GAE" {
		h, err := NewStackDriverHook("app", c.ProjectID)
		if err != nil {
			return nil, err
		}
		log.AddHook(h)

		// disalbe std logging
		// See https://github.com/Sirupsen/logrus/issues/328
		log.SetOutput(ioutil.Discard)
		log.Info("Gcloud logging success")
	}

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
	log.Infof("Running on %s", s.c.AppPort)
	log.Fatal(s.srv.ListenAndServe())
}
