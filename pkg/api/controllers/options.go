package controllers

import (
	"net/http"
	"strings"

	"github.com/pazams/go-create-api/pkg/api/middlewares"
)

const (
	accessControlAllowOrigin  = "Access-Control-Allow-Origin"
	accessControlAllowMethods = "Access-Control-Allow-Methods"
	accessControlAllowHeaders = "Access-Control-Allow-Headers"
)

// OptionsController  ..
type OptionsController struct {
	corsMethods string
	corsHeaders string
}

// NewOptionsController ..
func NewOptionsController(
	apiMid *middlewares.APIAuthMiddleware,
) *OptionsController {

	corsMethods := strings.Join(
		[]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		", ")

	corsHeaders := strings.Join(
		[]string{
			apiMid.APIAuthHeaderKey(),
			"Content-Type",
			"Cache-Control",
			// add more if needed
		},
		", ")

	return &OptionsController{
		corsMethods: corsMethods,
		corsHeaders: corsHeaders,
	}

}

// Options ..
func (c *OptionsController) Options(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(accessControlAllowMethods, c.corsMethods)
	w.Header().Set(accessControlAllowHeaders, c.corsHeaders)
	w.Header().Set(accessControlAllowOrigin, "*")
}
