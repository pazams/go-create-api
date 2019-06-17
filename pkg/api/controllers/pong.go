package controllers

import (
	"net/http"
)

// PongController ..
type PongController struct{}

// NewPongController returns a new PongController
func NewPongController() *PongController {
	return &PongController{}
}

// Pong ..
func (pc *PongController) Pong(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return http.StatusOK, &struct{ ping string }{"ping pong"}
}
