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

// Define all routes (name, http method, path, and controller api)
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
	Route{
		Name:        " RegisterNodesBulk",
		Method:      "POST",
		Pattern:     "/register-nodes-bulk",
		HandlerFunc: controller.RegisterNodesBulk,
	},
	Route{
		Name:        "RegisterBid",
		Method:      "POST",
		Pattern:     "/bid",
		HandlerFunc: controller.RegisterBid,
	},
	Route{
		Name:        "RegisterAndBroadcastBid",
		Method:      "POST",
		Pattern:     "/bid/broadcast",
		HandlerFunc: controller.RegisterAndBroadcastBid,
	},
	Route{
		Name:        "Mine",
		Method:      "GET",
		Pattern:     "/mine",
		HandlerFunc: controller.Mine,
	},
	Route{
		Name:        "ReceiveNewBlock",
		Method:      "POST",
		Pattern:     "/review-new-block",
		HandlerFunc: controller.ReceiveNewBlock,
	},
	Route{
		Name:        "Consensus",
		Method:      "GET",
		Pattern:     "/consensus",
		HandlerFunc: controller.Consensus,
	},
	Route{
		Name:        "GetBidsForAuction",
		Method:      "GET",
		Pattern:     "/auction/{auctionId}",
		HandlerFunc: controller.GetBidsForAuction,
	},
	Route{
		Name:        "GetBidsForPlayer",
		Method:      "GET",
		Pattern:     "/player/{playerId}",
		HandlerFunc: controller.GetBidsForPlayer,
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