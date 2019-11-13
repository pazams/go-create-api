package api

import (
	"net/http"

	"github.com/AndrewBurian/powermux"
	"github.com/pazams/go-create-api/pkg/api/controllers"
	"github.com/pazams/go-create-api/pkg/api/middlewares"
)

// NewRouter ..
func NewRouter(
	bc *controllers.BookController,
	pc *controllers.PongController,
	apiMid *middlewares.APIAuthMiddleware,
	corsMid *middlewares.CORSMiddleware,
) http.Handler {
	mux := powermux.NewServeMux()
	core := mux.Route("/").Middleware(corsMid)

	veririfedAPIToken := core.Middleware(apiMid)

	veririfedAPIToken.
		Route("/ping").
		Get(toHandler(pc.Pong))

	veririfedAPIToken.
		Route("/book").
		Get(toHandler(bc.Books))

	veririfedAPIToken.
		Route("/book/:id").
		Get(toHandler(bc.Book))

	veririfedAPIToken.
		Route("/book").
		Post(toHandler(bc.InsertBook))

	return mux
}
