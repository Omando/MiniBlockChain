package bid

import (
	"github.com/gorilla/mux"
)

func NewRouter(nodeAddress string) *mux.Router {

	/* mux.Router matches incoming requests against a list of registered routes and calls
	a handler for the route that matches the URL or other condition. It implements the
	http.Handler interface so it is compatible with the standard http.ServeMux.
	*/
	var router *mux.Router = mux.NewRouter().StrictSlash(true)
	return router;
}