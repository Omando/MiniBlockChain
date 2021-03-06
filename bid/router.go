/* Initialize web api routes where each route defines a path, an http method,
an API function, and a name (for easy lookup of routes) */
package bid

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Instantiate a controller object so that routes can be initialized
var controller *Controller = &Controller{
	blockChain:     &BlockChain{
		Chain:        Blocks{},
		PendingBids:  Bids{},
		NetworkNodes: map[string]bool{},
	},
	currentNodeUrl: "",
}

// Define all routes (name, http method, path, and controller api)
var routes []Route = []Route{
	Route{
		Name:        "Index",
		Method:      "GET",
		Path:        "/",
		HandlerFunc: controller.Index,
	},
	Route{
		Name:        "GetBlockChain",
		Method:      "GET",
		Path:        "/blockchain",
		HandlerFunc: controller.GetBlockChain,
	},
	Route{
		Name:        "RegisterAndBroadcastBid",
		Method:      "POST",
		Path:        "/bid/broadcast",
		HandlerFunc: controller.RegisterAndBroadcastBid,
	},
	Route{
		Name:        "RegisterBid",
		Method:      "POST",
		Path:        "/bid",
		HandlerFunc: controller.RegisterBid,
	},
	Route{
		Name:        "RegisterAndBroadcastNode",
		Method:      "POST",
		Path:        "/register-and-broadcast-node",
		HandlerFunc: controller.RegisterAndBroadcastNode,
	},
	Route{
		Name:        "RegisterNode",
		Method:      "POST",
		Path:        "/register-node",
		HandlerFunc: controller.RegisterNode,
	},
	Route{
		Name:        " RegisterNodesBulk",
		Method:      "POST",
		Path:        "/register-nodes-bulk",
		HandlerFunc: controller.RegisterNodesBulk,
	},
	Route{
		Name:        "Mine",
		Method:      "GET",
		Path:        "/mine",
		HandlerFunc: controller.Mine,
	},
	Route{
		Name:        "ReceiveNewBlock",
		Method:      "POST",
		Path:        "/receive-new-block",
		HandlerFunc: controller.ReceiveNewBlock,
	},
	Route{
		Name:        "Consensus",
		Method:      "GET",
		Path:        "/consensus",
		HandlerFunc: controller.Consensus,
	},
	Route{
		Name:        "GetBidsForAuction",
		Method:      "GET",
		Path:        "/auction/{auctionId}",
		HandlerFunc: controller.GetBidsForAuction,
	},
	Route{
		Name:        "GetBidsForPlayer",
		Method:      "GET",
		Path:        "/player/{playerId}",
		HandlerFunc: controller.GetBidsForPlayer,
	},
}

func NewRouter(port string) *mux.Router {
	// Initialize a controller object that holds the blockchain and the url
	// of the node on which this controller  is running
	controller.currentNodeUrl  = "http://localhost" + port
	controller.blockChain.CreateNewBlock(100, "0", "0")	// genesis block

	/* mux.Router matches incoming requests against a list of registered routes and calls
	a handler for the route that matches the URL or other condition. It implements the
	http.Handler interface so it is compatible with the standard http.ServeMux.
	*/
	var router *mux.Router = mux.NewRouter().StrictSlash(true)

	// Configure the router with all route elements in the routes array. For example, this code:
	// router.Methods("GET").Path("/consensus").Handler(controller.Consensus).Name("Consensus")
	// means that any GET method sent to /consensus will be routed to controller.Consensus method
	// The Name() method has no effect on the path; it allows us to easily locate a route by name
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Handler(route.HandlerFunc).
			Name(route.Name)
	}

	// Utilize Router.Use to append logging middleware to the chain of middleware (in this case we only
	// have the logging middleware). Note the signature for Router.Use:
	// 	func (r *Router) Use(mwf ...MiddlewareFunc)
	//  type MiddlewareFunc func(http.Handler) http.Handler
	// Recall that http.Handler is an interface that defines a ServeHTTP method.
	// We return http.HandlerFunc which is a struct that implements the Handler interface
	router.Use( func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			handler.ServeHTTP(writer, request)
			log.Printf("[%s] [%s] [Ellapsed: %s]", request.Method, request.RequestURI, time.Since(start))
		})
	})

	/*//type MiddlewareFunc func(http.Handler) http.Handler
	var loggingHandlerFunc http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		log.Printf("[%s] [%s] [Ellapsed: %s]", request.Method, request.RequestURI, time.Since(start))
	}
	var mwLogging mux.MiddlewareFunc = func(handler http.Handler) http.Handler {
		return loggingHandlerFunc
	}
	router.Use( mwLogging )*/

	// Return the fully configured router
	return router
}