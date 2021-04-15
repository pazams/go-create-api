package api

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// ResultFunc is like http.HandlerFunc but returns values
type resultFunc func(http.ResponseWriter, *http.Request) (int, interface{})

// toHandler wraps a ResultFunc with a http.Handler func that handels the results and calls rendering logic
func toHandler(rf resultFunc) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		status, i := rf(w, r)

		switch v := i.(type) {
		case error:
			renderError(w, v, status)
		default:
			if v == nil {
				renderEmpty(w, status)
				return
			}
			renderJSON(w, v, status)
		}

	}
	return http.HandlerFunc(f)
}

func renderJSON(w http.ResponseWriter, model interface{}, status int) {
	w.Header().Set("Content-Type", "application/json") // Must be called before WriteHeader
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(model); err != nil {
		renderError(w, err, 500)
	}
}

func renderEmpty(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func renderError(w http.ResponseWriter, err error, status int) {
	log.WithError(err).Error(status)
	w.WriteHeader(status)
}
