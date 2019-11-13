package middlewares

import (
	"net/http"
	"strings"
)

const (
	accessControlAllowOrigin  = "Access-Control-Allow-Origin"
	accessControlAllowMethods = "Access-Control-Allow-Methods"
	accessControlAllowHeaders = "Access-Control-Allow-Headers"
)

// CORSMiddleware  ..
type CORSMiddleware struct {
	corsMethods string
	corsHeaders string
}

// NewCORSMiddleware ..
func NewCORSMiddleware(
	apiMid *APIAuthMiddleware,
) *CORSMiddleware {

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

	return &CORSMiddleware{
		corsMethods: corsMethods,
		corsHeaders: corsHeaders,
	}

}

// ServeHTTPMiddleware writes CORS headers
func (m *CORSMiddleware) ServeHTTPMiddleware(rw http.ResponseWriter, req *http.Request, next func(http.ResponseWriter, *http.Request)) {
	rw.Header().Set(accessControlAllowMethods, m.corsMethods)
	rw.Header().Set(accessControlAllowHeaders, m.corsHeaders)
	rw.Header().Set(accessControlAllowOrigin, "*")

	if req.Method == http.MethodOptions {
		return
	}

	next(rw, req)
}
