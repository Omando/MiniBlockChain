/* Define all API methods for each available route. These methods are defined on the
Controller struct. Recall that the Controller struct holds the current node's URL
and own copy of the blockchain */
package bid

import (
	"encoding/json"
	"net/http"
)

// GetBlockChain GET /blockchain
// Retrieves the blockchain in JSON
func (c *Controller) GetBlockChain(writer http.ResponseWriter, request *http.Request) {
	// Setup headers
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)

	// Convert the blockchain into JSON
	data, _ := json.Marshal(c.blockChain)
	writer.Write(data)
	return
}

// RegisterAndBroadcastBid POST /bid/broadcast
func (c *Controller) RegisterAndBroadcastBid(writer http.ResponseWriter, request *http.Request) {
	// continue here
}


// Index GET/
func (c *Controller) Index(writer http.ResponseWriter, request *http.Request) {

}

// RegisterAndBroadcastNode POST /register-and-broadcast-node
func (c *Controller) RegisterAndBroadcastNode(writer http.ResponseWriter, request *http.Request) {

}

// RegisterNode POST /register-node
func (c *Controller) RegisterNode(writer http.ResponseWriter, request *http.Request) {

}

// RegisterNodesBulk POST /register-nodes-bulk
func (c *Controller) RegisterNodesBulk(writer http.ResponseWriter, request *http.Request) {

}

// RegisterBid POST/bid
func (c *Controller) RegisterBid(writer http.ResponseWriter, request *http.Request) {

}

// Mine GET /mine
func (c *Controller) Mine(writer http.ResponseWriter, request *http.Request) {

}

// ReceiveNewBlock POST /review-new-block
func (c *Controller) ReceiveNewBlock(writer http.ResponseWriter, request *http.Request) {

}
// Consensus GET /consensus
func (c *Controller) Consensus(writer http.ResponseWriter, request *http.Request) {

}

// GetBidsForAuction GET /auction/{auctionId} retrieves all bids for an auction
func (c *Controller) GetBidsForAuction(writer http.ResponseWriter, request *http.Request) {

}

// GetBidsForPlayer GET /player/{playerId}
func (c *Controller) GetBidsForPlayer(writer http.ResponseWriter, request * http.Request) {

}


