package middlewares

import (
	"net/http"

	"github.com/pazams/go-create-api/pkg/api/config"
)

const apiAuthHeaderKey = "x-api-token"

// APIAuthMiddleware  ..
type APIAuthMiddleware struct {
	c *config.Config
}

// NewAPIAuthMiddleware ..
func NewAPIAuthMiddleware(c *config.Config) *APIAuthMiddleware {
	return &APIAuthMiddleware{
		c: c,
	}
}

// ServeHTTPMiddleware ..
func (m *APIAuthMiddleware) ServeHTTPMiddleware(rw http.ResponseWriter, req *http.Request, next func(http.ResponseWriter, *http.Request)) {

	if req.Method == http.MethodOptions {
		next(rw, req)
		return
	}

	token := req.Header.Get(apiAuthHeaderKey)
	if token != m.c.APIToken {
		rw.Header().Set("Content-Type", "")
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	next(rw, req)
}

// APIAuthHeaderKey returns the api authentication header key string
func (m *APIAuthMiddleware) APIAuthHeaderKey() string {
	return apiAuthHeaderKey
}
