// "testhelpers" binary exposes an http server to perform side-effects requested by integration tests
// The motivation for a separate library is to prevent this logic to ever end up running in production

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pazams/go-create-api/pkg/api/config"
	"github.com/pazams/go-create-api/pkg/api/data"
)

func main() {

	config := config.New()
	dal := data.New(config)

	http.HandleFunc("/reset-db", func(w http.ResponseWriter, r *http.Request) {
		err := dal.MigrateDown()
		if err != nil {
			renderError(w, err, 500)
			return
		}

		err = dal.MigrateUp()
		if err != nil {
			renderError(w, err, 500)
			return
		}

		w.WriteHeader(200)
	})

	http.ListenAndServe(":3002", nil)
}

func renderError(w http.ResponseWriter, err error, status int) {
	log.Println(fmt.Sprintf("%v: %v", status, err))
	w.WriteHeader(status)
}
