package bid

import (
	"github.com/gorilla/mux"
)

var controller *Controller = &Controller{
	blockChain:     &BlockChain{
		Chain:        Blocks{},
		PendingBids:  Bids{},
		NetworkNodes: Nodes{},
	},
	currentNodeUrl: "",
}

var routes []Route = []Route{
	Route{
		Name:    "Index",
		Method:  "GET",
		Pattern: "/",
		HandlerFunc: controller.Index,
	},
	Route{
		Name:        "GetBlockChain",
		Method:      "GET",
		Pattern:     "/blockchain",
		HandlerFunc: controller.GetBlockChain,
	},
	Route{
		Name:        "RegisterAndBroadcastNode",
		Method:      "POST",
		Pattern:     "/register-and-broadcast-node",
		HandlerFunc: controller.RegisterAndBroadcastNode,
	},
	Route{
		Name:        "RegisterNode",
		Method:      "POST",
		Pattern:     "/register-node",
		HandlerFunc: controller.RegisterNode,
	},
}

func NewRouter(nodeAddress string) *mux.Router {

	/* mux.Router matches incoming requests against a list of registered routes and calls
	a handler for the route that matches the URL or other condition. It implements the
	http.Handler interface so it is compatible with the standard http.ServeMux.
	*/
	var router *mux.Router = mux.NewRouter().StrictSlash(true)
	return router;
}