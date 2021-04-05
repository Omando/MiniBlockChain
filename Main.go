package main

import (
	"MiniBlockChain/bid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	// Port to listen to
	port := os.Args[1]

	/* The 'handlers' package is a collection of handlers (aka "HTTP middleware") for use with
	Go's net/http package. The package includes handlers for logging in standardised formats,
	compressing HTTP responses, validating content types and other useful tools for manipulating
	requests and responses.

	AllowedOrigins sets the allowed origins for CORS requests, as used in the
	Allow-Access-Control-Origin' HTTP header. Passing in a "*" will allow any domain.
	AllowedMethods explicitly allow methods in the Access-Control-Allow-Methods header.

	Note the following definitions:
	type CORSOption func(*cors) error
	type cors struct {
		h                      http.Handler
		allowedHeaders         []string
		allowedMethods         []string
		allowedOrigins         []string
		allowedOriginValidator OriginValidator
		exposedHeaders         []string
		maxAge                 int
		ignoreOptions          bool
		allowCredentials       bool
		optionStatusCode       int
	} */
	var allowedOrigins handlers.CORSOption = handlers.AllowedOrigins([]string{"*"})
	var allowedMethods handlers.CORSOption = handlers.AllowedMethods([]string{"GET","POST"})

	// Initialize headers: accept calls from any origin and work with GET and POST requests
	var funcHandler func(http.Handler) http.Handler = handlers.CORS(allowedMethods, allowedOrigins)

	// Listen to port defined in port
	var router *mux.Router = bid.NewRouter(port)
	var handler http.Handler = funcHandler(router);
	http.ListenAndServe(":"+port, handler)
}

