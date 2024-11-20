package controllers

import (
	"net/http"

	"github.com/pazams/go-create-api/pkg/api/models"
)

// PingController ..
type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

func (c *PingController) Ping(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return http.StatusOK, models.PingResponse{Message: "pong"}
}
