package main

import (
	"MiniBlockChain/bid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
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

type MyHandler struct {
}

// Implement Handler interface
func (my MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func basicHttpUsage() {
	// error handling removed for brevity ...

	// http.Handle registers the handler for the given pattern in the DefaultServeMu
	// You can handle a given pattern using either http.Handle or http.HandleFunc
	// For http.Handle, you need have a type that implements the Handler interface
	myHandler := MyHandler{}
	http.Handle("/foo",  myHandler)

	// This is a much easier way to handle a given pattern
	http.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		// ...
	})

	// ListenAndServe listens on the TCP network address addr and then calls Serve with
	// handler to handle requests on incoming connections. Accepted connections are
	// configured to enable TCP keep-alives. The handler is typically nil, in which case
	// the DefaultServeMux is used.
	http.ListenAndServe(":8080", nil)

	// More control over the server's behavior is available by creating a custom Server
	// (comment the previous line above)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}


